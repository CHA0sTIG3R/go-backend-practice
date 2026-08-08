package jsonutil

import (
	"bytes"
	"encoding/json"
	"errors"
	"io"
	"os"

	"github.com/CHA0sTIG3R/go-backend-practice/sessions/04-json-tasks/task"
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
	cleanedData := bytes.TrimSpace(data)
	
	incomingTasks := []task.Task{}
	if cleanedData[0] == '[' {
		err := json.Unmarshal(cleanedData, &incomingTasks)
		if err != nil {
			return err
		}
	} else {
		var singleTask task.Task
		err := json.Unmarshal(cleanedData, &singleTask)
		if err != nil {
			return err
		}
		incomingTasks = append(incomingTasks, singleTask)
	}

	existingTasks := []task.Task{}

	file, err := os.OpenFile(filename, os.O_RDWR|os.O_CREATE, 0644)
	if err != nil {
		return err
	}

	defer file.Close()

	decodedData := json.NewDecoder(file)
	if err := decodedData.Decode(&existingTasks); err != nil && !errors.Is(err, io.EOF) {
		return err
	}

	println("Existing tasks before appending:", existingTasks)

	existingTasks = append(existingTasks, incomingTasks...)
	println("Existing tasks after appending:", existingTasks)

	encodedData, err := EncodeJson(existingTasks)
	if err != nil {
		return err
	}

	println("Encoded data to be written:", string(encodedData))

	err = os.WriteFile(filename, encodedData, 0644)
	if err != nil {
		return err
	}

	println("Tasks saved successfully to", filename)

	return nil

}

func LoadTasks(filename string) ([]task.Task, error) {
	fileBytes, err := os.ReadFile(filename)
	if err != nil {
		return nil, err
	}

	return DecodeJson(fileBytes)
}
