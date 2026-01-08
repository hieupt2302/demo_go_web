package handle_routes

import (
	"demowebgo/controllers"
	"demowebgo/middlewares"
	"github.com/gin-gonic/gin"
)

func UserRoutes(router *gin.RouterGroup) {
    router.GET("/:id", middlewares.AuthorizeJWT([]string{"admin", "user"}), controllers.GetUserByID)
    router.PUT("/update/:id", middlewares.AuthorizeJWT([]string{"admin", "user"}), controllers.UpdateUser)
    router.DELETE("/delete/:id", middlewares.AuthorizeJWT([]string{"admin"}), controllers.DeleteUser)
}
