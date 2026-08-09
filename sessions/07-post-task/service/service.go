package service

import (
	"encoding/json"
	"log"

	"github.com/CHA0sTIG3R/go-backend-practice/sessions/04-json-tasks/task"
	"github.com/CHA0sTIG3R/go-backend-practice/sessions/07-post-task/jsonutil"
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

func (s *TaskService) AddTask(task task.Task) error {

	jsonTask, err := json.Marshal(task)
	if err != nil {
		return err
	}

	err = jsonutil.SaveTasks(s.filename, jsonTask)
	if err != nil {
		// check the error type and raise 409 conflict if it's a duplicate task error
		if _, ok := err.(*jsonutil.DuplicateTaskError); ok {
			return &jsonutil.DuplicateTaskError{TaskName: task.Name}
		}

		log.Println("Error saving tasks:", err)
		return err
	}

	return nil
}
