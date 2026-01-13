package controllers

import (
	"demowebgo/config"
	"demowebgo/models"
	"demowebgo/utlis"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
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
        // Sửa: Nhận mảng ID thay vì ID đơn lẻ
        CategoryIDs   []uint  `form:"category_ids"` 
        AuthorIDs     []uint  `form:"author_ids"`
    }

    if err := c.ShouldBind(&input); err != nil {
        utils.Error(c, http.StatusBadRequest, err.Error())
        return
    }

    var publishDate time.Time
    if input.PublishDate != "" {
        t, err := time.Parse("2006-01-02", input.PublishDate)
        if err == nil {
            publishDate = t
        }
    }

    coverImageURL := ""
    if path, exists := c.Get("file_path"); exists {
        coverImageURL = path.(string)
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
    }

    if len(input.CategoryIDs) > 0 {
        config.DB.Find(&book.Categories, input.CategoryIDs)
    }
    if len(input.AuthorIDs) > 0 {
        config.DB.Find(&book.Authors, input.AuthorIDs)
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
	if err := config.DB.Preload("Categories").Preload("Authors").Find(&books).Error; err != nil {
		utils.Error(c, http.StatusInternalServerError, err.Error())
		return
	}
	utils.Success(c, http.StatusOK, books, "Books retrieved successfully")
}

// Lấy chi tiết sách theo id
func GetBookByIDs(c *gin.Context) {
	var book models.Book
	id := c.Param("id")
	if err := config.DB.Preload("Categories").Preload("Authors").First(&book, id).Error; err != nil {
		utils.Error(c, http.StatusNotFound, "Book not found")
		return
	}
	utils.Success(c, http.StatusOK, book, "Book retrieved successfully")
}

// Cập nhật sách
func UpdateBook(c *gin.Context) {
	var book models.Book
	id := c.Param("id")
	if err := config.DB.Preload("Categories").Preload("Authors").First(&book, id).Error; err != nil {
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
		CategoryIDs   []uint  `form:"category_ids"`
        AuthorIDs     []uint  `form:"author_ids"`
	}
	if err := c.ShouldBind(&input); err != nil {
		utils.Error(c, http.StatusBadRequest, err.Error())
		return
	}

	err := config.DB.Transaction(func(tx *gorm.DB) error {
		// Cập nhật các trường cơ bản
		if input.Title != "" { book.Title = input.Title }
		if input.ISBN != "" { book.ISBN = input.ISBN }
		if input.Price != 0 { book.Price = input.Price }
		if input.StockQuantity != 0 { book.StockQuantity = input.StockQuantity }
		if input.Publisher != "" { book.Publisher = input.Publisher }
		if input.Description != "" { book.Description = input.Description }
		
		if input.PublishDate != "" {
			t, err := time.Parse("2006-01-02", input.PublishDate)
			if err == nil {
				book.PublishDate = t
			}
		}

		if filePath, exists := c.Get("file_path"); exists {
			book.CoverImageURL = filePath.(string)
		}

		// Lưu thông tin sách trước
		if err := tx.Save(&book).Error; err != nil {
			return err
		}

		// Cập nhật quan hệ Many-to-Many cho Categories
		if input.CategoryIDs != nil {
			var categories []models.Category
			tx.Find(&categories, input.CategoryIDs)
			if err := tx.Model(&book).Association("Categories").Replace(categories); err != nil {
				return err
			}
		}

		// Cập nhật quan hệ Many-to-Many cho Authors
		if input.AuthorIDs != nil {
			var authors []models.Author
			tx.Find(&authors, input.AuthorIDs)
			if err := tx.Model(&book).Association("Authors").Replace(authors); err != nil {
				return err
			}
		}

		return nil
	})

	if err != nil {
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

	err := config.DB.Transaction(func(tx *gorm.DB) error {
		// Clear các liên kết trong bảng trung gian
		tx.Model(&book).Association("Categories").Clear()
		tx.Model(&book).Association("Authors").Clear()
		
		// Xóa bản ghi sách
		if err := tx.Delete(&book).Error; err != nil {
			return err
		}
		return nil
	})

	if err != nil {
		utils.Error(c, http.StatusInternalServerError, err.Error())
		return
	}
	utils.Success(c, http.StatusOK, nil, "Book deleted successfully")
}

// server/controllers/book_controller.go

func AdvancedSearch(c *gin.Context) {
	searchType := c.Query("type")
	query := c.Query("q")
	
	if query == "" {
		utils.Error(c, http.StatusBadRequest, "Từ khóa tìm kiếm không được để trống")
		return
	}

	var bookIDs []uint
	switch searchType {
    case "author":
        config.DB.Model(&models.Book{}).
            Joins("JOIN book_authors ON book_authors.book_id = books.id").
            Joins("JOIN authors ON authors.id = book_authors.author_id").
            Where("MATCH(authors.name) AGAINST(? IN NATURAL LANGUAGE MODE)", query).
            Distinct().Pluck("books.id", &bookIDs)
            
    case "category":
        config.DB.Model(&models.Book{}).
            Joins("JOIN book_categories ON book_categories.book_id = books.id").
            Joins("JOIN categories ON categories.id = book_categories.category_id").
            Where("MATCH(categories.name) AGAINST(? IN NATURAL LANGUAGE MODE)", query).
            Distinct().Pluck("books.id", &bookIDs)
            
    case "book":
        config.DB.Model(&models.Book{}).
            Where("MATCH(title) AGAINST(? IN NATURAL LANGUAGE MODE)", query).
            Pluck("id", &bookIDs)
            
    default:
        utils.Error(c, http.StatusBadRequest, "Loại tìm kiếm không hợp lệ")
        return
    }

	if len(bookIDs) == 0 {
		utils.Success(c, http.StatusOK, []models.Book{}, "Không tìm thấy kết quả")
		return
	}

	var books []models.Book
	err := config.DB.Preload("Categories").Preload("Authors").
		Where("id IN ?", bookIDs).
		Find(&books).Error

	if err != nil {
		utils.Error(c, http.StatusInternalServerError, "Lỗi khi truy vấn dữ liệu")
		return
	}

	utils.Success(c, http.StatusOK, books, "Tìm thấy kết quả")
}
