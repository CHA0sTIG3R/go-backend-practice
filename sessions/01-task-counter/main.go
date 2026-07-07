package main

import (
	"github.com/CHA0sTIG3R/go-backend-practice/sessions/01-task-counter/internal/counter"
	"github.com/CHA0sTIG3R/go-backend-practice/sessions/01-task-counter/internal/printer"
)

func main() {
	tasks := []string{
		"Implement login",
		"Write tests",
		"Deploy service",
		"Fix bugs",
	}

	emptyTasks := []string{}

	if counter.IsEmpty(tasks) {
		printer.PrintEmptyTask()
	} else {
		printer.PrintTaskInfo(tasks, counter.Count(tasks))
	}

	if counter.IsEmpty(emptyTasks) {
		printer.PrintEmptyTask()
	} else {
		printer.PrintTaskInfo(emptyTasks, counter.Count(emptyTasks))
	}
}
