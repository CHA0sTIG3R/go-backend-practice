package service

import (
	"github.com/CHA0sTIG3R/go-backend-practice/sessions/04-json-tasks/task"
	"github.com/CHA0sTIG3R/go-backend-practice/sessions/10-interfaces/repository"
)

type TaskService struct {
	repo repository.TaskRepository
}

func NewTaskService(repo repository.TaskRepository) *TaskService {
	return &TaskService{repo: repo}
}

func (s *TaskService) GetTasks() ([]task.Task, error) {
	return s.repo.GetTasks()
}

func (s *TaskService) AddTask(task task.Task) error {
	return s.repo.AddTask(task)
}
