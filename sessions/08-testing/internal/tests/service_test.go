package test

import (
	"os"
	"testing"

	"github.com/CHA0sTIG3R/go-backend-practice/sessions/04-json-tasks/task"
	"github.com/CHA0sTIG3R/go-backend-practice/sessions/07-post-task/service"
)

func TestService(t *testing.T) {
	service := service.NewTaskService("test_tasks.json")

	// table-driven tests for adding tasks and rejecting duplicates
	tests := []struct {
		name    string
		task    task.Task
		wantErr bool
	}{
		{
			name: "add valid task",
			task: task.Task{
				Name:     "Test Task 1",
				Priority: 1,
			},
			wantErr: false,
		},
		{
			name: "add duplicate task",
			task: task.Task{
				Name:     "Test Task 1",
				Priority: 1,
			},
			wantErr: true,
		},
	}

	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			err := service.AddTask(tc.task)
			if (err != nil) != tc.wantErr {
				t.Errorf("AddTask() error = %v, wantErr %v", err, tc.wantErr)
			}
		})
	}

	// table-driven tests for retrieving tasks
	retrievalTests := []struct {
		name     string
		expected int
	}{
		{
			name:     "retrieve tasks after adding",
			expected: 1, // we expect only one unique task to be added
		},
	}

	for _, tc := range retrievalTests {
		t.Run(tc.name, func(t *testing.T) {
			tasks, err := service.GetTasks()
			if err != nil {
				t.Errorf("GetTasks() error = %v", err)
				return
			}
			if len(tasks) != tc.expected {
				t.Errorf("GetTasks() returned %d tasks, expected %d", len(tasks), tc.expected)
			}
		})
	}
	// Clean up test file after tests
	err := os.Remove("test_tasks.json")
	if err != nil {
		t.Errorf("Failed to clean up test file: %v", err)
	}
}
