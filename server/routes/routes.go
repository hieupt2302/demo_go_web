package routes

import (
	"demowebgo/routes/handle_routes"
	"github.com/gin-gonic/gin"
	"demowebgo/middlewares"
)

func SetupRoutes(router *gin.Engine) {
	authRoute := router.Group("/auth")
	handle_routes.AuthRoutes(authRoute)
	userRoute := router.Group("/user", middlewares.AuthMiddleware())
	handle_routes.UserRoutes(userRoute)
	cartRoute := router.Group("/cart", middlewares.AuthMiddleware())
    handle_routes.CartRoutes(cartRoute)
}