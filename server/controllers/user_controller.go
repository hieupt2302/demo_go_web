package controllers

import (
	"demowebgo/config"
	"demowebgo/models"
	"github.com/gin-gonic/gin"
)


func GetUserByID(c *gin.Context) {
       var user models.User
       id := c.Param("id")

       if err := config.DB.Where("id = ?", id).First(&user).Error; err != nil {
	       c.JSON(404, gin.H{"error": "User not found"})
	       return
       }

       c.JSON(200, user)
}
func UpdateUser(c *gin.Context) {
       var user models.User
       id := c.Param("id")
       if err := config.DB.First(&user, id).Error; err != nil {
	       c.JSON(404, gin.H{"error": "User not found"})
	       return
       }

       var input struct {
	       Fullname string `json:"fullname"`
	       Phone    string `json:"phone"`
	       Address  string `json:"address"`
	       Role     string `json:"role"`
       }
       if err := c.ShouldBindJSON(&input); err != nil {
	       c.JSON(400, gin.H{"error": err.Error()})
	       return
       }

       user.Fullname = input.Fullname
       user.Phone = input.Phone
       user.Address = input.Address
       user.Role = input.Role

       if err := config.DB.Save(&user).Error; err != nil {
	       c.JSON(500, gin.H{"error": err.Error()})
	       return
       }
       c.JSON(200, user)
}

func DeleteUser(c *gin.Context) {
       var user models.User
       id := c.Param("id")

       if err := config.DB.First(&user, id).Error; err != nil {
	       c.JSON(404, gin.H{"error": "User not found"})
	       return
       }

       if err := config.DB.Delete(&user).Error; err != nil {
	       c.JSON(500, gin.H{"error": err.Error()})
	       return
       }
       c.JSON(200, gin.H{"message": "User deleted successfully"})
}

