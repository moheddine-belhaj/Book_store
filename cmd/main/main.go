package main

import (
	"log"
	"net/http"
	"github.com/gorilla/mux"
	"github.com/rs/cors"
	"github.com/moheddine-belhaj/Book_store/pkg/routes"
)

func main() {
	r := mux.NewRouter()

	// Register the routes for the book store
	routes.RegisterBookStoreRoutes(r)

	// Use CORS middleware to allow requests from the Vue.js app
	c := cors.New(cors.Options{
		AllowedOrigins:   []string{"http://localhost:8081"}, // Your Vue.js app origin
		AllowedMethods:   []string{"GET", "POST", "PUT", "DELETE", "OPTIONS"},
		AllowedHeaders:   []string{"Content-Type"},
		AllowCredentials: true,
	})

	// Wrap the router with CORS middleware
	handler := c.Handler(r)

	// Start the server
	log.Fatal(http.ListenAndServe(":9010", handler))
}
