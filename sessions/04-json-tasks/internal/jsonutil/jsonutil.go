package jsonutil

import (
	"encoding/json"
	"os"

	"github.com/CHA0sTIG3R/go-backend-practice/sessions/04-json-tasks/internal/task"
)

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

func SaveTasks(filename string, data []byte) error {
	file, err := os.Create(filename)
	if err != nil {
		return err
	}

	defer file.Close()

	_, err = file.Write(data)
	if err != nil {
		return err
	}

	return nil
}

func LoadTasks(filename string) ([]task.Task, error) {
	fileBytes, err := os.ReadFile(filename)
	if err != nil {
		return nil, err
	}

	return DecodeJson(fileBytes)
}
