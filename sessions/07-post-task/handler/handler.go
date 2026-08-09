package handler

import (
	"encoding/json"
	"net/http"

	"github.com/CHA0sTIG3R/go-backend-practice/sessions/07-post-task/jsonutil"
	"github.com/CHA0sTIG3R/go-backend-practice/sessions/04-json-tasks/task"
	"github.com/CHA0sTIG3R/go-backend-practice/sessions/07-post-task/service"
	"github.com/CHA0sTIG3R/go-backend-practice/sessions/07-post-task/validation"
)

func NewTaskHandler(service *service.TaskService) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		switch r.Method {
		case http.MethodGet:
			getTaskHandler(w, r, service)
		case http.MethodPost:
			postTaskHandler(w, r, service)
		default:
			http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		}
	}
}

func getTaskHandler(w http.ResponseWriter, r *http.Request, service *service.TaskService) {

	tasks, err := service.GetTasks()
	if err != nil {
		http.Error(w, "Failed to load tasks", http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(tasks)

}

func postTaskHandler(w http.ResponseWriter, r *http.Request, service *service.TaskService) {

	var newTask task.Task
	err := json.NewDecoder(r.Body).Decode(&newTask)
	if err != nil {
		http.Error(w, "Invalid request body", http.StatusBadRequest)
		return
	}

	err = validation.ValidateTask(newTask)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	err = service.AddTask(newTask)
	if err != nil {
		if _, ok := err.(*jsonutil.DuplicateTaskError); ok {
			http.Error(w, err.Error(), http.StatusConflict)
			return
		}
		http.Error(w, "Failed to add task", http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(newTask)
}