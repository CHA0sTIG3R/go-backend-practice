package test

import (
	"testing"

	"github.com/CHA0sTIG3R/go-backend-practice/sessions/04-json-tasks/task"
	"github.com/CHA0sTIG3R/go-backend-practice/sessions/10-interfaces/service"
)

func TestTaskService(t *testing.T) {
	mockRepo := &MockTaskRepository{
		Tasks: []task.Task{},
		err:   nil,
	}

	taskService := service.NewTaskService(mockRepo)

	t.Run("AddTask should add a task to the repository", func(t *testing.T) {
		newTask := task.Task{Name: "Test Task", Priority: 1}
		err := taskService.AddTask(newTask)
		if err != nil {
			t.Fatalf("expected no error, got %v", err)
		}

		tasks, err := taskService.GetTasks()
		if err != nil {
			t.Fatalf("expected no error, got %v", err)
		}

		if len(tasks) != 1 || tasks[0].Name != "Test Task" {
			t.Fatalf("expected task to be added, got %v", tasks)
		}
	})

	t.Run("GetTasks should return all tasks from the repository", func(t *testing.T) {
		tasks, err := taskService.GetTasks()
		if err != nil {
			t.Fatalf("expected no error, got %v", err)
		}

		if len(tasks) != 1 || tasks[0].Name != "Test Task" {
			t.Fatalf("expected to get the added task, got %v", tasks)
		}
	})
}
