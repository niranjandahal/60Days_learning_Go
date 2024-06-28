module chat-app

go 1.22.3

require (
	github.com/joho/godotenv v1.5.1
	golang.org/x/crypto v0.0.0-20191205180655-e7c4368fe9dd
)

require github.com/gorilla/securecookie v1.1.2 // indirect

require (
	github.com/go-sql-driver/mysql v1.7.0 // indirect
	github.com/jinzhu/now v1.1.5 // indirect
	gorm.io/gorm v1.25.10
)

require (
	github.com/dgrijalva/jwt-go v3.2.0+incompatible
	github.com/gorilla/mux v1.8.1
	github.com/gorilla/sessions v1.3.0
	github.com/jinzhu/inflection v1.0.0 // indirect
	gorm.io/driver/mysql v1.5.7
)
