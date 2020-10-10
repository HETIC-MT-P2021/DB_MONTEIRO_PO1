package controllers

import (
	"log"
	"net/http"

	"github.com/wyllisMonteiro/DB_MONTEIRO_PO1/services"
)

// GetOffices : Return JSON of all offices with associated employees
func GetOffices(w http.ResponseWriter, req *http.Request) {
	order, err := services.GetOffices()
	if err != nil {
		log.Println(err)
		return
	}

	services.WriteJSON(w, http.StatusOK, order)
}
