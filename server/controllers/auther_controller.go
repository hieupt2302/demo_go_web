package controllers

import (
	"net/http"
	"demowebgo/config"
	"demowebgo/models"
	"demowebgo/utlis"
	"github.com/gin-gonic/gin"
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
	if err := config.DB.First(&author, id).Error; err != nil {
		utils.Error(c, http.StatusNotFound, "Author not found")
		return
	}
	var input struct {
		Name string `json:"name"`
		Bio  string `json:"bio"`
	}
	if err := c.ShouldBindJSON(&input); err != nil {
		utils.Error(c, http.StatusBadRequest, err.Error())
		return
	}
	author.Name = input.Name
	author.Bio = input.Bio
	if err := config.DB.Save(&author).Error; err != nil {
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
	if err := config.DB.Delete(&author).Error; err != nil {
		utils.Error(c, http.StatusInternalServerError, err.Error())
		return
	}
	utils.Success(c, http.StatusOK, nil, "Author deleted successfully")
}
