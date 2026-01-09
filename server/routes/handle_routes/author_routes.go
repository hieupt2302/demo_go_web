package handle_routes
import (
	"github.com/gin-gonic/gin"
	"demowebgo/controllers"
)
func AuthorRoutes(router *gin.RouterGroup) {
	router.GET("/", controllers.GetAllAuthors)
	router.GET("/:id", controllers.GetAuthorByID)
}