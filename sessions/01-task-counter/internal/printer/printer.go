package printer

import "fmt"

func PrintEmptyTask () {
	fmt.Println("No tasks available.")
}

func PrintTaskInfo(task []string, count int) {
	fmt.Println("Tasks: ")
	for _, task := range task {
		fmt.Println("- ", task)
	}

	fmt.Println("Total tasks:", count)
}