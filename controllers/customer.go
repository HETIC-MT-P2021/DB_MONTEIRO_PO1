package controllers

import (
	"fmt"
	"net/http"

	"github.com/wyllisMonteiro/DB_MONTEIRO_PO1/models"
	"github.com/wyllisMonteiro/DB_MONTEIRO_PO1/services"
)

func GetCustomer(w http.ResponseWriter, req *http.Request) {
	customers, err := models.GetCustomer(119)
	if err != nil {
		fmt.Println("erreur db")
		return
	}

	services.WriteJSON(w, http.StatusOK, customers)
}
