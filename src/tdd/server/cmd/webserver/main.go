package main

import (
	"log"
	"net/http"
	"os"

	poker "tdd/server"
)

const dbName = "game.db.json"

func main() {
	db, err := os.OpenFile(dbName, os.O_RDWR|os.O_CREATE, 0666)
	if err != nil {
		log.Fatalf(`problem opening %s, %v`, dbName, err)
	}

	store, err := poker.NewFileSystemPlayerStore(db)
	if err != nil {
		log.Fatalf(`problem creating file system store, %v`, err)
	}

	server := poker.NewPlayerServer(store)

	if err := http.ListenAndServe(":5000", server); err != nil {
		log.Fatalf(`could not listen on :5000, %v`, err)
	}
}
