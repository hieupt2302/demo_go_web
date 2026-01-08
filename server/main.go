package main
import(
	"github.com/gin-gonic/gin"
	"demowebgo/routes"
	"demowebgo/config"
)
func main() {
	router := gin.Default()
	router.Use()
	config.ConnectDatabase()
	routes.SetupRoutes(router)
	router.Run(":8080")
}