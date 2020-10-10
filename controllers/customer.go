package controllers

import (
	"log"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
	"github.com/wyllisMonteiro/DB_MONTEIRO_PO1/services"
)

// GetCustomer : Return JSON of a customer with nb product ordered and total price
// valid id : 103
func GetCustomer(w http.ResponseWriter, req *http.Request) {
	urlParams := mux.Vars(req)
	idStr := urlParams["id"]
	id, err := strconv.Atoi(idStr)
	if err != nil {
		log.Println(err)
		return
	}

	customer, err := services.GetCustomer(id)
	if err != nil {
		log.Println(err)
		return
	}

	services.WriteJSON(w, http.StatusOK, customer)
}
