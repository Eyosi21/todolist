package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strings"
)

var tasks = make(map[int]string)

func addTasks() {
	fmt.Println("Enter to do list")
	reader := bufio.NewReader(os.Stdin)
	var lines []string
	for {
		line, err := reader.ReadString('\n')
		if err != nil {
			log.Fatal(err)
		}
		if len(strings.TrimSpace(line)) == 0 {
			break
		}
		lines = append(lines, line)
	}

	fmt.Println("Task added:")
	for _, l := range lines {
		fmt.Println(l)
	}

	if len(lines) > 0 {
		taskDescription := strings.Join(lines, "\n")
		tasks[len(tasks)] = taskDescription
		fmt.Println("Task added:")
		fmt.Println(taskDescription)
	} else {
		fmt.Println("No task entered, returning to actions.")
	}
}

func markDone() {
	var num int
	fmt.Println("Enter to task number")
	fmt.Scan(&num)
	if task, exists := tasks[num]; exists {
		tasks[num] = task + "--- Done"
	} else {
		fmt.Println("Task not found")
	}
}

func viewTasks() {
	if len(tasks) == 0 {
		fmt.Println("There are no tasks available")

	} else {
		for i, task := range tasks {
			fmt.Printf("Task %d: %s\n", i+1, task)
		}
	}
}

func deleteTask() {
	var num int
	fmt.Println("Enter to task number")
	fmt.Scan(&num)
	if _, exists := tasks[num]; exists {
		delete(tasks, num)
		fmt.Println("Task deleted Successfully")
	} else {
		fmt.Println("Task not found")
	}
}

func actionsTab() {
	var action int
	for {
		fmt.Println("Actions: ")
		fmt.Println("1. Add Task")
		fmt.Println("2. View Task")
		fmt.Println("3. Delete Task")
		fmt.Println("4. Mark done ")
		fmt.Println("5. Exit ")
		fmt.Println("Please choose action and enter number")
		fmt.Scan(&action)

		switch action {
		case 1:
			addTasks()
		case 2:
			viewTasks()
		case 3:
			deleteTask()
		case 4:
			markDone()
		case 5:
			fmt.Println("Exiting the program.")
			return
		default:
			fmt.Println("Enter correct number")
			continue

		}

		fmt.Println("Action completed enter another action or exit")
	}

}

func main() {
	actionsTab()

}
