package handlers

import (
	"encoding/json"
	"log"
	"net/http"

	"chat-app/database"
	"chat-app/models"

	"github.com/gorilla/sessions"
	"golang.org/x/crypto/bcrypt"
	"gorm.io/gorm"
)

var store = sessions.NewCookieStore([]byte("your-secret-key"))

type Credentials struct {
	Username string `json:"username"`
	Password string `json:"password"`
}

// SignUp handles user registration and session creation
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

	// Create a session and set the user ID in the session
	session, _ := store.Get(r, "session-name")
	session.Values["user_id"] = user.ID
	session.Save(r, w)

	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(user)
	log.Printf("User created and session started: %s", creds.Username)
}

func LogIn(w http.ResponseWriter, r *http.Request) {
	var creds Credentials
	err := json.NewDecoder(r.Body).Decode(&creds)
	if err != nil {
		http.Error(w, "Invalid request payload", http.StatusBadRequest)
		log.Printf("Failed to decode JSON payload: %v", err)
		return
	}

	session, _ := store.Get(r, "session-name")
	userID := session.Values["user_id"]
	if userID != nil {
		log.Printf("session found")
		http.Redirect(w, r, "/static/homepage.html", http.StatusSeeOther)
		return
	}

	log.Printf("session not found")

	var user models.User
	result := database.DB.Where("username = ?", creds.Username).First(&user)
	if result.Error == gorm.ErrRecordNotFound || bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(creds.Password)) != nil {
		http.Error(w, "Invalid username or password", http.StatusUnauthorized)
		log.Printf("Invalid username or password for user: %s", creds.Username)
		return
	}

	session.Values["user_id"] = user.ID
	session.Save(r, w)

	w.WriteHeader(http.StatusOK)
	w.Write([]byte("Login successful"))
	log.Printf("User logged in: %s", creds.Username)
}
