package main

import (
	"log"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/wyllisMonteiro/DB_MONTEIRO_PO1/models"
	"github.com/wyllisMonteiro/DB_MONTEIRO_PO1/router"
)

func main() {
	r := mux.NewRouter()
	router.InitRoutes(r)

	models.ConnectToBDD()

	log.Fatal(http.ListenAndServe(":9000", r))
}
