package routes

import (
	"github.com/gorilla/mux"
	"github.com/pudjamansyurin/gomux-gorm/controllers"
)

func Product(router *mux.Router) {
	p := router.PathPrefix("/product").Subrouter()

	p.HandleFunc("/", controllers.GetProducts).Methods("GET")
	p.HandleFunc("/", controllers.CreateProduct).Methods("POST")
	p.HandleFunc("/{id:[0-9]+}", controllers.GetProduct).Methods("GET")
	p.HandleFunc("/{id:[0-9]+}", controllers.UpdateProduct).Methods("PUT")
	p.HandleFunc("/{id:[0-9]+}", controllers.DeleteProduct).Methods("DELETE")
}
