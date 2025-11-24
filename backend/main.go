package main

import "net/http"

type Todo struct {
	Title       string `json:"title"`
	Description string `json:"description"`
}

var todos []Todo

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
