package handler

import (
	"encoding/json"
	"net/http"

	"github.com/CHA0sTIG3R/go-backend-practice/sessions/04-json-tasks/task"
)

func TasksHandler(tasks []task.Task) http.HandlerFunc {

	handler := func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")

		json.NewEncoder(w).Encode(tasks)
	}

	return http.HandlerFunc(handler)
}
