package main

import (
	"fmt"
	"log"
	"net/http"

	"example.com/todo/router"
)

func main() {
	r := router.Router()

	fmt.Println("Staring server on :8080")
	log.Fatal(http.ListenAndServe(":8080", r))
}
