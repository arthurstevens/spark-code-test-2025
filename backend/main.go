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

	// Your code here
}
