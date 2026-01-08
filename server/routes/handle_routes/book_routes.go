package handle_routes

import (
	"demowebgo/controllers"
	"demowebgo/middlewares"
	"github.com/gin-gonic/gin"
)

func BookRoutes(router *gin.RouterGroup) {
	router.POST("/create", middlewares.AuthMiddleware(), middlewares.AuthorizeJWT([]string{"admin"}),middlewares.UploadFileMiddleware("file",20 << 20), controllers.CreateBook)
	router.GET("/:id", controllers.GetBookById)
	router.PUT("/update/:id", middlewares.AuthMiddleware(), middlewares.AuthorizeJWT([]string{"admin"}), controllers.UpdateBook)
	router.DELETE("/delete/:id", middlewares.AuthMiddleware(), middlewares.AuthorizeJWT([]string{"admin"}), controllers.DeleteBook)
}