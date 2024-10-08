package routes

import (
	"github.com/gorilla/mux"
	"github.com/moheddine-belhaj/Book_store/pkg/controllers"
)

var RegisterBookStoreRoutes = func(router *mux.Router) {
	router.HandleFunc("/book/create", controllers.CreateBook).Methods("POST")
	router.HandleFunc("/book/all", controllers.GetBook).Methods("GET")
	router.HandleFunc("/book/{bookId}", controllers.GetBookById).Methods("GET")
	router.HandleFunc("/book/update/{bookId}", controllers.UpdateBook).Methods("PUT")
	router.HandleFunc("/book/delete/{bookId}", controllers.DeleteBook).Methods("DELETE")
}