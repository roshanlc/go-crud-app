package main

import (
	"log"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/roshanlc/go-crud-app/pkg/routes"
)

func main() {

	// Gorilla mux router
	mux := mux.NewRouter()

	// Register gorilla mux as router for the route handling
	routes.RegisterBookStoreRoutes(mux)

	// Handle homepage by mux router
	http.Handle("/", mux)

	log.Println("Starting web server at :9000")

	// Starting server, and it logs if anything unexpected occurs
	log.Fatal(http.ListenAndServe(":9000", mux))
}
