package counter

import "github.com/CHA0sTIG3R/go-backend-practice/sessions/03-task-struct/internal/task"

func Count(tasks []task.Task) int {
	return len(tasks)
}

func IsEmpty(tasks []task.Task) bool {
	return len(tasks) == 0
}

func CountCompleted(tasks []task.Task) int {
	count := 0

	for _, task := range tasks {
		if task.Completed {
			count++
		}
	}

	return count
}

func CountPending(tasks []task.Task) int {
	count := 0

	for _, task := range tasks {
		if !task.Completed {
			count++
		}
	}

	return count
}