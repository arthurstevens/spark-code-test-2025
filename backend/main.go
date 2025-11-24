package main

import (
	"encoding/json"
	"net/http"
)

type Todo struct {
	Title       string `json:"title"`
	Description string `json:"description"`
}

var todos = []Todo{}

func main() {
	http.HandleFunc("/", ToDoListHandler)
	http.ListenAndServe(":8080", nil)
}

func ToDoListHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Access-Control-Allow-Origin", "*")

	switch r.Method {
	case http.MethodGet:
		handleGetList(w, r)
	case http.MethodPost:
		handleAddTodo(w, r)
	default:
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
	}
}

func handleGetList(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	encoder := json.NewEncoder(w)
	encoder.Encode(todos)
}

func handleAddTodo(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	// Decode and validate per openAPI spec
	var t Todo

	decoder := json.NewDecoder(r.Body)
	err := decoder.Decode(&t)
	if err != nil {
		http.Error(w, "Failed to decode JSON, bad request", http.StatusBadRequest)
	}

	if t.Title == "" || t.Description == "" {
		http.Error(w, "Missing fields", http.StatusBadRequest)
	}

	// Store in-memory and return serialised struct
	todos = append(todos, t)

	encoder := json.NewEncoder(w)
	encoder.Encode(t)
}
