package controllers

import (
	"demowebgo/config"
	"demowebgo/models"
	"github.com/gin-gonic/gin"
       "net/http"
       "time"
       "demowebgo/utlis"
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
		UserID          uint    `json:"user_id" binding:"required"`
		TotalAmount     float64 `json:"total_amount" binding:"required"`
		Status          string  `json:"status"`
		ShippingAddress string  `json:"shipping_address"`
		OrderItems      []struct {
			BookID   uint    `json:"book_id" binding:"required"`
			Quantity int     `json:"quantity" binding:"required"`
			Price    float64 `json:"price" binding:"required"`
		} `json:"order_items" binding:"required"`
	}
	if err := c.ShouldBindJSON(&input); err != nil {
		utils.Error(c, http.StatusBadRequest, err.Error())
		return
	}
	order := models.Order{
		UserID:          input.UserID,
		OrderDate:       time.Now(),
		TotalAmount:     input.TotalAmount,
		Status:          input.Status,
		ShippingAddress: input.ShippingAddress,
	}
	for _, item := range input.OrderItems {
		order.OrderItems = append(order.OrderItems, models.OrderItem{
			BookID:          item.BookID,
			Quantity:        item.Quantity,
			PriceAtPurchase: item.Price,
		})
	}
	if err := config.DB.Create(&order).Error; err != nil {
		utils.Error(c, http.StatusInternalServerError, err.Error())
		return
	}
	utils.Success(c, http.StatusCreated, order, "Order created successfully")
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
