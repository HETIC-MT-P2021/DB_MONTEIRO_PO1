package controllers

import (
	"log"
	"net/http"

	"github.com/wyllisMonteiro/DB_MONTEIRO_PO1/services"
)

func GetCustomer(w http.ResponseWriter, req *http.Request) {
	customers, err := services.GetCustomer(103)
	if err != nil {
		log.Println(err)
		return
	}

	services.WriteJSON(w, http.StatusOK, customers)
}
