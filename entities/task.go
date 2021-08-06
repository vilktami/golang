package entities

type Task struct {
	Title string
	Done bool
}

type NewTaskTodo struct {
	Task string `json:"task"`
}