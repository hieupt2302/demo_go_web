package controllers

import (
	"demowebgo/config"
	"demowebgo/models"
	"demowebgo/utlis"
	"net/http"
	"strconv"
	"github.com/gin-gonic/gin"
)

func Login(c *gin.Context) {
	var input struct {
		Email    string `json:"email" binding:"required"`
		Password string `json:"password" binding:"required"`
	}
	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	var user models.User
	if err := config.DB.Where("email = ?", input.Email).First(&user).Error; err != nil {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "Invalid email or password"})
		return
	}
	if !utils.CheckPasswordHash(input.Password, user.Password) {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "Invalid email or password"})
		return
	}
	accessToken, refreshToken, err := utils.GenerateTokens(user.Email, strconv.Itoa(int(user.ID)), user.Role)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Could not generate tokens"})
		return
	}
	
	c.JSON(http.StatusOK, gin.H{
		"access_token":  accessToken,
		"refresh_token": refreshToken,
	})
}

func Register(c *gin.Context) {
	var input struct {
		Email       string `json:"email" binding:"required,email"`
		Password 	string `json:"password" binding:"required,min=6"`
		PhoneNumber string `json:"phone_number"`
		Role        string `json:"role"`
	}
	if err := c.ShouldBindJSON(&input); err != nil {
        c.JSON(http.StatusBadRequest, gin.H{"error": "Dữ liệu không hợp lệ: " + err.Error()})
        return
    }
	var existingUser models.User
    if err := config.DB.Where("email = ?", input.Email).First(&existingUser).Error; err == nil {
        c.JSON(http.StatusConflict, gin.H{"error": "Email này đã được sử dụng"})
        return
    }

	hashedPassword := utils.HashPassword(&input.Password)
	user := models.User{
		Email:       input.Email,
        Password:    *hashedPassword,
        PhoneNumber: input.PhoneNumber,
        Role:        input.Role,
    }

	if err := config.DB.Create(&user).Error; err != nil {
        c.JSON(http.StatusInternalServerError, gin.H{"error": "Không thể tạo tài khoản"})
        return
    }

    c.JSON(http.StatusCreated, gin.H{"message": "Đăng ký tài khoản thành công"})
}