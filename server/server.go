package server

import (
	"fmt"
	"log"
	"net/http"
	
	"github.com/Masayuki-Suzuki/instaclone/config"
	"github.com/gorilla/handlers"
)

func Init() {
	port := ":" + config.GetConfig().GetString("server.port")
	router := NewRouter()
	
	// SetUp CORS
	allowedOrigins := handlers.AllowedOrigins([]string{"http://localhost:3000", "http://localhost:8088"})
	allowedMethods := handlers.AllowedMethods([]string{"GET", "POST", "DELETE", "PUT"})
	allowedHeaders := handlers.AllowedHeaders([]string{"Authorization", "Origin", "Content-Type", "Accept"})
	
	// Shows message.
	fmt.Println("Server: Listening..")
	log.Fatal(http.ListenAndServe(port, handlers.CORS(allowedOrigins, allowedMethods, allowedHeaders)(router)))
}
