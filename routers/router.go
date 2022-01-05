package routers

import (
	"crud-golang/controllers"
	"github.com/gorilla/mux"
)

var Routers = func(router *mux.Router) {
	router.HandleFunc("/book", controllers.CreateBook).Methods("POST")
	router.HandleFunc("/book", controllers.GetBook).Methods("GET")
	router.HandleFunc("/book/{id}", controllers.GetBookById).Methods("GET")
	router.HandleFunc("/update-book/{id}", controllers.UpdateBook).Methods("POST")
	router.HandleFunc("/delete-book/{id}", controllers.DeleteBook).Methods("POST")
}
