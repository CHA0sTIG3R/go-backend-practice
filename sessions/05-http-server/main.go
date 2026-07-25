package main

import (
	"log"
	"net/http"

	"github.com/CHA0sTIG3R/go-backend-practice/sessions/04-json-tasks/task"
	"github.com/CHA0sTIG3R/go-backend-practice/sessions/05-http-server/handler"
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
	}

	
	http.HandleFunc("/tasks", handler.TasksHandler(tasks))
	log.Println("Server listening on :8080")
	log.Fatal(http.ListenAndServe(":8080", nil))

}

