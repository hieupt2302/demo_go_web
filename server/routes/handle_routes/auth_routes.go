package handle_routes
import (
	"github.com/gin-gonic/gin"
	"demowebgo/controllers"
)
func AuthRoutes(router *gin.RouterGroup) {
	router.POST("/login", controllers.Login)
	router.POST("/register", controllers.Register)
}