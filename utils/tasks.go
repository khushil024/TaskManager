package utils

import "fmt"

func (tm *Taskmanager) AddTask(content string) {
	task := Task{
		ID:       len(tm.Tasks) + 1,
		Content:  content,
		Complete: false,
	}
	tm.Tasks = append(tm.Tasks, task)
	fmt.Printf("Task added: %s\n", content)
}

func (tm *Taskmanager) ListTasks() {
	fmt.Println("Tasks:")

	for _, task := range tm.Tasks {
		status := "Incomplete"
		if task.Complete {
			status = "Complete"
		}
		fmt.Printf("%d. %s (%s)\n", task.ID, task.Content, status)
	}
}

func (tm *Taskmanager) CompleteTask(taskID int) {
	if taskID > 0 && taskID <= len(tm.Tasks) {
		tm.Tasks[taskID-1].Complete = true
		fmt.Printf("Task %d marked as complete\n", taskID)
	} else {
		fmt.Println("Invalid task ID")
	}
}

func (tm *Taskmanager) RemoveTask(taskID int) {
	if taskID > 0 && taskID <= len(tm.Tasks) {
		fmt.Println("Yet to be implemented")
		//remove index `taskID`
	} else {
		fmt.Println("Invalid task ID")
	}
}
