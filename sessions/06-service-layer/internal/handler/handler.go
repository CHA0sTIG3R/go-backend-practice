package handler

import (
	"encoding/json"
	"net/http"

	"github.com/CHA0sTIG3R/go-backend-practice/sessions/06-service-layer/internal/service"
)

func NewTaskHandler(service *service.TaskService) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		if r.Method != http.MethodGet {
			http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
			return
		}

		tasks, err := service.GetTasks()
		if err != nil {
			http.Error(w, "Failed to load tasks", http.StatusInternalServerError)
			return
		}

		w.Header().Set("Content-Type", "application/json")
		json.NewEncoder(w).Encode(tasks)

	}
}
