package config

import (
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
	"demowebgo/models"
)

var TestDB *gorm.DB

func SetupTestDB() {
	db, _ := gorm.Open(sqlite.Open(":memory:"), &gorm.Config{})
	db.AutoMigrate(&models.User{})
	TestDB = db
}
