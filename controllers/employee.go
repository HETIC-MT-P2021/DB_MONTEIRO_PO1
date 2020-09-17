package controllers

import (
	"net/http"

	"github.com/wyllisMonteiro/mailing/api/service"
)

type Employee struct {
	Name string `json:"name"`
	Age  int    `json:"age"`
}

func GetEmployee(w http.ResponseWriter, req *http.Request) {
	employee := Employee{
		Name: "Wyllis",
		Age:  22,
	}
	service.WriteJSON(w, http.StatusOK, employee)
}
