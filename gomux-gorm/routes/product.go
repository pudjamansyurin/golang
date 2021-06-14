package routes

import (
	"github.com/gorilla/mux"
	"github.com/pudjamansyurin/gomux-gorm/controllers"
)

func Product(router *mux.Router) {
	r := router.PathPrefix("/products").Subrouter()

	r.HandleFunc("/", controllers.GetProducts).Methods("GET")
	r.HandleFunc("/", controllers.CreateProduct).Methods("POST")
	r.HandleFunc("/{id:[0-9]+}", controllers.GetProduct).Methods("GET")
	r.HandleFunc("/{id:[0-9]+}", controllers.UpdateProduct).Methods("PUT")
	r.HandleFunc("/{id:[0-9]+}", controllers.DeleteProduct).Methods("DELETE")
}
