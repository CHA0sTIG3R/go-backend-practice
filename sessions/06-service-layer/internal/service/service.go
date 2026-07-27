package service

import (
	"github.com/CHA0sTIG3R/go-backend-practice/sessions/04-json-tasks/task"
	"github.com/CHA0sTIG3R/go-backend-practice/sessions/04-json-tasks/jsonutil"
)

type TaskService struct {
	filename string
}

func NewTaskService(filename string) *TaskService {
	return &TaskService{filename: filename}
}

func (s *TaskService) GetTasks() ([]task.Task, error) {
	tasks, err := jsonutil.LoadTasks(s.filename)
	if err != nil {
		return nil, err
	}
	return tasks, nil
}