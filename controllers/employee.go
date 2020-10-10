package controllers

import (
	"log"
	"net/http"

	"github.com/wyllisMonteiro/DB_MONTEIRO_PO1/services"
)

// GetEmployees : Return JSON of all employees with associated office
func GetEmployees(w http.ResponseWriter, req *http.Request) {
	customer, err := services.GetEmployees()
	if err != nil {
		log.Println(err)
		return
	}

	services.WriteJSON(w, http.StatusOK, customer)
}
