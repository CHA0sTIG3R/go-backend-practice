package main

import (
	"log"
	"net/http"

	"github.com/CHA0sTIG3R/go-backend-practice/sessions/07-post-task/handler"
	"github.com/CHA0sTIG3R/go-backend-practice/sessions/07-post-task/service"
)

func main() {
	service := service.NewTaskService("tasks.json")
	handler := handler.NewTaskHandler(service)

	http.HandleFunc("/tasks", handler)
	log.Println("Server listening on :8080")
	log.Fatal(http.ListenAndServe(":8080", nil))
}
