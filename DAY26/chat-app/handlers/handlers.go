package handlers

import (
	"encoding/json"
	"log"
	"net/http"

	"chat-app/database"
	"chat-app/models"

	"golang.org/x/crypto/bcrypt"
	"gorm.io/gorm"
)

// Credentials is used for user signup and login
type Credentials struct {
	Username string `json:"username"`
	Password string `json:"password"`
}

// SignUp handles user registration
func SignUp(w http.ResponseWriter, r *http.Request) {
	var creds Credentials
	err := json.NewDecoder(r.Body).Decode(&creds)
	if err != nil {
		http.Error(w, "Invalid request payload", http.StatusBadRequest)
		log.Printf("Failed to decode JSON payload: %v", err)
		return
	}

	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(creds.Password), bcrypt.DefaultCost)
	if err != nil {
		http.Error(w, "Internal Server Error", http.StatusInternalServerError)
		log.Printf("Failed to hash password: %v", err)
		return
	}

	user := models.User{
		Username: creds.Username,
		Password: string(hashedPassword),
	}

	result := database.DB.Create(&user)
	if result.Error != nil {
		log.Printf("Failed to create user: %v", result.Error)
		http.Error(w, "Failed to create user", http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(user)
	log.Printf("User created: %s", creds.Username)
}

// LogIn handles user login (without JWT) and issues a simple response
func LogIn(w http.ResponseWriter, r *http.Request) {
	var creds Credentials
	err := json.NewDecoder(r.Body).Decode(&creds)
	if err != nil {
		http.Error(w, "Invalid request payload", http.StatusBadRequest)
		log.Printf("Failed to decode JSON payload: %v", err)
		return
	}

	var user models.User
	result := database.DB.Where("username = ?", creds.Username).First(&user)
	if result.Error == gorm.ErrRecordNotFound || bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(creds.Password)) != nil {
		http.Error(w, "Invalid username or password", http.StatusUnauthorized)
		log.Printf("Invalid username or password for user: %s", creds.Username)
		return
	}

	w.WriteHeader(http.StatusOK)
	w.Write([]byte("Login successful"))
	log.Printf("User logged in: %s", creds.Username)
}

// AddMessage handles adding a new chat message
func AddMessage(w http.ResponseWriter, r *http.Request) {
	type MessageRequest struct {
		UserID  uint   `json:"user_id"`
		Content string `json:"content"`
	}

	var msgReq MessageRequest
	err := json.NewDecoder(r.Body).Decode(&msgReq)
	if err != nil {
		http.Error(w, "Invalid request payload", http.StatusBadRequest)
		log.Printf("Failed to decode JSON payload: %v", err)
		return
	}

	msg := models.Message{
		UserID:  msgReq.UserID,
		Content: msgReq.Content,
	}

	result := database.DB.Create(&msg)
	if result.Error != nil {
		log.Printf("Failed to create message: %v", result.Error)
		http.Error(w, "Failed to create message", http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(msg)
	log.Printf("Message created: %v", msg)
}

// GetMessages handles retrieving all chat messages
func GetMessages(w http.ResponseWriter, r *http.Request) {
	var messages []models.Message
	result := database.DB.Preload("User").Order("created_at asc").Find(&messages)
	if result.Error != nil {
		log.Printf("Failed to retrieve messages: %v", result.Error)
		http.Error(w, "Failed to retrieve messages", http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(messages)
	log.Println("Messages retrieved")
}
