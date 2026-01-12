package handle_routes

import (
	"demowebgo/controllers"
	"demowebgo/middlewares"
	"github.com/gin-gonic/gin"
)

func OrderRoutes(router *gin.RouterGroup) {
	router.POST("/create", controllers.CreateOrder)
	router.GET("/:id", controllers.GetOrderByID)
	router.GET("/", middlewares.AuthorizeJWT([]string{"admin"}), controllers.GetAllOrders)
	router.GET("/user/:userId", controllers.GetOrdersByUserID)
	router.PUT("/:id", controllers.UpdateOrder)
	router.DELETE("/:id", controllers.DeleteOrder)
}