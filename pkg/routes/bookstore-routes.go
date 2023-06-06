package routes

import (
	"github.com/DragonCoderz/go-bookstore/pkg/controllers" // "./controllers" would not work, golang does absolute routing and absolute routing only 
	"github.com/gorilla/mux"
)

var RegisterBookStoreRoutes = func(router *mux.Router) {
	//Creation and registration of handlers
	router.HandleFunc("/book/", controllers.CreateBook).Methods("POST") //Essentially delegates any requests to be handled by CreateBook function in controllers file
	router.HandleFunc("/book/", controllers.GetBook).Methods("GET") // ""
	router.HandleFunc("/book/{bookId}", controllers.GetBookById).Methods("GET")
	router.HandleFunc("/book/{bookId}", controllers.UpdateBook).Methods("PUT")
	router.HandleFunc("/book/{bookId}", controllers.DeleteBook).Methods("DELETE")
}