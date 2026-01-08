package handle_routes
import (
	"github.com/gin-gonic/gin"
	"demowebgo/controllers"
)


func UserRoutes(router *gin.RouterGroup) {
    router.POST("/create", controllers.CreateUser)
    router.GET("/:id", controllers.GetUserByID)
    router.PUT("/update/:id", controllers.UpdateUser)
    router.DELETE("/delete/:id", controllers.DeleteUser)
}

