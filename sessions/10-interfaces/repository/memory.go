package repository

import (
	"sync"

	"github.com/CHA0sTIG3R/go-backend-practice/sessions/04-json-tasks/task"
)

type MemoryTaskRepository struct {
	tasks []task.Task
	mu	 sync.Mutex
}

func NewMemoryTaskRepository() *MemoryTaskRepository {
	return &MemoryTaskRepository{
		tasks: []task.Task{},
	}
}

func (r *MemoryTaskRepository) GetTasks() ([]task.Task, error) {
	r.mu.Lock()
	defer r.mu.Unlock()
	return append([]task.Task{}, r.tasks...), nil
}

func (r *MemoryTaskRepository) AddTask(newTask task.Task) error {
	r.mu.Lock()
	defer r.mu.Unlock()

	for _, t := range r.tasks {
		if t.Name == newTask.Name {
			return &DuplicateTaskError{TaskName: newTask.Name}
		}
	}
	
	r.tasks = append(r.tasks, newTask)
	return nil
}

type DuplicateTaskError struct {
	TaskName string
}

func (e *DuplicateTaskError) Error() string {
	return "task with name '" + e.TaskName + "' already exists"
}
