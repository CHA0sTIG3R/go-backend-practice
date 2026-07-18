package task

type Task struct {
	Name      string `json:"name"`
	Completed bool   `json:"completed"`
	Priority  int    `json:"priority"`
}
