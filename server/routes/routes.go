package routes

import (
	"demowebgo/controllers"
	"demowebgo/routes/handle_routes"

	"github.com/gin-gonic/gin"
)

func SetupRoutes(router *gin.Engine) {
	userRoute := router.Group("/user")
	handle_routes.UserRoutes(userRoute)
}

func setupRouter() *gin.Engine {
	gin.SetMode(gin.TestMode)
	r := gin.Default()

	r.POST("/user", controllers.CreateUser)

	return r
}