package main

import (
	"log"
	"net/http"
	"github.com/gorilla/mux"
	"github/bettaburger/my-mbta/routes"
)

func main() {
	// Create the router
	router := mux.NewRouter()
	// Handle
	routes.AlertRoutes(router)

	// Starting the server on port 3000
	log.Println("Server is starting on port 9000")
	http.ListenAndServe(":9000", router) 
}