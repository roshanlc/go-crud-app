package routes

import (
	"github.com/gorilla/mux"
	"github.com/roshanlc/go-crud-app/pkg/controllers"
)

var RegisterBookStoreRoutes = func(router *mux.Router) {
	// Add a book
	router.HandleFunc("/book", controllers.CreateBook).Methods("POST")

	// Get all books
	router.HandleFunc("/book", controllers.GetBook).Methods("GET")

	// Get a book by id
	router.HandleFunc("/book/{id}", controllers.GetBookByID).Methods("GET")

	// delete a book by ids
	router.HandleFunc("/book/{id}", controllers.DeleteBook).Methods("DELETE")

	// TODO: update book details
	//router.HandleFunc("/book/{id}", controllers.UpdateBook).Methods("PUT")

}
