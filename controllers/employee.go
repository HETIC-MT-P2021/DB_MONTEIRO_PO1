package controllers

import (
	"net/http"

	"github.com/wyllisMonteiro/DB_MONTEIRO_PO1/services"
)

func GetEmployees(w http.ResponseWriter, req *http.Request) {
	customer, err := services.GetEmployees()
	if err != nil {
		return
	}

	services.WriteJSON(w, http.StatusOK, customer)
}
