package controllers

import (
	"demowebgo/config"
	"demowebgo/models"
	"github.com/gin-gonic/gin"
)

func GetAllCategories(c *gin.Context) {
       var categories []models.Category
       if err := config.DB.Preload("Books").Find(&categories).Error; err != nil {
	       c.JSON(500, gin.H{"error": err.Error()})
	       return
       }
       c.JSON(200, categories)
}


func GetCategoryByID(c *gin.Context) {
       var category models.Category
       id := c.Param("id")
       if err := config.DB.Preload("Books").First(&category, id).Error; err != nil {
	       c.JSON(404, gin.H{"error": "Category not found"})
	       return
       }
       c.JSON(200, category)
}


func CreateCategory(c *gin.Context) {
       var input struct {
	       Name        string `json:"name" binding:"required"`
	       Description string `json:"description"`
       }
       if err := c.ShouldBindJSON(&input); err != nil {
	       c.JSON(400, gin.H{"error": err.Error()})
	       return
       }
       category := models.Category{
	       Name:        input.Name,
	       Description: input.Description,
       }
       if err := config.DB.Create(&category).Error; err != nil {
	       c.JSON(500, gin.H{"error": err.Error()})
	       return
       }
       c.JSON(201, category)
}

func UpdateCategory(c *gin.Context) {
       var category models.Category
       id := c.Param("id")
       if err := config.DB.First(&category, id).Error; err != nil {
	       c.JSON(404, gin.H{"error": "Category not found"})
	       return
       }
       var input struct {
	       Name        string `json:"name"`
	       Description string `json:"description"`
       }
       if err := c.ShouldBindJSON(&input); err != nil {
	       c.JSON(400, gin.H{"error": err.Error()})
	       return
       }
       if input.Name != "" {
	       category.Name = input.Name
       }
       if input.Description != "" {
	       category.Description = input.Description
       }
       if err := config.DB.Save(&category).Error; err != nil {
	       c.JSON(500, gin.H{"error": err.Error()})
	       return
       }
       c.JSON(200, category)
}

func DeleteCategory(c *gin.Context) {
       var category models.Category
       id := c.Param("id")
       if err := config.DB.First(&category, id).Error; err != nil {
	       c.JSON(404, gin.H{"error": "Category not found"})
	       return
       }
       if err := config.DB.Delete(&category).Error; err != nil {
	       c.JSON(500, gin.H{"error": err.Error()})
	       return
       }
       c.JSON(200, gin.H{"message": "Category deleted successfully"})
}