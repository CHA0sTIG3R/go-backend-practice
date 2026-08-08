package jsonutil

import (
	"bytes"
	"encoding/json"
	"fmt"
	"os"

	"github.com/CHA0sTIG3R/go-backend-practice/sessions/04-json-tasks/task"
)

type DuplicateTaskError struct {
	TaskName string
}

func (e *DuplicateTaskError) Error() string {
	return fmt.Sprintf("task with name '%s' already exists", e.TaskName)
}

func EncodeJson(tasks []task.Task) ([]byte, error) {
	jsonTasks, err := json.MarshalIndent(tasks, "", " ")

	if err != nil {
		return nil, err
	}

	return jsonTasks, nil
}

func DecodeJson(data []byte) ([]task.Task, error) {
	var tasks []task.Task

	err := json.Unmarshal(data, &tasks)

	if err != nil {
		return nil, err
	}

	return tasks, nil
}

func decodeTasks(data []byte) ([]task.Task, error) {
	trimmed := bytes.TrimSpace(data)
	if len(trimmed) == 0 {
		return []task.Task{}, nil
	}

	tasks := []task.Task{}
	if trimmed[0] == '[' {
		if err := json.Unmarshal(trimmed, &tasks); err != nil {
			return nil, err
		}
		return tasks, nil
	}

	var single task.Task
	if err := json.Unmarshal(trimmed, &single); err != nil {
		return nil, err
	}
	return []task.Task{single}, nil
}

func SaveTasks(filename string, data []byte) error {
	incomingTasks, err := decodeTasks(data)
	if err != nil {
		return err
	}

	fileBytes, err := os.ReadFile(filename)
	if err != nil && !os.IsNotExist(err) {
		return err
	}

	existingTasks, err := decodeTasks(fileBytes)
	if err != nil {
		return err
	}

	// check for duplicates based on task name
	existingTaskNames := make(map[string]bool)
	for _, task := range existingTasks {
		existingTaskNames[task.Name] = true
	}

	for _, task := range incomingTasks {
		if existingTaskNames[task.Name] {
			return &DuplicateTaskError{TaskName: task.Name}
		}
	}

	existingTasks = append(existingTasks, incomingTasks...)

	encodedData, err := EncodeJson(existingTasks)
	if err != nil {
		return err
	}

	return os.WriteFile(filename, encodedData, 0644)
}

func LoadTasks(filename string) ([]task.Task, error) {
	fileBytes, err := os.ReadFile(filename)
	if err != nil {
		return nil, err
	}

	return DecodeJson(fileBytes)
}
