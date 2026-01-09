package controllers

import (
	"demowebgo/config"
	"demowebgo/models"

	"github.com/gin-gonic/gin"
)

func GetAuthorByID(c *gin.Context) {
	var author models.Author
	id := c.Param("id")
	if err := config.DB.Preload("Books").First(&author, id).Error; err != nil {
		c.JSON(404, gin.H{"error": "Author not found"})
		return
	}
	c.JSON(200, author)
}

func GetAllAuthors(c *gin.Context) {
	var authors []models.Author
	if err := config.DB.Preload("Books").Find(&authors).Error; err != nil {
		c.JSON(500, gin.H{"error": "Could not retrieve authors"})
		return
	}
	c.JSON(200, authors)
}
