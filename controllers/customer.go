package controllers

import (
	"net/http"

	"github.com/wyllisMonteiro/DB_MONTEIRO_PO1/services"
)

func GetCustomer(w http.ResponseWriter, req *http.Request) {
	customers, err := services.GetCustomer(103)
	if err != nil {
		return
	}

	services.WriteJSON(w, http.StatusOK, customers)
}