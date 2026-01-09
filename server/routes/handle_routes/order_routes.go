package handle_routes

import (
	"demowebgo/controllers"
	"demowebgo/middlewares"
	"github.com/gin-gonic/gin"
)

func OrderRoutes(router *gin.RouterGroup) {
	router.POST("/create", middlewares.AuthorizeJWT([]string{"user"}), controllers.CreateOrder)
	router.GET("/:id", controllers.GetOrderByID)
	router.GET("/", middlewares.AuthorizeJWT([]string{"admin"}), controllers.GetAllOrders)
}