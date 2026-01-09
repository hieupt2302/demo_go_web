package handle_routes

import (
	"demowebgo/controllers"
	"demowebgo/middlewares"
	"github.com/gin-gonic/gin"
)

func UserRoutes(router *gin.RouterGroup) {
    router.GET("/:id", controllers.GetUserByID)
    router.PUT("/update/:id", controllers.UpdateUser)
    router.DELETE("/delete/:id", middlewares.AuthorizeJWT([]string{"admin"}), controllers.DeleteUser)
}
