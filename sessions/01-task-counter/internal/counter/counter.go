package counter

func Count(task []string) int {
	return len(task)
}

func IsEmpty(task []string) bool{
	return Count(task) < 1
}