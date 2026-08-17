package repository

import (
	"sync"

	"github.com/CHA0sTIG3R/go-backend-practice/sessions/04-json-tasks/task"
)

type MemoryTaskRepository struct {
	tasks map[string]task.Task
	mu	 sync.Mutex
}

func NewMemoryTaskRepository() *MemoryTaskRepository {
	return &MemoryTaskRepository{
		tasks: make(map[string]task.Task),
	}
}

func (r *MemoryTaskRepository) GetTasks() ([]task.Task, error) {
	r.mu.Lock()
	defer r.mu.Unlock()
	var tasks []task.Task
	for _, t := range r.tasks {
		tasks = append(tasks, t)
	}
	return tasks, nil
}

func (r *MemoryTaskRepository) AddTask(newTask task.Task) error {
	r.mu.Lock()
	defer r.mu.Unlock()

	if _, exists := r.tasks[newTask.Name]; exists {
		return &DuplicateTaskError{TaskName: newTask.Name}
	}
	
	r.tasks[newTask.Name] = newTask
	return nil
}

type DuplicateTaskError struct {
	TaskName string
}

func (e *DuplicateTaskError) Error() string {
	return "task with name '" + e.TaskName + "' already exists"
}
