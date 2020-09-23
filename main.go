package main

import (
	"log"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/wyllisMonteiro/DB_MONTEIRO_PO1/router"
	"github.com/wyllisMonteiro/DB_MONTEIRO_PO1/models"
)

func main() {
	r := mux.NewRouter()
	router.InitRoutes(r)
	
	models.ConnectToBDD()
	
	log.Fatal(http.ListenAndServe(":9000", r))
}
