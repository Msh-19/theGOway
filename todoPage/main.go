package main

import (
	"fmt"
	"net/http"
	"todoPage/handlers"
)

func main() {
	http.HandleFunc("/tasks",handleTasks)
	http.HandleFunc("/tasks/",handlers.DeleteTask)


	fs := http.FileServer(http.Dir("./static"))
	http.Handle("/static",http.StripPrefix("/static",fs))

	fmt.Println("Starting server on: 8080")
	if err := http.ListenAndServe(":8080",nil); err != nil{
		fmt.Printf("Error starting server: %s\n",err)
	}
}

func handleTasks(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case http.MethodGet:
		handlers.ListTasks(w, r)
	case http.MethodPost:
		handlers.AddTask(w, r)
	default:
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
	}
}