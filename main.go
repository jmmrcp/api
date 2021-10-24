package main

import (
	"github.com/jmmrcp/api/store"
	"net/http"
	"github.com/gorilla/handlers"
	"log"
	"os"
)

func main() {
	port := os.Getenv("PORT")
	if port == "" {
		log.Fatal("$PORT must be set.")
	}
	router := store.NewRouter()

	allowedOrigins := handlers.AllowedOrigins([]string{*})
	allowedMethods := handlers.AllowedMethods([]string{
		"GET",
		"POST",
		"PUT",
		"DELETE",
	})

	log.Fatal(http.ListenAndServe(":"+port, handlers.CORS(allowedOrigins, allowedMethods)(router)))
}
