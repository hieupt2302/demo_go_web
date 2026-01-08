package handle_routes
import (
	"github.com/gin-gonic/gin"
	"demowebgo/controllers"
    "demowebgo/middlewares"
)


func UserRoutes(router *gin.RouterGroup) {
    router.GET("/:id", middlewares.AuthMiddleware(), controllers.GetUserByID)
    router.PUT("/update/:id", middlewares.AuthMiddleware(), controllers.UpdateUser)
    router.DELETE("/delete/:id", middlewares.AuthMiddleware(), middlewares.AdminMiddleware(), controllers.DeleteUser)
}

