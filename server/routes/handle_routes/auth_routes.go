package handle_routes
import (
	"github.com/gin-gonic/gin"
	"demowebgo/controllers"
)
func AuthRoutes(router *gin.RouterGroup) {
	router.POST("/login", controllers.LoginUser)
	router.POST("/register", controllers.RegisterUser)
	router.POST("/refresh", controllers.RefreshToken)
}