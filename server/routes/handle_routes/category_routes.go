package handle_routes

import (
	"demowebgo/controllers"
	"demowebgo/middlewares"
	"github.com/gin-gonic/gin"
)

func CategoryRoutes(router *gin.RouterGroup) {
	router.GET("/", controllers.GetAllCategories)
	router.GET("/:id", controllers.GetCategoryByID)
	router.POST("/create", middlewares.AuthMiddleware(), middlewares.AuthorizeJWT([]string{"admin"}), controllers.CreateCategory)
	router.PUT("/update/:id", middlewares.AuthMiddleware(), middlewares.AuthorizeJWT([]string{"admin"}), controllers.UpdateCategory)
	router.DELETE("/delete/:id", middlewares.AuthMiddleware(), middlewares.AuthorizeJWT([]string{"admin"}), controllers.DeleteCategory)
}