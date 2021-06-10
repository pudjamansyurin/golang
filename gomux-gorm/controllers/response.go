package controllers

import (
	"encoding/json"
	"log"
	"net/http"
)

func respError(w http.ResponseWriter, code int, msg string) {
	respJSON(w, code, map[string]string{"error": msg})
}

func respJSON(w http.ResponseWriter, code int, payload interface{}) {
	res, err := json.Marshal(payload)
	if err != nil {
		log.Println("could not parse json:", err)
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(code)
	w.Write(res)
}
