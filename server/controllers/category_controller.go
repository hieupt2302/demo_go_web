package controllers

import (
	"demowebgo/config"
	"demowebgo/models"
	"demowebgo/utlis"
	"net/http"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

func GetAllCategories(c *gin.Context) {
       var categories []models.Category
       if err := config.DB.Preload("Books").Find(&categories).Error; err != nil {
	       utils.Error(c, http.StatusInternalServerError, err.Error())
	       return
       }
       utils.Success(c, http.StatusOK, categories, "Get categories successfully")
}


func GetCategoryByID(c *gin.Context) {
       var category models.Category
       id := c.Param("id")
       if err := config.DB.Preload("Books.Authors").First(&category, id).Error; err != nil {
	       utils.Error(c, http.StatusNotFound, "Category not found")
	       return
       }
       utils.Success(c, http.StatusOK, category, "Get category successfully")
}


func CreateCategory(c *gin.Context) {
       var input struct {
	       Name        string `json:"name" binding:"required"`
	       Description string `json:"description"`
       }
       if err := c.ShouldBindJSON(&input); err != nil {
	       utils.Error(c, http.StatusBadRequest, err.Error())
	       return
       }
       category := models.Category{
	       Name:        input.Name,
	       Description: input.Description,
       }
       if err := config.DB.Create(&category).Error; err != nil {
	       utils.Error(c, http.StatusInternalServerError, err.Error())
	       return
       }
       utils.Success(c, http.StatusCreated, category, "Category created successfully")
}

func UpdateCategory(c *gin.Context) {
       var category models.Category
       id := c.Param("id")
       if err := config.DB.First(&category, id).Error; err != nil {
	       utils.Error(c, http.StatusNotFound, "Category not found")
	       return
       }
       var input struct {
	       Name        string `json:"name"`
	       Description string `json:"description"`
              BookIDs     []uint `json:"book_ids"`
       }
       if err := c.ShouldBindJSON(&input); err != nil {
	       utils.Error(c, http.StatusBadRequest, err.Error())
	       return
       }
       err := config.DB.Transaction(func(tx *gorm.DB) error {
		// Cập nhật thông tin cơ bản
		if input.Name != "" {
			category.Name = input.Name
		}
		category.Description = input.Description

		if err := tx.Save(&category).Error; err != nil {
			return err
		}

		// Cập nhật quan hệ Many-to-Many với Books
		if input.BookIDs != nil {
			var books []models.Book
			if len(input.BookIDs) > 0 {
				tx.Find(&books, input.BookIDs)
			}
			// Replace xóa các liên kết cũ và thêm mới vào bảng book_categories
			if err := tx.Model(&category).Association("Books").Replace(books); err != nil {
				return err
			}
		}
		return nil
	})

	if err != nil {
		utils.Error(c, http.StatusInternalServerError, err.Error())
		return
	}

	utils.Success(c, http.StatusOK, category, "Category updated successfully")
}

func DeleteCategory(c *gin.Context) {
	var category models.Category
	id := c.Param("id")

	if err := config.DB.First(&category, id).Error; err != nil {
		utils.Error(c, http.StatusNotFound, "Category not found")
		return
	}

	err := config.DB.Transaction(func(tx *gorm.DB) error {
		if err := tx.Model(&category).Association("Books").Clear(); err != nil {
			return err
		}
		if err := tx.Delete(&category).Error; err != nil {
			return err
		}
		return nil
	})

	if err != nil {
		utils.Error(c, http.StatusInternalServerError, err.Error())
		return
	}

	utils.Success(c, http.StatusOK, nil, "Category deleted successfully")
}