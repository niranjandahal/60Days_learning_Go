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
	r.HandleFunc("/signup", handlers.SignUp).Methods("POST")
	r.HandleFunc("/login", handlers.LogIn).Methods("POST")
	r.HandleFunc("/ws", handlers.HandleWebSocket) // WebSocket endpoint
	r.PathPrefix("/").Handler(http.FileServer(http.Dir("./static/")))

	log.Println("Server is running on port 8080")
	http.ListenAndServe(":8080", r)
}