package handle_routes

import (
	"demowebgo/controllers"
	"demowebgo/middlewares"
	"github.com/gin-gonic/gin"
)

func UserRoutes(router *gin.RouterGroup) {
	router.POST("/create", controllers.CreateUser)
	router.GET("/:id", middlewares.AuthMiddleware(), middlewares.AuthorizeJWT([]string{"admin","user"}), controllers.GetUserByID)
	router.PUT("/update/:id", controllers.UpdateUser)
	router.DELETE("/delete/:id", controllers.DeleteUser)
}
