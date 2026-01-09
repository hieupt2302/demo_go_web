
package controllers

import (
	"demowebgo/config"
	"demowebgo/models"
	"time"

	"github.com/gin-gonic/gin"
)

func CreateBook(c *gin.Context) {
	var input struct {
		Title         string  `form:"title" binding:"required"`
		ISBN          string  `form:"isbn"`
		Price         float64 `form:"price" binding:"required"`
		StockQuantity int     `form:"stock_quantity"`
		Publisher     string  `form:"publisher"`
		PublishDate   string  `form:"publish_date"`
		Description   string  `form:"description"`
		CategoryID    uint    `form:"category_id"`
		AuthorID      uint    `form:"author_id"`
	}
	if err := c.ShouldBind(&input); err != nil {
		c.JSON(400, gin.H{"error": err.Error()})
		return
	}

	// Handle file upload
	file, err := c.FormFile("cover_image")
	var coverImageURL string
	if err == nil {
		// Save file to uploads folder
		uploadPath := "uploads/" + file.Filename
		if err := c.SaveUploadedFile(file, uploadPath); err != nil {
			c.JSON(500, gin.H{"error": "Failed to save image"})
			return
		}
		coverImageURL = uploadPath
	} else {
		coverImageURL = ""
	}

	var publishDate time.Time
	if input.PublishDate != "" {
		t, err := time.Parse("2006-01-02", input.PublishDate)
		if err != nil {
			c.JSON(400, gin.H{"error": "Invalid publish date format. Use YYYY-MM-DD."})
			return
		}
		publishDate = t
	}

	book := models.Book{
		Title:         input.Title,
		ISBN:          input.ISBN,
		Price:         input.Price,
		StockQuantity: input.StockQuantity,
		Publisher:     input.Publisher,
		PublishDate:   publishDate,
		CoverImageURL: coverImageURL,
		Description:   input.Description,
		CategoryID:    input.CategoryID,
		AuthorID:      input.AuthorID,
	}

	if err := config.DB.Create(&book).Error; err != nil {
		c.JSON(500, gin.H{"error": err.Error()})
		return
	}
	c.JSON(201, book)
}

// Lấy tất cả sách
func GetAllBooks(c *gin.Context) {
	var books []models.Book
	if err := config.DB.Find(&books).Error; err != nil {
		c.JSON(500, gin.H{"error": err.Error()})
		return
	}
	c.JSON(200, books)
}

// Lấy chi tiết sách theo id
func GetBookByID(c *gin.Context) {
	var book models.Book
	id := c.Param("id")
	if err := config.DB.First(&book, id).Error; err != nil {
		c.JSON(404, gin.H{"error": "Book not found"})
		return
	}
	c.JSON(200, book)
}

// Cập nhật sách
func UpdateBook(c *gin.Context) {
	var book models.Book
	id := c.Param("id")
	if err := config.DB.First(&book, id).Error; err != nil {
		c.JSON(404, gin.H{"error": "Book not found"})
		return
	}

	var input struct {
		Title         string  `form:"title"`
		ISBN          string  `form:"isbn"`
		Price         float64 `form:"price"`
		StockQuantity int     `form:"stock_quantity"`
		Publisher     string  `form:"publisher"`
		PublishDate   string  `form:"publish_date"`
		Description   string  `form:"description"`
		CategoryID    uint    `form:"category_id"`
		AuthorID      uint    `form:"author_id"`
	}
	if err := c.ShouldBind(&input); err != nil {
		c.JSON(400, gin.H{"error": err.Error()})
		return
	}

	if input.Title != "" {
		book.Title = input.Title
	}
	if input.ISBN != "" {
		book.ISBN = input.ISBN
	}
	if input.Price != 0 {
		book.Price = input.Price
	}
	if input.StockQuantity != 0 {
		book.StockQuantity = input.StockQuantity
	}
	if input.Publisher != "" {
		book.Publisher = input.Publisher
	}
	if input.PublishDate != "" {
		t, err := time.Parse("2006-01-02", input.PublishDate)
		if err == nil {
			book.PublishDate = t
		}
	}
	if input.Description != "" {
		book.Description = input.Description
	}
	if input.CategoryID != 0 {
		book.CategoryID = input.CategoryID
	}
	if input.AuthorID != 0 {
		book.AuthorID = input.AuthorID
	}

	// Lấy ảnh từ middleware nếu có
	if filePath, exists := c.Get("file_path"); exists {
		book.CoverImageURL = filePath.(string)
	}

	if err := config.DB.Save(&book).Error; err != nil {
		c.JSON(500, gin.H{"error": err.Error()})
		return
	}
	c.JSON(200, book)
}

// Xóa sách
func DeleteBook(c *gin.Context) {
	var book models.Book
	id := c.Param("id")
	if err := config.DB.First(&book, id).Error; err != nil {
		c.JSON(404, gin.H{"error": "Book not found"})
		return
	}
	if err := config.DB.Delete(&book).Error; err != nil {
		c.JSON(500, gin.H{"error": err.Error()})
		return
	}
	c.JSON(200, gin.H{"message": "Book deleted successfully"})
}

// Lấy sách theo tên (title)
func GetBookByName(c *gin.Context) {
       name := c.Query("name")
       if name == "" {
	       c.JSON(400, gin.H{"error": "Missing name query param"})
	       return
       }
       var books []models.Book
       if err := config.DB.Where("title LIKE ?", "%"+name+"%").Find(&books).Error; err != nil {
	       c.JSON(500, gin.H{"error": err.Error()})
	       return
       }
       c.JSON(200, books)
}