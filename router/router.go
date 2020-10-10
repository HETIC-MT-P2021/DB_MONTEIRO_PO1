package router

import (
	"github.com/gorilla/mux"

	"github.com/wyllisMonteiro/DB_MONTEIRO_PO1/controllers"
)

// Load controller (handler)
func InitRoutes(r *mux.Router) *mux.Router {
	r.HandleFunc("/customers", controllers.GetCustomer).Methods("GET")
	return r
}
