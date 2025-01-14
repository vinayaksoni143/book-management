package routes

import (
	"github.com/gorilla/mux"
	"github.com/vinayaksoni143/book-management/pkg/controllers"
)

var RegisterBookStoreRoutes = func(router *mux.Router) {
	// create book
	router.HandleFunc("/book/", controllers.CreateBook).Methods("POST")

	// get books
	router.HandleFunc("/book/", controllers.GetBook).Methods("GET")

	// get book by id
	router.HandleFunc("/book/{bookId}", controllers.GetBookById).Methods("GET")

	// update book
	router.HandleFunc("/book/{bookId}", controllers.UpdateBook).Methods("PUT")

	// delete book
	router.HandleFunc("/book/{bookId}", controllers.DeleteBook).Methods("DELETE")
}
