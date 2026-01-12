package controllers

import (
	"demowebgo/config"
	"demowebgo/models"
	"github.com/gin-gonic/gin"
    "demowebgo/utlis"
	"net/http"
)


func GetUserByID(c *gin.Context) {
	var user models.User
	id := c.Param("id")

	if err := config.DB.Where("id = ?", id).First(&user).Error; err != nil {
		utils.Error(c, http.StatusNotFound, "User not found")
		return
	}

	// Chuẩn hóa: Trả về user trong trường "data"
	utils.Success(c, http.StatusOK, user, "Get user successfully")
}
func UpdateUser(c *gin.Context) {
	var user models.User
	id := c.Param("id")
	if err := config.DB.First(&user, id).Error; err != nil {
		utils.Error(c, http.StatusNotFound, "User not found")
		return
	}

	var input struct {
		Fullname string `json:"fullname"`
		Phone    string `json:"phone"`
		Address  string `json:"address"`
		Role     string `json:"role"`
	}
	
	if err := c.ShouldBindJSON(&input); err != nil {
		utils.Error(c, http.StatusBadRequest, err.Error())
		return
	}

	// Cập nhật các trường thông tin
	user.Fullname = input.Fullname
	user.Phone = input.Phone
	user.Address = input.Address
	user.Role = input.Role

	if err := config.DB.Save(&user).Error; err != nil {
		utils.Error(c, http.StatusInternalServerError, err.Error())
		return
	}

	// Chuẩn hóa: Trả về user đã cập nhật trong trường "data"
	utils.Success(c, http.StatusOK, user, "User updated successfully")
}

func DeleteUser(c *gin.Context) {
	var user models.User
	id := c.Param("id")

	if err := config.DB.First(&user, id).Error; err != nil {
		utils.Error(c, http.StatusNotFound, "User not found")
		return
	}

	if err := config.DB.Delete(&user).Error; err != nil {
		utils.Error(c, http.StatusInternalServerError, err.Error())
		return
	}

	// Chuẩn hóa: Data là nil khi xóa, sử dụng message để thông báo
	utils.Success(c, http.StatusOK, nil, "User deleted successfully")
}

