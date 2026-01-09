package controllers

import (
	"demowebgo/config"
	"demowebgo/models"
	"demowebgo/utlis"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

func LoginUser(c *gin.Context) {
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
       if !utils.CheckPasswordHash(input.Password, user.PasswordHash) {
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

func RegisterUser(c *gin.Context) {
       var input struct {
	       Fullname    string `json:"fullname" binding:"required"`
	       Email       string `json:"email" binding:"required,email"`
	       Password    string `json:"password" binding:"required,min=6"`
	       Phone       string `json:"phone"`
	       Address     string `json:"address"`
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
	       Fullname:     input.Fullname,
	       Email:        input.Email,
	       PasswordHash: hashedPassword,
	       Phone:        input.Phone,
	       Address:      input.Address,
	       Role:         input.Role,
       }

       if err := config.DB.Create(&user).Error; err != nil {
	       c.JSON(http.StatusInternalServerError, gin.H{"error": "Không thể tạo tài khoản"})
	       return
       }

       c.JSON(http.StatusCreated, gin.H{"message": "Đăng ký tài khoản thành công"})
}

func RefreshToken(c *gin.Context) {
	var input struct {
		RefreshToken string `json:"refresh_token" binding:"required"`
	}

	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Cần cung cấp refresh token"})
		return
	}

	// Xác thực Refresh Token
	claims, err := utils.ValidateRefreshToken(input.RefreshToken)
	if err != nil {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "Refresh token không hợp lệ hoặc hết hạn"})
		return
	}

	// Tìm user trong DB để lấy thông tin mới nhất
	var user models.User
	if err := config.DB.First(&user, claims.UserID).Error; err != nil {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "Người dùng không tồn tại"})
		return
	}

	// Tạo cặp token mới (Rotation)
	newAT, newRT, err := utils.GenerateTokens(user.Email, strconv.Itoa(int(user.ID)), user.Role)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Lỗi tạo token"})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"access_token":  newAT,
		"refresh_token": newRT,
	})
}
