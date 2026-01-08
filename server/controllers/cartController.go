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
    userID, _ := strconv.Atoi(userIDStr)

    var existingCart models.Cart
    err := config.DB.Joins("JOIN cart_items ON cart_items.id = carts.cart_item_id").
        Where("carts.user_id = ? AND cart_items.book_id = ?", userID, input.BookID).
        Preload("CartItem").
        First(&existingCart).Error

    if err == nil {
        existingCart.CartItem.Quantity += input.Quantity
        config.DB.Save(&existingCart.CartItem)
        c.JSON(http.StatusOK, gin.H{"message": "Đã cập nhật số lượng trong giỏ hàng"})
        return
    }

    cartItem := models.CartItem{
        BookID:   input.BookID,
        Quantity: input.Quantity,
    }

    if err := config.DB.Create(&cartItem).Error; err != nil {
        c.JSON(http.StatusInternalServerError, gin.H{"error": "Không thể tạo mục giỏ hàng"})
        return
    }

    cartLink := models.Cart{
        UserID:     userID,
        CartItemID: cartItem.ID,
    }

    if err := config.DB.Create(&cartLink).Error; err != nil {
        c.JSON(http.StatusInternalServerError, gin.H{"error": "Không thể liên kết giỏ hàng"})
        return
    }

    c.JSON(http.StatusOK, gin.H{"message": "Đã thêm vào giỏ hàng thành công"})
}

func GetCart(c *gin.Context) {
    // 1. Lấy userID từ Middleware
    userIDStr := c.MustGet("user_id").(string)
    
    // 2. Khai báo một lát (slice) vì giỏ hàng có nhiều món
    var cartItems []models.Cart 

    // 3. Thực hiện truy vấn
    err := config.DB.
		Preload("User").            // Nạp thông tin User cho từng dòng
        Preload("CartItem.Book").   // Nạp thông tin món hàng và sách
        Where("user_id = ?", userIDStr). // Lọc đúng theo ID người dùng
        Find(&cartItems).Error      // Dùng Find để lấy danh sách
		
    if err != nil {
        c.JSON(http.StatusInternalServerError, gin.H{"error": "Lỗi truy vấn dữ liệu"})
        return
    }

    // 4. Kiểm tra nếu giỏ hàng trống
    if len(cartItems) == 0 {
        c.JSON(http.StatusOK, gin.H{
            "message": "Giỏ hàng hiện đang trống",
            "cart":    []interface{}{},
        })
        return
    }

    c.JSON(http.StatusOK, gin.H{
        "cart": cartItems,
    })
}