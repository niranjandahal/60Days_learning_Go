
package database

import (
	"fmt"
	"log"
	"os"

	"chat-app/models"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)
	
	var DB *gorm.DB
	
	func InitDB() {
		dsn := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=utf8mb4&parseTime=True&loc=Local",
			os.Getenv("DB_USER"),
			os.Getenv("DB_PASSWORD"),
			os.Getenv("DB_HOST"),
			os.Getenv("DB_PORT"),
			os.Getenv("DB_NAME"))
	
		var err error
		DB, err = gorm.Open(mysql.Open(dsn), &gorm.Config{})
		if err != nil {
			log.Fatalf("Failed to connect to database: %v", err)
		}
	
		DB.AutoMigrate(&models.User{}, &models.Message{})
	}
	