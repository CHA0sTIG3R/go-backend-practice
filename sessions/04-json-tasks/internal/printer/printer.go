package printer

import (
	"fmt"

	"github.com/CHA0sTIG3R/go-backend-practice/sessions/04-json-tasks/internal/task"
)

func PrintEmptyTask() {
	fmt.Println("No tasks available.")
}

func PrintTaskInfo(task []task.Task, completed int, pending int, total int) {

	fmt.Println("Tasks:", total)
	fmt.Println("Completed:", completed)
	fmt.Println("Pending:", pending)
	for _, task := range task {
		if task.Completed {
			fmt.Printf("[x] %s (Priority %d) \n", task.Name, task.Priority)
		} else {
			fmt.Printf("[ ] %s (Priority %d) \n", task.Name, task.Priority)
		}

	}
}

func PrintTaskJson(jsonstr string) {
	fmt.Println(jsonstr)
}
