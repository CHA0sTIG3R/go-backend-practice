package counter

func Count(tasks []string) int {
	return len(tasks)
}

func IsEmpty(tasks []string) bool {
	return len(tasks) == 0
}
