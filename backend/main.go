package main

import "net/http"

type Todo struct {
	Title       string `json:"title"`
	Description string `json:"description"`
}

func main() {
	// Your code here
}

func ToDoListHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Access-Control-Allow-Origin", "*")

	// Your code here
}
