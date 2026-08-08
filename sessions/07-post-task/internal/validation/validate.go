package validation

import (
	"net/http"
	"github.com/CHA0sTIG3R/go-backend-practice/sessions/04-json-tasks/task"
)

func ValidateTask(w http.ResponseWriter, task task.Task) error {
	if task.Name == "" {
		http.Error(w, "Name is required", http.StatusBadRequest)
		return nil
	}
	if task.Priority < 1 || task.Priority > 3 {
		http.Error(w, "Priority must be between 1 and 3", http.StatusBadRequest)
		return nil
	}

	if len(task.Name) > 100 {
		http.Error(w, "Name cannot exceed 100 characters", http.StatusBadRequest)
		return nil
	}

	return nil
}