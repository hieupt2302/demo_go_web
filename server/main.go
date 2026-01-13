package main
import(
	"github.com/gin-gonic/gin"
	"demowebgo/routes"
	"demowebgo/config"
	"github.com/gin-contrib/cors"
	"time"
)
func main() {
	router := gin.Default()
	router.Use(cors.New(cors.Config{
		AllowOrigins:     []string{"http://localhost:5173"},
		AllowMethods:     []string{"GET", "POST", "PUT", "DELETE"},
		AllowHeaders:     []string{"Origin", "Content-Type", "Authorization"},
		ExposeHeaders:    []string{"Content-Length"},
		AllowCredentials: true,
		MaxAge: 12 * time.Hour,
	}))
	config.ConnectDatabase()
	router.Static("/uploads", "./uploads")
	routes.SetupRoutes(router)
	router.Run(":8080")
}