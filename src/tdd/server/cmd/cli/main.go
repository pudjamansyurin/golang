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
	poker.NewCLI(store, os.Stdin).PlayPoker()
}
