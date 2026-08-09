package test

import (
	"testing"

	"github.com/CHA0sTIG3R/go-backend-practice/sessions/04-json-tasks/task"
	"github.com/CHA0sTIG3R/go-backend-practice/sessions/07-post-task/validation"
)

func TestValidation(t *testing.T) {
	test := []struct {
		name    string
		task    task.Task
		wantERR bool
	}{
		{
			name: "valid task",
			task: task.Task{
				Name:     "Practice Go",
				Priority: 2,
			},
			wantERR: false,
		},
		{
			name: "empty name",
			task: task.Task{
				Name:     " ",
				Priority: 2,
			},
			wantERR: true,
		},
		{
			name: "priority too high",
			task: task.Task{
				Name:     "Deploy",
				Priority: 5,
			},
			wantERR: true,
		},
		{
			name: "priority too low",
			task: task.Task{
				Name:     "Refactor",
				Priority: 0,
			},
			wantERR: true,
		},
	}

	for _, tc := range test {
		t.Run(tc.name, func(t *testing.T) {
			err := validation.ValidateTask(tc.task)
			if (err != nil) != tc.wantERR {
				t.Errorf("ValidateTask() error = %v, wantERR %v", err, tc.wantERR)
			}
		})
	}
}
