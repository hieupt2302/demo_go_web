package controllers

import (
	"demowebgo/config"
	"demowebgo/models"
	"demowebgo/utlis"
	"log"
	"net/http"
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
		utils.Error(c, http.StatusBadRequest, err.Error())
		return
	}
	log.Printf("Input Data Binded: %+v", input)
	// Handle file upload
	coverImageURL := ""
    if path, exists := c.Get("file_path"); exists {
        coverImageURL = path.(string)
    }

	var publishDate time.Time
	if input.PublishDate != "" {
		t, err := time.Parse("2006-01-02", input.PublishDate)
		if err != nil {
			utils.Error(c, http.StatusBadRequest, "Invalid publish date format. Use YYYY-MM-DD.")
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
		utils.Error(c, http.StatusInternalServerError, err.Error())
		return
	}
	utils.Success(c, http.StatusCreated, book, "Book created successfully")
}

// Lấy tất cả sách
func GetAllBooks(c *gin.Context) {
	var books []models.Book
	if err := config.DB.Preload("Category").Preload("Author").Find(&books).Error; err != nil {
		utils.Error(c, http.StatusInternalServerError, err.Error())
		return
	}
	utils.Success(c, http.StatusOK, books, "Books retrieved successfully")
}

// Lấy chi tiết sách theo id
func GetBookByIDs(c *gin.Context) {
	var book models.Book
	id := c.Param("id")
	if err := config.DB.Preload("Author").Preload("Category").First(&book, id).Error; err != nil {
		utils.Error(c, http.StatusNotFound, "Book not found")
		return
	}
	utils.Success(c, http.StatusOK, book, "Book retrieved successfully")
}

// Cập nhật sách
func UpdateBook(c *gin.Context) {
	var book models.Book
	id := c.Param("id")
	if err := config.DB.First(&book, id).Error; err != nil {
		utils.Error(c, http.StatusNotFound, "Book not found")
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
		utils.Error(c, http.StatusBadRequest, err.Error())
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
		utils.Error(c, http.StatusInternalServerError, err.Error())
		return
	}
	utils.Success(c, http.StatusOK, book, "Book updated successfully")
}

// Xóa sách
func DeleteBook(c *gin.Context) {
	var book models.Book
	id := c.Param("id")
	if err := config.DB.First(&book, id).Error; err != nil {
		utils.Error(c, http.StatusNotFound, "Book not found")
		return
	}
	if err := config.DB.Delete(&book).Error; err != nil {
		utils.Error(c, http.StatusInternalServerError, err.Error())
		return
	}
	utils.Success(c, http.StatusOK, nil, "Book deleted successfully")
}

// server/controllers/book_controller.go

func AdvancedSearch(c *gin.Context) {
    searchType := c.Query("type") // book, author, hoặc category
    query := c.Query("q")
    
    if query == "" {
        utils.Error(c, http.StatusBadRequest, "Từ khóa tìm kiếm không được để trống")
        return
    }

    var bookIDs []uint
    searchPattern := query + "%" // Tận dụng B-Tree Index (tìm từ đầu chuỗi)

    // BƯỚC 1: Tìm ID sách dựa trên lựa chọn của người dùng
    switch searchType {
    case "author":
        // Tìm ID sách từ tên tác giả
        config.DB.Model(&models.Book{}).
            Joins("JOIN authors ON authors.id = books.author_id").
            Where("authors.name LIKE ?", searchPattern).
            Pluck("books.id", &bookIDs)
            
    case "category":
        // Tìm ID sách từ tên thể loại
        config.DB.Model(&models.Book{}).
            Joins("JOIN categories ON categories.id = books.category_id").
            Where("categories.name LIKE ?", searchPattern).
            Pluck("books.id", &bookIDs)
            
    case "book":
        // Tìm ID sách trực tiếp từ tiêu đề sách
        config.DB.Model(&models.Book{}).
            Where("title LIKE ?", searchPattern).
            Pluck("id", &bookIDs)
            
    default:
        utils.Error(c, http.StatusBadRequest, "Loại tìm kiếm không hợp lệ")
        return
    }

    // BƯỚC 2: Nếu không tìm thấy ID nào, trả về mảng rỗng ngay lập tức
    if len(bookIDs) == 0 {
        utils.Success(c, http.StatusOK, []models.Book{}, "Không tìm thấy kết quả")
        return
    }

    // BƯỚC 3: Join 3 bảng để lấy thông tin chi tiết dựa trên danh sách ID đã tìm được
    // Thao tác này cực nhanh vì lọc theo Primary Key (ID IN ?)
    var books []models.Book
    err := config.DB.Preload("Category").Preload("Author").
        Where("id IN ?", bookIDs).
        Find(&books).Error

    if err != nil {
        utils.Error(c, http.StatusInternalServerError, "Lỗi khi truy vấn dữ liệu")
        return
    }

    utils.Success(c, http.StatusOK, books, "Tìm thấy kết quả")
}
