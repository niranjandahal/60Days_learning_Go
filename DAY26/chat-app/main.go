package main

import (
	"log"
	"net/http"

	"chat-app/database"
	"chat-app/handlers"

	"github.com/gorilla/mux"
	"github.com/joho/godotenv"
)

func main() {
	err := godotenv.Load()
	if err != nil {
		log.Fatalf("Error loading .env file")
	}

	database.InitDB()

	r := mux.NewRouter()

	// Route to handle user registration
	r.HandleFunc("/signup", handlers.SignUp).Methods("POST")

	// Route to handle user login
	r.HandleFunc("/login", handlers.LogIn).Methods("POST")

	// Route to handle adding a new chat message
	r.HandleFunc("/messages", handlers.AddMessage).Methods("POST")

	// Route to handle retrieving all chat messages
	r.HandleFunc("/messages", handlers.GetMessages).Methods("GET")

	// Serve static files
	r.PathPrefix("/").Handler(http.FileServer(http.Dir("./static/")))

	log.Println("Server is running on port 8080")
	http.ListenAndServe(":8080", r)
}
