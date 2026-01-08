package handle_routes

import (
    "demowebgo/controllers"
    "github.com/gin-gonic/gin"
)

func CartRoutes(router *gin.RouterGroup) {
    router.POST("/add", controllers.AddToCart)
    router.GET("/", controllers.GetCart)
}