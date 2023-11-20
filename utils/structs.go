package utils

type Task struct {
	ID       int
	Content  string
	Complete bool
}

type Taskmanager struct {
	Tasks []Task
}
