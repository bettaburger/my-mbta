package routes

import (
	"github.com/gorilla/mux"
	"github/bettaburger/my-mbta/controllers"
)

var AlertRoutes = func(router *mux.Router) {
	router.HandleFunc("/alerts", controllers.AlertHandler).Methods("GET","OPTIONS")
}
