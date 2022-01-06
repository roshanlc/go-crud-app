package controllers

import (
	"encoding/json"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
	"github.com/roshanlc/go-crud-app/pkg/models"
	"github.com/roshanlc/go-crud-app/pkg/utils"
)

// CreateBook handles handles POST request to create new book resource
func CreateBook(w http.ResponseWriter, r *http.Request) {

	newBook := &models.Book{}
	utils.ParseBody(r, newBook)

	b := newBook.CreateBook()

	w.WriteHeader(http.StatusOK)
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(b)

}

// GetBook handles GET request to list existing book resources
func GetBook(w http.ResponseWriter, r *http.Request) {

	newBooks := models.GetAllBooks()

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(newBooks)

}

// GetBookByID handles GET request to details of a specific book by its id, if it exists

func GetBookByID(w http.ResponseWriter, r *http.Request) {

	vars := mux.Vars(r)
	id := vars["id"]
	ID, err := strconv.Atoi(id)

	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte(err.Error()))
		return
	}
	book, _ := models.GetBookByID(ID)

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(book)

}

// DeleteBook handles DELETE request which deletes a book resource by its id
func DeleteBook(w http.ResponseWriter, r *http.Request) {

	vars := mux.Vars(r)
	id := vars["id"]
	ID, err := strconv.Atoi(id)

	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte(err.Error()))
		return
	}
	x := models.DeleteBook(ID)

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(x)

}
