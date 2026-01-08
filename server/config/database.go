package config

import (
	"fmt"
	"log"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"github.com/joho/godotenv"
	"os"
	"demowebgo/models"
)

var DB *gorm.DB

func ConnectDatabase() {
	_ = godotenv.Load()
	fmt.Println(os.Getenv("DB_USER"))
	fmt.Println(os.Getenv("DB_PASS"))
	dsn := os.Getenv("DB_USER") + ":" + os.Getenv("DB_PASSWORD") + "@tcp(127.0.0.1:3306)/" + os.Getenv("DB_NAME") + "?charset=utf8mb4&parseTime=True&loc=Local"

	database, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatal("Failed to connect to database: ", err)
	}
	DB = database
	DB.AutoMigrate(&models.User{})
	log.Println("✅ Database connected!")
}