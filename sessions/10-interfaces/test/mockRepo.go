package test

import "github.com/CHA0sTIG3R/go-backend-practice/sessions/04-json-tasks/task"

type MockTaskRepository struct {
	Tasks []task.Task
	err   error
}

func (m *MockTaskRepository) GetTasks() ([]task.Task, error) {
	return m.Tasks, m.err
}

func (m *MockTaskRepository) AddTask(newTask task.Task) error {
	if m.err != nil {
		return m.err
	}
	m.Tasks = append(m.Tasks, newTask)
	return nil
}
