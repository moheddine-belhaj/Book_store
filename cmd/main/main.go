package main

import (
	"log"
	"net/http"
	"fmt"
	"github.com/gorilla/mux"
	_ "github.com/jinzhu/gorm/dialects/postgres" // Use PostgreSQL dialect
	"github.com/moheddine-belhaj/Book_store/pkg/routes"
)

func main() {
	fmt.Println("Starting...")
	r := mux.NewRouter()
	routes.RegisterBookStoreRoutes(r)
	http.Handle("/", r)
	log.Fatal(http.ListenAndServe("localhost:9010", r))
}
