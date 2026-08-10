package repository

import (
	"encoding/json"

	"github.com/CHA0sTIG3R/go-backend-practice/sessions/04-json-tasks/task"
	"github.com/CHA0sTIG3R/go-backend-practice/sessions/07-post-task/jsonutil"
)

type JSONTaskRepository struct {
	filename string
}

func NewJSONTaskRepository(filename string) *JSONTaskRepository {
	return &JSONTaskRepository{filename: filename}
}

func (r *JSONTaskRepository) GetTasks() ([]task.Task, error) {
	tasks, err := jsonutil.LoadTasks(r.filename)
	if err != nil {
		return nil, err
	}
	return tasks, nil
}

func (r *JSONTaskRepository) AddTask(newTask task.Task) error {
	jsonTask, err := json.Marshal(newTask)
	if err != nil {
		return err
	}

	err = jsonutil.SaveTasks(r.filename, jsonTask)
	if err != nil {
		// check the error type and raise 409 conflict if it's a duplicate task error
		if _, ok := err.(*jsonutil.DuplicateTaskError); ok {
			return &jsonutil.DuplicateTaskError{TaskName: newTask.Name}
		}
		return err
	}
	return nil
}
