package middleware

import (
	"encoding/json"
	"log"
	"net/http"

	"github.com/gorilla/mux"

	"example.com/todo/models"
)

func GetAllTask(w http.ResponseWriter, r *http.Request) {
	setHeader(w, "")

	payload := getAllTask()
	json.NewEncoder(w).Encode(payload)
}

func CreateTask(w http.ResponseWriter, r *http.Request) {
	setHeader(w, "POST")

	var task models.ToDoList
	err := json.NewDecoder(r.Body).Decode(&task)
	if err != nil {
		log.Println(err)
	}

	insertOneTask(task)
	json.NewEncoder(w).Encode(task)
}

func TaskComplete(w http.ResponseWriter, r *http.Request) {
	setHeader(w, "PUT")
	taskById(w, r, taskComplete)
}

func UndoTask(w http.ResponseWriter, r *http.Request) {
	setHeader(w, "PUT")
	taskById(w, r, undoTask)
}

func DeleteTask(w http.ResponseWriter, r *http.Request) {
	setHeader(w, "DELETE")
	taskById(w, r, deleteOneTask)
}

func DeleteAllTask(w http.ResponseWriter, r *http.Request) {
	setHeader(w, "")

	count := deleteAllTask()
	json.NewEncoder(w).Encode(count)
}

func taskById(w http.ResponseWriter, r *http.Request, fn func(string)) {
	params := mux.Vars(r)
	fn(params["id"])
	json.NewEncoder(w).Encode(params["id"])
}

func setHeader(w http.ResponseWriter, method string) {
	w.Header().Set("Context-Type", "application/x-www-form-urlencoded")
	w.Header().Set("Access-Control-Allow-Origin", "*")
	if method != "" {
		w.Header().Set("Access-Control-Allow-Headers", "Content-Type")
		w.Header().Set("Access-Control-Allow-Methods", method)
	}
}
