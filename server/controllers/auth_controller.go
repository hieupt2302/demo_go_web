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
	       	utils.Error(c, http.StatusBadRequest, err.Error())
	       	return
       	}
       	var user models.User
       	if err := config.DB.Where("email = ?", input.Email).First(&user).Error; err != nil {
	       	utils.Error(c, http.StatusUnauthorized, "Invalid email or password")
	       	return
       	}
       	if !utils.CheckPasswordHash(input.Password, user.PasswordHash) {
	       	utils.Error(c, http.StatusUnauthorized, "Invalid email or password")
	       	return
       	}
       	accessToken, refreshToken, err := utils.GenerateTokens(user.Email, strconv.Itoa(int(user.ID)), user.Role)
       	if err != nil {
	       	utils.Error(c, http.StatusInternalServerError, "Could not generate tokens")
	       	return
       	}
		loginResponse := gin.H{
			"access_token":  accessToken,
			"refresh_token": refreshToken,
		}
		utils.Success(c, http.StatusOK, loginResponse, "Login successful")
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
			utils.Error(c, http.StatusBadRequest, err.Error())
			return
		}
		var existingUser models.User
		if err := config.DB.Where("email = ?", input.Email).First(&existingUser).Error; err == nil {
			utils.Error(c, http.StatusBadRequest, "Email already exists")
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
			utils.Error(c, http.StatusInternalServerError, "Could not create user")
			return
		}

		utils.Success(c, http.StatusCreated, nil, "User registered successfully")
}

func RefreshToken(c *gin.Context) {
    var input struct {
        RefreshToken string `json:"refresh_token" binding:"required"`
    }

    // Sử dụng utlis.Error cho tính nhất quán
    if err := c.ShouldBindJSON(&input); err != nil {
        utils.Error(c, http.StatusBadRequest, "Cần cung cấp refresh token")
        return
    }

    // Xác thực Refresh Token qua helper jwt
    claims, err := utils.ValidateRefreshToken(input.RefreshToken)
    if err != nil {
        utils.Error(c, http.StatusUnauthorized, "Refresh token không hợp lệ hoặc hết hạn")
        return
    }

    // Tìm user trong DB để lấy thông tin mới nhất
    var user models.User
    if err := config.DB.First(&user, claims.UserID).Error; err != nil {
        utils.Error(c, http.StatusUnauthorized, "Người dùng không tồn tại")
        return
    }

    // Tạo cặp token mới (Rotation)
    newAT, newRT, err := utils.GenerateTokens(user.Email, strconv.Itoa(int(user.ID)), user.Role)
    if err != nil {
        utils.Error(c, http.StatusInternalServerError, "Lỗi tạo token")
        return
    }

    // CHUẨN HÓA: Bọc dữ liệu vào trường "data" để khớp với ApiResponse ở FE
    // Bao gồm cả user info để FE cập nhật lại store nếu cần
    refreshResponse := gin.H{
        "access_token":  newAT,
        "refresh_token": newRT,
    }

    utils.Success(c, http.StatusOK, refreshResponse, "Lấy mã xác thực mới thành công")
}
