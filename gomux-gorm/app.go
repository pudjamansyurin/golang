package main

import (
	"log"
	"net/http"

	"github.com/gorilla/mux"
	_ "github.com/jinzhu/gorm/dialects/sqlite"
	"github.com/pudjamansyurin/gomux-gorm/models"
	"github.com/pudjamansyurin/gomux-gorm/routes"

	"github.com/gorilla/handlers"
)

type App struct {
	Router *mux.Router
}

func (a *App) Initialize(dbname string) {
	models.Connect(dbname)
	a.initializeRoutes()
}

func (a *App) Run(addr string) {
	log.Println("server started on ", addr)

	err := http.ListenAndServe(addr, handlers.CORS(
		handlers.AllowedHeaders([]string{"X-Requested-With", "X-Total-Count", "Content-Type", "Authorization"}),
		handlers.AllowedMethods([]string{"GET", "POST", "PUT", "HEAD", "OPTIONS"}),
		handlers.AllowedOrigins([]string{"*"}),
	)(a.Router))
	if err != nil {
		log.Fatal("failed to start server", err.Error())
	}
}

func (a *App) initializeRoutes() {
	a.Router = mux.NewRouter().StrictSlash(true)
	routes.Product(a.Router)
}
