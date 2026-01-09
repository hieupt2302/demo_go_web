package controllers

import (
       "demowebgo/config"
       "demowebgo/models"
       "github.com/gin-gonic/gin"
       "time"
)

// Lấy tất cả đơn hàng
func GetAllOrders(c *gin.Context) {
       var orders []models.Order
       if err := config.DB.Preload("OrderItems.Book").Find(&orders).Error; err != nil {
	       c.JSON(500, gin.H{"error": err.Error()})
	       return
       }
       c.JSON(200, orders)
}

// Lấy đơn hàng theo id
func GetOrderByID(c *gin.Context) {
       var order models.Order
       id := c.Param("id")
       if err := config.DB.Preload("OrderItems.Book").First(&order, id).Error; err != nil {
	       c.JSON(404, gin.H{"error": "Order not found"})
	       return
       }
       c.JSON(200, order)
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
	       c.JSON(400, gin.H{"error": err.Error()})
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
	       c.JSON(500, gin.H{"error": err.Error()})
	       return
       }
       c.JSON(201, order)
}

// Cập nhật đơn hàng
func UpdateOrder(c *gin.Context) {
       var order models.Order
       id := c.Param("id")
       if err := config.DB.Preload("OrderItems").First(&order, id).Error; err != nil {
	       c.JSON(404, gin.H{"error": "Order not found"})
	       return
       }
       var input struct {
	       Status          string `json:"status"`
	       ShippingAddress string `json:"shipping_address"`
       }
       if err := c.ShouldBindJSON(&input); err != nil {
	       c.JSON(400, gin.H{"error": err.Error()})
	       return
       }
       if input.Status != "" {
	       order.Status = input.Status
       }
       if input.ShippingAddress != "" {
	       order.ShippingAddress = input.ShippingAddress
       }
       if err := config.DB.Save(&order).Error; err != nil {
	       c.JSON(500, gin.H{"error": err.Error()})
	       return
       }
       c.JSON(200, order)
}

// Xóa đơn hàng
func DeleteOrder(c *gin.Context) {
       var order models.Order
       id := c.Param("id")
       if err := config.DB.First(&order, id).Error; err != nil {
	       c.JSON(404, gin.H{"error": "Order not found"})
	       return
       }
       if err := config.DB.Delete(&order).Error; err != nil {
	       c.JSON(500, gin.H{"error": err.Error()})
	       return
       }
       c.JSON(200, gin.H{"message": "Order deleted successfully"})
}
