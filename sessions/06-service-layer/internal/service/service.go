package service

import "github.com/CHA0sTIG3R/go-backend-practice/sessions/04-json-tasks/task"

type TaskService struct {
	filename string
}

func NewTaskService(filename string) *TaskService 

func (s *TaskService) GetTasks() ([]task.Task, error) 