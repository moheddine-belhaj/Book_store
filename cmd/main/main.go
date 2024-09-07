package main

import (
	"log"
	"net/http"
	// path/filepath
	"github.com/gorilla/mux"
	_ "github.com/jinzhu/gorm/dialects/postgres"
	"github.com/moheddine-belhaj/Book_store/pkg/routes"
)

func main() {
	r := mux.NewRouter()

	// Serve static files from the dist directory
	r.PathPrefix("/").Handler(http.StripPrefix("/", http.FileServer(http.Dir("./dist/"))))

	routes.RegisterBookStoreRoutes(r)
	http.Handle("/", r)
	log.Fatal(http.ListenAndServe("localhost:9010", r))
}
