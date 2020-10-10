package controllers

import (
	"log"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
	"github.com/wyllisMonteiro/DB_MONTEIRO_PO1/services"
)

// GetOrder : Return JSON of an order with associated products
// valid id : 10123
func GetOrder(w http.ResponseWriter, req *http.Request) {
	urlParams := mux.Vars(req)
	idStr := urlParams["id"]
	id, err := strconv.Atoi(idStr)
	if err != nil {
		log.Println(err)
		return
	}

	order, err := services.GetOrder(id)
	if err != nil {
		log.Println(err)
		return
	}

	services.WriteJSON(w, http.StatusOK, order)
}
