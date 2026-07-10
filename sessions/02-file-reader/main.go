package main

import (
	"fmt"

	"github.com/CHA0sTIG3R/go-backend-practice/sessions/02-file-reader/internal/counter"
	"github.com/CHA0sTIG3R/go-backend-practice/sessions/02-file-reader/internal/printer"
	"github.com/CHA0sTIG3R/go-backend-practice/sessions/02-file-reader/internal/reader"
)

func main() {
	tasks, err := reader.ReadTasks("tasks.txt")

	if err != nil {
		fmt.Println(err)
		return
	}

	if counter.IsEmpty(tasks) {
		printer.PrintEmptyTask()
	} else {
		printer.PrintTaskInfo(tasks, counter.Count(tasks))
	}

}
