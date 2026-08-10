package main

import (
	"log"
	"net/http"

	"github.com/CHA0sTIG3R/go-backend-practice/sessions/10-interfaces/handler"
	"github.com/CHA0sTIG3R/go-backend-practice/sessions/10-interfaces/repository"
	"github.com/CHA0sTIG3R/go-backend-practice/sessions/10-interfaces/service"
)

func main() {
	repo := repository.NewJSONTaskRepository("tasks.json")

	service := service.NewTaskService(repo)

	handler := handler.NewTaskHandler(service)

	http.HandleFunc("/tasks", handler)
	log.Println("Server listening on :8080")
	log.Fatal(http.ListenAndServe(":8080", nil))
}
