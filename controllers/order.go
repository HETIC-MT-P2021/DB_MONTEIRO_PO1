package controllers

import (
	"log"
	"net/http"

	"github.com/wyllisMonteiro/DB_MONTEIRO_PO1/services"
)

func GetOrder(w http.ResponseWriter, req *http.Request) {
	order, err := services.GetOrder(10123)
	if err != nil {
		log.Println(err)
		return
	}

	services.WriteJSON(w, http.StatusOK, order)
}
