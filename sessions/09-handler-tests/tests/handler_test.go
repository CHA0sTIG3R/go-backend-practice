package tests

import (
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"path/filepath"
	"testing"

	"github.com/CHA0sTIG3R/go-backend-practice/sessions/04-json-tasks/task"
	"github.com/CHA0sTIG3R/go-backend-practice/sessions/07-post-task/handler"
	"github.com/CHA0sTIG3R/go-backend-practice/sessions/07-post-task/service"
)

func TestTaskHandler(t *testing.T) {
	filename := filepath.Join(t.TempDir(), "tasks.json")

	taskService := service.NewTaskService(filename)

	err := taskService.AddTask(task.Task{
		Name:      "Practice handler tests",
		Priority:  1,
		Completed: false,
	})
	if err != nil {
		t.Fatalf("Failed to arrange test data: %v", err)
	}

	taskHandler := handler.NewTaskHandler(taskService)

	req := httptest.NewRequest(http.MethodGet, "/tasks", nil)
	rr := httptest.NewRecorder()

	taskHandler.ServeHTTP(rr, req)

	if rr.Code != http.StatusOK {
		t.Fatalf("status = %d, want %d", rr.Code, http.StatusOK)
	}

	if got := rr.Header().Get("Content-Type"); got != "application/json" {
		t.Errorf("Content-Type = %q, want %q", got, "application/json")
	}

	var tasks []task.Task
	if err := json.Unmarshal(rr.Body.Bytes(), &tasks); err != nil {
		t.Fatalf("Failed to decode response body: %v", err)
	}

	if len(tasks) != 1 {
		t.Fatalf("received %d tasks, want 1", len(tasks))
	}

	if tasks[0].Name != "Practice handler tests" {
		t.Errorf(
			"task name = %q, want %q",
			tasks[0].Name,
			"Practice handler tests",
		)
	}
	if tasks[0].Priority != 1 {
		t.Errorf(
			"task priority = %d, want %d",
			tasks[0].Priority,
			1,
		)
	}
}
