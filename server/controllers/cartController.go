package controllers

import (
    "net/http"
    "demowebgo/models"
    "demowebgo/config"
    "github.com/gin-gonic/gin"
    "strconv"
)

func AddToCart(c *gin.Context) {
    var input struct {
        BookID   int `json:"book_id" binding:"required"`
        Quantity int `json:"quantity" binding:"required"`
    }

    if err := c.ShouldBindJSON(&input); err != nil {
        c.JSON(http.StatusBadRequest, gin.H{"error": "Dữ liệu không hợp lệ"})
        return
    }

    userIDStr := c.MustGet("user_id").(string)
    userID, _ := strconv.ParseUint(userIDStr, 10, 32)

    cartItem := models.CartItem{
        BookID:   input.BookID,
        Quantity: input.Quantity,
    }

    if err := config.DB.Create(&cartItem).Error; err != nil {
        c.JSON(http.StatusInternalServerError, gin.H{"error": "Không thể tạo mục giỏ hàng"})
        return
    }

    cartLink := models.Cart{
        UserID:     uint(userID),
        CartItemID: uint(cartItem.ID),
    }

    if err := config.DB.Create(&cartLink).Error; err != nil {
        c.JSON(http.StatusInternalServerError, gin.H{"error": "Không thể liên kết giỏ hàng"})
        return
    }

    c.JSON(http.StatusOK, gin.H{"message": "Đã thêm vào giỏ hàng thành công"})
}

func GetCart(c *gin.Context) {
    userIDStr := c.MustGet("user_id").(string)
    
    var user models.User
    err := config.DB.Preload("CartItem.Book").First(&user, userIDStr).Error

    if err != nil {
        c.JSON(http.StatusNotFound, gin.H{"error": "Giỏ hàng trống"})
        return
    }

    c.JSON(http.StatusOK, gin.H{
        "cart": user.CartItem,
    })
}