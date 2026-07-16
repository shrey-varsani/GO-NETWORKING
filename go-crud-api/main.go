package gocrudapi

import (
	"encoding/json"
	"net/http"
)

type Task struct {
	ID    int
	Title string
}

var tasks []Task

func TaskHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	switch r.Method {
	case http.MethodGet: 
		json.NewEncoder(w).Encode(tasks)

	case http.MethodPost:
		var task Task

		err := json.NewDecoder(r.Body).Decode(&task)

		if err != nil {
			http.Error(w, "Invalid JSON", http.StatusBadRequest)
			return
		}

		tasks = append(tasks, task)
		json.NewEncoder(w).Encode(task)

	case http.MethodPut:

		var updatedTask Task

		err := json.NewDecoder(r.Body).Decode(&updatedTask)
		if err != nil {
			http.Error(w, "Invalid JSON", http.StatusBadRequest)
			return
		}

		for i, task := range tasks {
			if task.ID == updatedTask.ID {
				tasks[i].Title = updatedTask.Title
				json.NewEncoder(w).Encode(tasks[i])
				return
			}
		}

		// didn't find in the iteration
		http.Error(w, "Task not found", http.StatusNotFound)

	case http.MethodDelete:
		var taskToDelete Task

		err := json.NewDecoder(r.Body).Decode(&taskToDelete)
		if err != nil {
			http.Error(w, "Invalid JSON", http.StatusBadRequest)
			return
		}

		for i, task := range tasks {
			if task.ID == taskToDelete.ID {
				tasks = append(tasks[:i], tasks[i+1:]...)
				w.Write([]byte("Task Deleted"))
				return 
			}
		}

		http.Error(w, "Task not found", http.StatusNotFound)

	default:
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
	}
}

func main() {
	http.HandleFunc("/tasks", TaskHandler)

	http.ListenAndServe(":8080", nil)
}