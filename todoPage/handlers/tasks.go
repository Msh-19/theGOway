package handlers

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"
	"sync"
)

type Task struct{
	ID int `json:"id"`
	Name string `json:"name"`
}

var (
	tasks = []Task{}
	idCount = 1
	mu 		sync.Mutex
)

func ListTasks(w http.ResponseWriter, r *http.Request)  {
	mu.Lock()
	defer mu.Unlock()
	w.Header().Set("Content-Type","application/json")
	json.NewEncoder(w).Encode(tasks)
}

func AddTask(w http.ResponseWriter, r *http.Request)  {
	mu.Lock()
	defer mu.Unlock()

	var task Task
	if err := json.NewDecoder(r.Body).Decode(&task); err!= nil{
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	task.ID = idCount
	idCount++
	tasks = append(tasks, task)

	w.Header().Set("Content-Type","application/json")
	json.NewEncoder(w).Encode(task)
}

func DeleteTask(w http.ResponseWriter, r *http.Request){
	mu.Lock()
	defer mu.Unlock()

	id, err := strconv.Atoi(r.URL.Path[len("/tasks/"):])
	if err != nil{
		http.Error(w,"Invalid task ID",http.StatusBadRequest)
		return
	}

	for i, task := range tasks{
		if task.ID == id{
			tasks = append(tasks[:i], tasks[i+1:]...)
			fmt.Fprintf(w,"Task %d deleted",id)
			return
		}
	}

	http.Error(w,"Task not found",http.StatusNotFound)
}