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

func GetBookById(c *gin.Context) {
	var book models.Book
	id := c.Param("id")

	if err := config.DB.Where("is_deleted = ?", false).First(&book, id).Error; err != nil {
		c.JSON(404, gin.H{"error": "Book not found"})
		return
	}

	c.JSON(200, book)
}
func UpdateBook(c *gin.Context) {
	var book models.Book
	id := c.Param("id")
	if err := config.DB.Where("is_deleted = ?", false).First(&book, id).Error; err != nil {
		c.JSON(404, gin.H{"error": "Book not found"})
		return
	}
	if err := c.ShouldBindJSON(&book); err != nil {
		c.JSON(400, gin.H{"error": err.Error()})
		return
	}
	if err := config.DB.Save(&book).Error; err != nil {
		c.JSON(500, gin.H{"error": err.Error()})
		return
	}	
	c.JSON(200, book)
}
func DeleteBook(c *gin.Context) {
	var book models.Book
	id := c.Param("id")	
	if err := config.DB.Where("is_deleted = ?", false).First(&book, id).Error; err != nil {
		c.JSON(404, gin.H{"error": "Book not found"})
		return
	}
	book.IsDeleted = true
	if err := config.DB.Save(&book).Error; err != nil {
		c.JSON(500, gin.H{"error": err.Error()})
		return
	}	
	c.JSON(200, gin.H{"message": "Book deleted successfully"})
}