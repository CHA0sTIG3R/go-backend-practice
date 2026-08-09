package validation

import (
	"fmt"
	"strings"

	"github.com/CHA0sTIG3R/go-backend-practice/sessions/04-json-tasks/task"
)

func ValidateTask(task task.Task) error {
	if strings.TrimSpace(task.Name) == "" {
		return fmt.Errorf("Name is required")
	}
	if task.Priority < 1 || task.Priority > 3 {
		return fmt.Errorf("Priority must be between 1 and 3")
	}

	if len(task.Name) > 100 {
		return fmt.Errorf("Name cannot exceed 100 characters")
	}

	return nil
}
