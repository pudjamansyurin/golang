package main

import (
	"log"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/pudjamansyurin/gomux-gorm/models"
	"github.com/pudjamansyurin/gomux-gorm/routes"
)

type App struct {
	Router *mux.Router
}

func (a *App) Initialize(user, pass, dbname string) {
	models.ConnectDatabase(dbname)
	a.Router = mux.NewRouter()
	a.initializeRoutes()
}

func (a *App) Run(addr string) {
	log.Println("server started on ", addr)
	log.Fatal(http.ListenAndServe(addr, a.Router))
}

func (a *App) initializeRoutes() {
	routes.Product(a.Router)
}
