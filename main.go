package main

import (
	"crud-golang/routers"
	_ "github.com/go-sql-driver/mysql"
	"github.com/gorilla/mux"
	"log"
	"net/http"
)

func main() {
	r := mux.NewRouter()
	routers.Routers(r)
	http.Handle("/", r)
	log.Fatal(http.ListenAndServe("localhost:8900", r))
}
