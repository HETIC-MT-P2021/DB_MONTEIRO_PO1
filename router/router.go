package router

import (
	"github.com/gorilla/mux"

	"github.com/wyllisMonteiro/DB_MONTEIRO_P01/controllers"
)

// Load controller (handler)
func InitRoutes(r *mux.Router) *mux.Router {
	r.HandleFunc("/employee", controllers.GetEmployee).Methods("GET")
	return r
}
