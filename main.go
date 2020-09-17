package main

import (
	"log"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/wyllisMonteiro/DB_MONTEIRO_P01/router"
)

func main() {
	r := mux.NewRouter()
	router.InitRoutes(r)
	log.Fatal(http.ListenAndServe(":9000", r))
}
