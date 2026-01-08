package controllers
import (
	"github.com/gin-gonic/gin"
	"demowebgo/models"
	"demowebgo/config"
)

func CreateUser(c *gin.Context) {
	var user models.User
	if err := c.ShouldBindJSON(&user); err != nil {
		c.JSON(400, gin.H{"error": err.Error()})
		return
	}
	if err := config.DB.Create(&user).Error; err != nil {
		c.JSON(500, gin.H{"error": err.Error()})
		return
	}
	c.JSON(201, user)
}

func GetUserByID(c *gin.Context) {
	var user models.User
	id := c.Param("id")
	if err := config.DB.First(&user, id).Error; err != nil {
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
	if err := c.ShouldBindJSON(&user); err != nil {
		c.JSON(400, gin.H{"error": err.Error()})
		return
	}
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