package repository

import "github.com/CHA0sTIG3R/go-backend-practice/sessions/04-json-tasks/task"

type MemoryTaskRepository struct {
	tasks []task.Task
}

func NewMemoryTaskRepository() *MemoryTaskRepository {
	return &MemoryTaskRepository{
		tasks: []task.Task{},
	}
}

func (r *MemoryTaskRepository) GetTasks() ([]task.Task, error) {
	return r.tasks, nil
}

func (r *MemoryTaskRepository) AddTask(newTask task.Task) error {
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
