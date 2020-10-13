package router

import (
	"github.com/gorilla/mux"

	"github.com/wyllisMonteiro/DB_MONTEIRO_PO1/controllers"
)

// InitRoutes : Load controller (handler)
func InitRoutes(r *mux.Router) *mux.Router {
	var api = r.PathPrefix("/api/v1").Subrouter()
	api.HandleFunc("/customers/{id}", controllers.GetCustomer).Methods("GET")
	api.HandleFunc("/orders/{id}", controllers.GetOrder).Methods("GET")
	api.HandleFunc("/employees", controllers.GetEmployees).Methods("GET")
	api.HandleFunc("/offices", controllers.GetOffices).Methods("GET")
	return r
}
