package router

import (
	"github.com/gorilla/mux"

	"github.com/wyllisMonteiro/DB_MONTEIRO_PO1/controllers"
)

// Load controller (handler)
func InitRoutes(r *mux.Router) *mux.Router {
	r.HandleFunc("/customer", controllers.GetCustomer).Methods("GET")
	r.HandleFunc("/order", controllers.GetOrder).Methods("GET")
	r.HandleFunc("/employees", controllers.GetEmployees).Methods("GET")
	r.HandleFunc("/offices", controllers.GetOffices).Methods("GET")
	return r
}
