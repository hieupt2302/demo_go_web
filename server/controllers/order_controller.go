package controllers

import (
	"demowebgo/config"
	"demowebgo/models"
	"demowebgo/services"
	"demowebgo/utlis"
	"fmt"
	"log"
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

func VNPAY_Verify(c *gin.Context) {
	// 1. Lấy thông tin từ Query Params do Frontend gửi lên
	vnp_ResponseCode := c.Query("vnp_ResponseCode")
	vnp_TransactionStatus := c.Query("vnp_TransactionStatus")
	vnp_TxnRef := c.Query("vnp_TxnRef")         
	vnp_TransactionNo := c.Query("vnp_TransactionNo") 
	vnp_Amount := c.Query("vnp_Amount")

	log.Println("vnp_ResponseCode", vnp_ResponseCode)
	log.Println("vnp_TransactionStatus", vnp_TransactionStatus)
	log.Println("vnp_TxnRef", vnp_TxnRef)
	log.Println("vnp_TransactionNo", vnp_TransactionNo)
	log.Println("vnp_Amount", vnp_Amount)

	// Dữ liệu trả về cho Frontend hiển thị
	responseData := gin.H{
		"order_id": vnp_TxnRef,
		"amount":   vnp_Amount,
		"txn_no":   vnp_TransactionNo,
	}

	// 2. Kiểm tra mã phản hồi thanh toán
	if vnp_ResponseCode != "00" || vnp_TransactionStatus != "00" {
		utils.Error(c, http.StatusBadRequest, "Giao dịch thất bại hoặc đã bị hủy")
		return
	}

	// 3. Xử lý Database trong Transaction
	tx := config.DB.Begin()

	var order models.Order
	if err := tx.First(&order, vnp_TxnRef).Error; err != nil {
		tx.Rollback()
		utils.Error(c, http.StatusNotFound, "Không tìm thấy đơn hàng trên hệ thống")
		return
	}

	// Nếu đơn hàng đã được cập nhật 'paid' (tránh xử lý trùng lặp với IPN)
	if order.Status == "shipped" {
		tx.Rollback()
		utils.Success(c, http.StatusOK, responseData, "Đơn hàng đã được xác nhận thanh toán trước đó")
		return
	}

	// Cập nhật thông tin đơn hàng
	updateData := map[string]interface{}{
		"Status":         "shipped",
		"PaymentStatus": "paid",
		"TxnID":         vnp_TransactionNo,
	}

	if err := tx.Model(&order).Updates(updateData).Error; err != nil {
		tx.Rollback()
		utils.Error(c, http.StatusInternalServerError, "Lỗi cập nhật trạng thái đơn hàng")
		return
	}

	// 4. Trừ tồn kho sách
	var orderItems []models.OrderItem
	tx.Where("order_id = ?", order.ID).Find(&orderItems)

	for _, item := range orderItems {
		result := tx.Model(&models.Book{}).
			Where("id = ? AND stock_quantity >= ?", item.BookID, item.Quantity).
			Update("stock_quantity", gorm.Expr("stock_quantity - ?", item.Quantity))

		if result.Error != nil {
			tx.Rollback()
			utils.Error(c, http.StatusInternalServerError, "Lỗi hệ thống khi trừ kho")
			return
		}

		if result.RowsAffected == 0 {
			tx.Rollback()
			utils.Error(c, http.StatusConflict, fmt.Sprintf("Sách ID %d đã hết hàng", item.BookID))
			return
		}
	}

	tx.Commit()

	// 5. Trả về phản hồi thành công chuẩn ApiResponse
	utils.Success(c, http.StatusOK, responseData, "Thanh toán và cập nhật đơn hàng thành công")
}