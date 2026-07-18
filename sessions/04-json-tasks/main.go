package main

import (
	"fmt"

	"github.com/CHA0sTIG3R/go-backend-practice/sessions/04-json-tasks/internal/counter"
	"github.com/CHA0sTIG3R/go-backend-practice/sessions/04-json-tasks/internal/jsonutil"
	"github.com/CHA0sTIG3R/go-backend-practice/sessions/04-json-tasks/internal/printer"
	"github.com/CHA0sTIG3R/go-backend-practice/sessions/04-json-tasks/internal/task"
)

func main() {
	tasks := []task.Task{
		{
			Name:      "Write tests",
			Completed: false,
			Priority:  2,
		},
		{
			Name:      "Review PR",
			Completed: true,
			Priority:  1,
		},
		{
			Name:      "Update documentation",
			Completed: false,
			Priority:  3,
		},
		{
			Name:      "Deploy to production",
			Completed: true,
			Priority:  1,
		},
		{
			Name:      "Refactor code",
			Completed: false,
			Priority:  2,
		},
		{
			Name:      "Fix bugs",
			Completed: true,
			Priority:  1,
		},
		{
			Name:      "Implement new feature",
			Completed: false,
			Priority:  3,
		},
		{
			Name:      "Conduct code review",
			Completed: true,
			Priority:  2,
		},
	}

	data := `
	[
		{
			"name":"Deploy",
			"completed":false,
			"priority":1
		},
		{
			"name":"Review",
			"completed":true,
			"priority":2
		}
	]
	`

	jsonTasks, err := jsonutil.EncodeJson(tasks)
	if err != nil {
		fmt.Println(err)
		return
	}

	printer.PrintTaskJson(string(jsonTasks))

	decodedTasks, err := jsonutil.DecodeJson([]byte(data))
	if err != nil {
		fmt.Println(err)
		return
	}

	printer.PrintTaskInfo(
		decodedTasks,
		counter.CountCompleted(decodedTasks),
		counter.CountPending(decodedTasks),
		counter.Count(decodedTasks),
	)

	err = jsonutil.SaveTasks("tasks.json", jsonTasks)
	if err != nil {
		fmt.Println(err)
		return
	}
	fileLoadedTasks, err := jsonutil.LoadTasks("tasks.json")
	if err != nil {
		fmt.Println(err)
		return
	}

	printer.PrintTaskInfo(
		fileLoadedTasks,
		counter.CountCompleted(fileLoadedTasks),
		counter.CountPending(fileLoadedTasks),
		counter.Count(fileLoadedTasks),
	)
}
