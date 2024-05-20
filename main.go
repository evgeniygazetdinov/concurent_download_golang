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
	// Create a new router
	r := mux.NewRouter()

	// Define your HTTP routes using the router
	r.HandleFunc("/api/user", createUserHandler).Methods("POST")
	r.HandleFunc("/api/user/{id}", getUserHandler).Methods("GET")
	r.HandleFunc("/api/user/{id}", updateUserHandler).Methods("PUT")
	r.HandleFunc("/api/user/{id}", deleteUserHandler).Methods("DELETE")
	r.PathPrefix("/swagger/").Handler(httpSwagger.WrapHandler)

	// Start the HTTP server on port 8090
	log.Println("Server listening on http://0.0.0.0:8082")
	log.Fatal(http.ListenAndServe(":8082", r))
}
