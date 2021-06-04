package main

import (
	"github.com/gorilla/mux"
	"github.com/jinzhu/gorm"
)

type App struct {
	Router *mux.Router
	DB     *gorm.DB
}

func (a *App) Initialize(user, pass, dbname string) {
	a.DB := models.ConnectDatabase(dbname) 
	a.Router := mux.NewRouter()
}

func (a *App) Run(addr string) {

}
