package controllers

import (
	"demowebgo/config"
	"demowebgo/models"
	"github.com/gin-gonic/gin"
       "demowebgo/utlis"
       "net/http"
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
       if err := config.DB.Preload("Books").First(&category, id).Error; err != nil {
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
       }
       if err := c.ShouldBindJSON(&input); err != nil {
	       utils.Error(c, http.StatusBadRequest, err.Error())
	       return
       }
       if input.Name != "" {
	       category.Name = input.Name
       }
       if input.Description != "" {
	       category.Description = input.Description
       }
       if err := config.DB.Save(&category).Error; err != nil {
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
       if err := config.DB.Delete(&category).Error; err != nil {
	       utils.Error(c, http.StatusInternalServerError, err.Error())
	       return
       }
       utils.Success(c, http.StatusOK, nil, "Category deleted successfully")
}