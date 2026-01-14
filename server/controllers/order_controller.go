package controllers

import (
	"demowebgo/config"
	"demowebgo/models"
	"demowebgo/services"
	"demowebgo/utlis"
	"fmt"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

// Lấy tất cả đơn hàng
func GetAllOrders(c *gin.Context) {
	var orders []models.Order
	if err := config.DB.Preload("OrderItems.Book").Find(&orders).Error; err != nil {
		utils.Error(c, http.StatusInternalServerError, err.Error())
		return
	}
	utils.Success(c, http.StatusOK, orders, "Get orders successfully")
}

// Lấy đơn hàng theo id
func GetOrderByID(c *gin.Context) {
	var order models.Order
	id := c.Param("id")
	if err := config.DB.Preload("OrderItems.Book").First(&order, id).Error; err != nil {
		utils.Error(c, http.StatusNotFound, "Order not found")
		return
	}
	utils.Success(c, http.StatusOK, order, "Get order successfully")
}

// Lấy đơn hàng theo user ID
func GetOrdersByUserID(c *gin.Context) {
	var orders []models.Order
	userID := c.Param("userId")

	if err := config.DB.Where("user_id = ?", userID).Preload("OrderItems.Book").Find(&orders).Error; err != nil {
		utils.Error(c, http.StatusInternalServerError, err.Error())
		return
	}
	utils.Success(c, http.StatusOK, orders, "Get user orders successfully")
}

// Tạo mới đơn hàng
func CreateOrder(c *gin.Context) {
	var input struct {
		ShippingAddress string `json:"shipping_address"`
		PaymentMethod   string `json:"payment_method"`
		Items []struct {
			BookID   uint `json:"book_id"`
			Quantity int  `json:"quantity"`
		} `json:"items"`
	}

	if err := c.ShouldBindJSON(&input); err != nil {
		utils.Error(c, http.StatusBadRequest, "Dữ liệu không hợp lệ")
		return
	}

	tx := config.DB.Begin()
	var totalAmount float64

	val, _ := c.Get("user_id")
	var userID uint

	if idStr, ok := val.(string); ok {
		// Nếu là string, chuyển sang int rồi sang uint
		idInt, _ := strconv.Atoi(idStr)
		userID = uint(idInt)
	} else if idUint, ok := val.(uint); ok {
		// Nếu đã là uint thì dùng luôn
		userID = idUint
	}

	order := models.Order{
		UserID:          userID,
		ShippingAddress: input.ShippingAddress,
		PaymentMethod:   input.PaymentMethod,
		Status:          "pending",
	}

	if err := tx.Create(&order).Error; err != nil {
		tx.Rollback()
		utils.Error(c, 500, "Lỗi tạo đơn hàng")
		return
	}

	for _, item := range input.Items {
		var book models.Book
		if err := tx.First(&book, item.BookID).Error; err != nil {
			tx.Rollback()
			utils.Error(c, 404, "Không tìm thấy sách")
			return
		}

		orderItem := models.OrderItem{
			OrderID:         order.ID,
			BookID:          item.BookID,
			Quantity:        item.Quantity,
			PriceAtPurchase: book.Price,
		}
		tx.Create(&orderItem)
		totalAmount += book.Price * float64(item.Quantity)
	}

	tx.Model(&order).Update("TotalAmount", totalAmount)
	tx.Commit()

	var paymentURL string
	switch input.PaymentMethod {
	case "VNPAY":
		paymentURL = services.CreateVNPAYPaymentURL(order.ID, totalAmount, c.ClientIP())
	case "BANK_TRANSFER":
		paymentURL = fmt.Sprintf("/checkout/bank-info/%d", order.ID)
	case "MOMO":
		paymentURL = "MOMO_API_URL_HERE"
	default:
		paymentURL = "/orders/success"
	}

	utils.Success(c, http.StatusOK, gin.H{
		"order_id":    order.ID,
		"payment_url": paymentURL,
	}, "Đã khởi tạo đơn hàng")
}

// Cập nhật đơn hàng
func UpdateOrder(c *gin.Context) {
	var order models.Order
	id := c.Param("id")
	if err := config.DB.Preload("OrderItems").First(&order, id).Error; err != nil {
		utils.Error(c, http.StatusNotFound, "Order not found")
		return
	}
	var input struct {
		Status          string `json:"status"`
		ShippingAddress string `json:"shipping_address"`
	}
	if err := c.ShouldBindJSON(&input); err != nil {
		utils.Error(c, http.StatusBadRequest, err.Error())
		return
	}
	if input.Status != "" {
		order.Status = input.Status
	}
	if input.ShippingAddress != "" {
		order.ShippingAddress = input.ShippingAddress
	}
	if err := config.DB.Save(&order).Error; err != nil {
		utils.Error(c, http.StatusInternalServerError, err.Error())
		return
	}
	utils.Success(c, http.StatusOK, order, "Order updated successfully")
}

// Xóa đơn hàng
func DeleteOrder(c *gin.Context) {
	var order models.Order
	id := c.Param("id")
	if err := config.DB.First(&order, id).Error; err != nil {
		utils.Error(c, http.StatusNotFound, "Order not found")
		return
	}
	if err := config.DB.Delete(&order).Error; err != nil {
		utils.Error(c, http.StatusInternalServerError, err.Error())
		return
	}
	utils.Success(c, http.StatusOK, nil, "Order deleted successfully")
}

func VNPAY_IPN(c *gin.Context) { // not complte the flow
	// 1. Lấy thông tin từ VNPAY
	vnp_ResponseCode := c.Query("vnp_ResponseCode")
	vnp_TransactionStatus := c.Query("vnp_TransactionStatus")
	vnp_TxnRef := c.Query("vnp_TxnRef")         // OrderID
	vnp_TransactionNo := c.Query("vnp_TransactionNo") // txn_id

	// 2. Kiểm tra thanh toán thành công (Mã '00' cho cả Response và Status)
	if vnp_ResponseCode == "00" && vnp_TransactionStatus == "00" {
		
		// Bắt đầu Transaction để đảm bảo tính nhất quán dữ liệu
		tx := config.DB.Begin()

		// A. Cập nhật trạng thái đơn hàng
		var order models.Order
		if err := tx.First(&order, vnp_TxnRef).Error; err != nil {
			tx.Rollback()
			c.JSON(200, gin.H{"RspCode": "01", "Message": "Order not found"})
			return
		}

		// Kiểm tra nếu đơn hàng đã được xử lý trước đó (tránh xử lý trùng lặp)
		if order.Status == "paid" {
			tx.Rollback()
			c.JSON(200, gin.H{"RspCode": "02", "Message": "Order already confirmed"})
			return
		}

		// Cập nhật thông tin thanh toán
		updateData := map[string]interface{}{
			"Status":         "paid",
			"PaymentStatus": "paid",
			"TxnID":         vnp_TransactionNo,
		}
		if err := tx.Model(&order).Updates(updateData).Error; err != nil {
			tx.Rollback()
			c.JSON(200, gin.H{"RspCode": "99", "Message": "Update order failed"})
			return
		}

		// B. Trừ tồn kho trong bảng Books
		var orderItems []models.OrderItem
		tx.Where("order_id = ?", order.ID).Find(&orderItems)

		for _, item := range orderItems {
			// Sử dụng câu lệnh Update với điều kiện để trừ kho an toàn (tránh số âm)
			result := tx.Model(&models.Book{}).
				Where("id = ? AND stock_quantity >= ?", item.BookID, item.Quantity).
				Update("stock_quantity", gorm.Expr("stock_quantity - ?", item.Quantity))

			if result.Error != nil {
				tx.Rollback()
				c.JSON(200, gin.H{"RspCode": "99", "Message": "Inventory update error"})
				return
			}

			if result.RowsAffected == 0 {
				// Nếu không có dòng nào bị ảnh hưởng -> Hết hàng
				tx.Rollback()
				c.JSON(200, gin.H{"RspCode": "99", "Message": "Out of stock for Book ID " + fmt.Sprint(item.BookID)})
				return
			}
		}

		// Commit tất cả thay đổi
		tx.Commit()
	}

	// Trả về cho VNPAY
	c.Redirect(http.StatusOK, "https://localhost:5173/success/" + vnp_TxnRef)
}

