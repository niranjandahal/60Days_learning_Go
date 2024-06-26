package models

import (
	"gorm.io/gorm"
)

// User represents a user in the chat application
type User struct {
	gorm.Model
	Username string `gorm:"unique"`
	Password string
}

// Message represents a chat message
type Message struct {
	gorm.Model
	UserID  uint
	Content string
}
