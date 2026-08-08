package service

import (
	"encoding/json"

	"github.com/CHA0sTIG3R/go-backend-practice/sessions/04-json-tasks/jsonutil"
	"github.com/CHA0sTIG3R/go-backend-practice/sessions/04-json-tasks/task"
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

func (s *TaskService) AddTask(task task.Task) (error){

	jsonTask, err := json.Marshal(task)
	if err != nil {
		return err
	}

	err = jsonutil.SaveTasks(s.filename, jsonTask)
	if err != nil {
		return err
	}

	return nil
}