package main

import (
	"fmt"
	"log"
	"os"
	poker "tdd/server"
)

const dbName = "game.db.json"

func main() {
	store, close, err := poker.FileSystemPlayerStoreFromFile(dbName)
	if err != nil {
		log.Fatal(err)
	}
	defer close()

	fmt.Println(`Lets play poker`)
	fmt.Println(`To record, please type "{Name} wins"`)

	alerter := poker.BlindAlerterFunc(poker.StdOutAlerter)
	game := poker.NewTexasHoldem(alerter, store)
	poker.NewCLI(os.Stdin, os.Stdout, game).PlayPoker()
}
