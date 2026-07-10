package reader

import (
	"bufio"
	"os"
	"strings"
)

func ReadTasks(filename string) ([]string, error) {

	tasks := []string{}

	file, err := os.Open(filename)
	if err != nil {
		return nil, err
	}

	defer file.Close()

	scanner := bufio.NewScanner(file)

	for scanner.Scan() {
		line := scanner.Text()

		if strings.TrimSpace(line) != "" {
			tasks = append(tasks, line)
		}

	}

	if err := scanner.Err(); err != nil {
		return nil, err
	} else {
		return tasks, nil
	}
}
