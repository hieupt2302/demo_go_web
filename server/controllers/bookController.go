package controllers

import (
	"demowebgo/config"
	"demowebgo/dto"
	"demowebgo/models"
	"strconv"

	"github.com/gin-gonic/gin"
)

func CreateBook(c *gin.Context) {
	image, exist := c.Get("file_name")
	if !exist {
		c.JSON(400, gin.H{"error": "Image is required"})
		return
	}
	priceStr := c.PostForm("price")
	price, err := strconv.ParseFloat(priceStr, 64)
	if err != nil {
		c.JSON(400, gin.H{"error": "Invalid price"})
		return
	}
	bookDto := dto.BookDto{
		BookName: c.PostForm("book_name"),
		Image:    image.(string),
		Price:    price,
	}
	book := models.Book{
		BookName: c.PostForm("book_name"),
		Image:    image.(string),
		Price:    price,
	}

	if err := config.DB.Create(&book).Error; err != nil {
		c.JSON(500, gin.H{"error": err.Error()})
		return
	}
	c.JSON(201, bookDto)
}
