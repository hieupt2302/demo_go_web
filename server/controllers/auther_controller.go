package controllers

import (
	"demowebgo/config"
	"demowebgo/models"
	"demowebgo/utlis"
	"net/http"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

func GetAuthorByID(c *gin.Context) {
	var author models.Author
	id := c.Param("id")
	if err := config.DB.Preload("Books").First(&author, id).Error; err != nil {
		utils.Error(c, http.StatusNotFound, "Author not found")
		return
	}
	utils.Success(c, http.StatusOK, author, "Get author successfully")
}

func GetAllAuthors(c *gin.Context) {
	var authors []models.Author
	if err := config.DB.Preload("Books").Find(&authors).Error; err != nil {
		utils.Error(c, http.StatusInternalServerError, err.Error())
		return
	}
	utils.Success(c, http.StatusOK, authors, "Get authors successfully")
}

func CreateAuthor(c *gin.Context) {
	var input struct {
		Name string `json:"name" binding:"required"`
		Bio  string `json:"bio"`
	}
	if err := c.ShouldBindJSON(&input); err != nil {
		utils.Error(c, http.StatusBadRequest, err.Error())
		return
	}
	author := models.Author{
		Name: input.Name,
		Bio:  input.Bio,
	}
	if err := config.DB.Create(&author).Error; err != nil {
		utils.Error(c, http.StatusInternalServerError, err.Error())
		return
	}
	utils.Success(c, http.StatusCreated, author, "Author created successfully")
}

func UpdateAuthor(c *gin.Context) {
    var author models.Author
    id := c.Param("id")
    
    // Tìm tác giả hiện tại
    if err := config.DB.First(&author, id).Error; err != nil {
        utils.Error(c, http.StatusNotFound, "Author not found")
        return
    }

    var input struct {
        Name    string `json:"name"`
        Bio     string `json:"bio"`
        BookIDs []uint `json:"book_ids"`
    }

    if err := c.ShouldBindJSON(&input); err != nil {
        utils.Error(c, http.StatusBadRequest, err.Error())
        return
    }

    author.Name = input.Name
    author.Bio = input.Bio

    err := config.DB.Transaction(func(tx *gorm.DB) error {
        if err := tx.Save(&author).Error; err != nil {
            return err
        }

        if input.BookIDs != nil {
            var books []models.Book
            if len(input.BookIDs) > 0 {
                tx.Find(&books, input.BookIDs)
            }
            if err := tx.Model(&author).Association("Books").Replace(books); err != nil {
                return err
            }
        }
        return nil
    })

    if err != nil {
        utils.Error(c, http.StatusInternalServerError, err.Error())
        return
    }

    utils.Success(c, http.StatusOK, author, "Author updated successfully")
}

func DeleteAuthor(c *gin.Context) {
    var author models.Author
    id := c.Param("id")

    if err := config.DB.First(&author, id).Error; err != nil {
        utils.Error(c, http.StatusNotFound, "Author not found")
        return
    }

    err := config.DB.Transaction(func(tx *gorm.DB) error {
        if err := tx.Model(&author).Association("Books").Clear(); err != nil {
            return err
        }

        if err := tx.Delete(&author).Error; err != nil {
            return err
        }

        return nil
    })

    if err != nil {
        utils.Error(c, http.StatusInternalServerError, err.Error())
        return
    }

    utils.Success(c, http.StatusOK, nil, "Author deleted successfully")
}
