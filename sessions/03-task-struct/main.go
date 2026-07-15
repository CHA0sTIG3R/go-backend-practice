package main

import (
	"github.com/CHA0sTIG3R/go-backend-practice/sessions/03-task-struct/internal/counter"
	"github.com/CHA0sTIG3R/go-backend-practice/sessions/03-task-struct/internal/printer"
	"github.com/CHA0sTIG3R/go-backend-practice/sessions/03-task-struct/internal/task"
)

func main() {
	tasks := []task.Task {
		{
			Name: "Write tests",
			Completed: false,
			Priority: 2,
		},
		{
			Name: "Review PR",
			Completed: true,
			Priority: 1,
		},
		{
			Name: "Update documentation",
			Completed: false,
			Priority: 3,
		},
		{
			Name: "Deploy to production",
			Completed: true,
			Priority: 1,
		},
		{
			Name: "Refactor code",
			Completed: false,
			Priority: 2,
		},
		{
			Name: "Fix bugs",
			Completed: true,
			Priority: 1,
		},
		{
			Name: "Implement new feature",
			Completed: false,
			Priority: 3,
		},
		{
			Name: "Conduct code review",
			Completed: true,
			Priority: 2,
		},
	}

	n := len(tasks)

	for i := 0; i < n-1; i++ {
		for j := 0; j < n-i-1; j++{
			if tasks[j].Priority > tasks[j+1].Priority {
				tasks[j], tasks[j+1] = tasks[j+1], tasks[j]
			}
		}
	}

	if counter.IsEmpty(tasks) {
		printer.PrintEmptyTask()
	} else {
		printer.PrintTaskInfo(tasks, counter.CountCompleted(tasks), counter.CountPending(tasks), counter.Count(tasks))
	}
}