package repository

import "github.com/CHA0sTIG3R/go-backend-practice/sessions/04-json-tasks/task"

type TaskRepository interface {
	GetTasks() ([]task.Task, error)
	AddTask(task.Task) error
}
