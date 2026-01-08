package routes

import (
	"demowebgo/routes/handle_routes"

	"github.com/gin-gonic/gin"
)

func SetupRoutes(router *gin.Engine) {
	userRoute := router.Group("/user")
	handle_routes.UserRoutes(userRoute)
	authRoute := router.Group("/auth")
	handle_routes.AuthRoutes(authRoute)
	bookRoute := router.Group("/book")
	handle_routes.BookRoutes(bookRoute)
}