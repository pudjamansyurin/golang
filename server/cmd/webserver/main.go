package main

import (
	"fmt"
	"log"
	"net/http"

	poker "example.com/server"
)

const dbName = "game.db.json"

func main() {
	store, close, err := poker.FileSystemPlayerStoreFromFile(dbName)
	if err != nil {
		log.Fatal(err)
	}
	defer close()

	alerter := poker.BlindAlerterFunc(poker.Alerter)
	game := poker.NewTexasHoldem(alerter, store)
	server, err := poker.NewPlayerServer(store, game)
	if err != nil {
		log.Fatalf("problem creating player server %v", err)
	}

	fmt.Println("listening on :5000")
	log.Fatal(http.ListenAndServe(":5000", server))
}
