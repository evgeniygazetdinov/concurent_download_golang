package main

import (
	"log"
	"net/http"

	_ "github.com/go-sql-driver/mysql"
	"github.com/gorilla/mux"
	httpSwagger "github.com/swaggo/http-swagger"
)

const (
	dbDriver = "mysql"
	dbUser   = "root"
	dbPass   = "1234"
	dbName   = "gocrud"
)

type User struct {
	ID    int
	Name  string
	Email string
}

type Chat struct {
	ID       int
	UsersIds string
}

type Message struct {
	ID     int
	Text   string
	ChatId int
}

type Like struct {
	ID      int
	Owner   int
	UserId  int
	Message string
}

// @title           Swagger Example API
// @version         1.0
// @description     This is a sample server celler server.
// @termsOfService  http://swagger.io/terms/

// @contact.name   API Support
// @contact.url    http://www.swagger.io/support
// @contact.email  support@swagger.io

// @license.name  Apache 2.0
// @license.url   http://www.apache.org/licenses/LICENSE-2.0.html

// @host      localhost:8080
// @BasePath  /api/v1

// @securityDefinitions.basic  BasicAuth

// @externalDocs.description  OpenAPI
// @externalDocs.url          https://swagger.io/resources/open-api/
func main() {
	r := mux.NewRouter()
	r.HandleFunc("/api/user", createUserHandler).Methods("POST")
	r.HandleFunc("/api/user/{id}", getUserHandler).Methods("GET")
	r.HandleFunc("/api/user/{id}", updateUserHandler).Methods("PUT")
	r.HandleFunc("/api/user/{id}", deleteUserHandler).Methods("DELETE")
	r.HandleFunc("/api/user/setLike/{id}/{id}", deleteUserHandler).Methods("POST")

	r.HandleFunc("/api/chat", createChatHandler).Methods("POST")

	r.PathPrefix("/swagger/").Handler(httpSwagger.WrapHandler)
	log.Println("Server listening on http://0.0.0.0:8082")
	log.Fatal(http.ListenAndServe(":8082", r))
}
