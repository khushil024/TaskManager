package main

import (
	"TaskManager/utils"
	"bufio"
	"fmt"
	"os"
	"strings"
)

func main() {
	taskManager := utils.Taskmanager{}

	for {

		fmt.Print("Enter command (add/list/complete/remove/exit): ")
		scanner := bufio.NewScanner(os.Stdin)
		scanner.Scan()
		command := scanner.Text()

		switch strings.ToLower(command) {

		case "add":
			fmt.Print("Enter task: ")
			scanner.Scan()
			content := scanner.Text()
			taskManager.AddTask(content)

		case "list":
			taskManager.ListTasks()

		case "complete":
			fmt.Print("Enter task ID to mark as complete: ")
			scanner.Scan()
			taskID := parseTaskID(scanner.Text())
			taskManager.CompleteTask(taskID)

		case "remove":
			fmt.Print("Enter task ID to remove: ")
			scanner.Scan()
			taskID := parseTaskID(scanner.Text())
			taskManager.RemoveTask(taskID)

		case "exit":
			fmt.Println("Exiting TaskManager. Goodbye!")
			os.Exit(0)

		default:
			fmt.Println("Invalid command. Try again.")
		}
	}

}

func parseTaskID(input string) int {
	taskID := 0

	_, err := fmt.Sscanf(input, "%d", &taskID)

	if err != nil {
		fmt.Println("Invalid input. Please enter a valid task ID.")
		return 0
	}
	return taskID
}
