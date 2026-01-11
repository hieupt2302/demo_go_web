package routes

import (
	"demowebgo/routes/handle_routes"
	"github.com/gin-gonic/gin"
	"demowebgo/middlewares"
)

func SetupRoutes(router *gin.Engine) {
	authRoute := router.Group("/auth")
	handle_routes.AuthRoutes(authRoute)
	bookRoute := router.Group("/book")
	handle_routes.BookRoutes(bookRoute)
	userRoute := router.Group("/user", middlewares.AuthMiddleware())
	handle_routes.UserRoutes(userRoute)
	authorRoute := router.Group("/author", middlewares.AuthMiddleware())
	handle_routes.AuthorRoutes(authorRoute)
	categoryRoute := router.Group("/category", middlewares.AuthMiddleware())
	handle_routes.CategoryRoutes(categoryRoute)
	orderRoute := router.Group("/order", middlewares.AuthMiddleware())
	handle_routes.OrderRoutes(orderRoute)
}