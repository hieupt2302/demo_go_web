package handle_routes
import (
	"github.com/gin-gonic/gin"
	"demowebgo/controllers"
	"demowebgo/middlewares"
)
func AuthorRoutes(router *gin.RouterGroup) {
	router.GET("/", controllers.GetAllAuthors)
	router.GET("/:id", controllers.GetAuthorByID)
	router.POST("/create", middlewares.AuthMiddleware(), middlewares.AuthorizeJWT([]string{"admin"}), controllers.CreateAuthor)
	router.PUT("/update/:id", middlewares.AuthMiddleware(), middlewares.AuthorizeJWT([]string{"admin"}), controllers.UpdateAuthor)
	router.DELETE("/delete/:id", middlewares.AuthMiddleware(), middlewares.AuthorizeJWT([]string{"admin"}), controllers.DeleteAuthor)
}