// main.go
package main

import (
	"encoding/json"
	"fmt"
	"net/http"
	"time"

	"math/rand"
	"strconv"

	"github.com/dgrijalva/jwt-go"
	"github.com/gorilla/mux"
	"golang.org/x/crypto/bcrypt"
)

var users []User
var jwtKey = []byte("my_secret_key")

type User struct {
	ID       string `json:"id"`
	Username string `json:"username"`
	Password string `json:"password"`
}

type Claims struct {
	Username string `json:"username"`
	jwt.StandardClaims
}

func signUp(w http.ResponseWriter, r *http.Request) {
	var user User
	_ = json.NewDecoder(r.Body).Decode(&user)

	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(user.Password), bcrypt.DefaultCost)
	if err != nil {
		http.Error(w, "Internal Server Error", http.StatusInternalServerError)
		return
	}
	user.ID = strconv.Itoa(rand.Intn(1000000))
	user.Password = string(hashedPassword)
	users = append(users, user)

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(user)
}

func logIn(w http.ResponseWriter, r *http.Request) {
	var credentials User
	_ = json.NewDecoder(r.Body).Decode(&credentials)

	var storedUser User
	for _, user := range users {
		if user.Username == credentials.Username {
			storedUser = user
		}
	}

	if err := bcrypt.CompareHashAndPassword([]byte(storedUser.Password), []byte(credentials.Password)); err != nil {
		http.Error(w, "Invalid Credentials", http.StatusUnauthorized)
		return
	}

	expirationTime := time.Now().Add(5 * time.Minute)
	claims := &Claims{
		Username: credentials.Username,
		StandardClaims: jwt.StandardClaims{
			ExpiresAt: expirationTime.Unix(),
		},
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	tokenString, err := token.SignedString(jwtKey)
	if err != nil {
		http.Error(w, "Internal Server Error", http.StatusInternalServerError)
		return
	}

	http.SetCookie(w, &http.Cookie{
		Name:    "token",
		Value:   tokenString,
		Expires: expirationTime,
	})
}

func main() {
	r := mux.NewRouter()
	r.HandleFunc("/signup", signUp).Methods("POST")
	r.HandleFunc("/login", logIn).Methods("POST")
	r.PathPrefix("/").Handler(http.FileServer(http.Dir("./static/")))


	fmt.Println("Server is running on port 8080")
	http.ListenAndServe(":8080", r)
}