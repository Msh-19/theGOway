package main

import (
	"fmt"
	"io/ioutil"
	"os"
	"strconv"
	"strings"
)

const todoFile = "tasks.txt"

func main() {
	if len(os.Args) < 2 {
		fmt.Println("Usage: todo add <task> | list | done <task number>")
		return
	}
	command := os.Args[1]

	switch command {
	case "add":
		if len(os.Args) < 3 {
			fmt.Println("Please provide a task to add.")
			return
		}
		task := strings.Join(os.Args[2:], " ")
		addTask(task)
	case "list":
		listTasks()
	case "done":
		if len(os.Args) < 3 {
			fmt.Println("Please provide the task number to mark as done.")
			return
		}
		taskNum, err := strconv.Atoi(os.Args[2])
		if err != nil {
			fmt.Println("Invalid task number.")
			return
		}
		markTaskDone(taskNum)
	default:
		fmt.Println("Unknown command:", command)
	}
}

func addTask(task string) {
	file, err := os.OpenFile(todoFile, os.O_APPEND|os.O_WRONLY|os.O_CREATE, 0600)
	if err != nil {
		fmt.Println("Error opening tasks file:", err)
		return
	}
	defer file.Close()

	_, err = fmt.Fprintln(file, "[ ] "+task)
	if err != nil {
		fmt.Println("Error writing to tasks file:", err)
		return
	}
	fmt.Println("Task added:", task)
}

func listTasks() {
	data, err := ioutil.ReadFile(todoFile)
	if err != nil {
		if os.IsNotExist(err) {
			fmt.Println("No tasks found.")
		} else {
			fmt.Println("Error reading tasks file:", err)
		}
		return
	}

	tasks := strings.Split(string(data), "\n")
	for i, task := range tasks {
		if task == "" {
			continue
		}
		fmt.Printf("%d. %s\n", i+1, task)
	}
}

func markTaskDone(taskNum int) {
	data, err := ioutil.ReadFile(todoFile)
	if err != nil {
		fmt.Println("Error reading tasks file:", err)
		return
	}

	tasks := strings.Split(string(data), "\n")
	if taskNum <= 0 || taskNum > len(tasks) {
		fmt.Println("Invalid task number.")
		return
	}

	tasks[taskNum-1] = strings.Replace(tasks[taskNum-1], "[ ]", "[x]", 1)

	err = ioutil.WriteFile(todoFile, []byte(strings.Join(tasks, "\n")), 0600)
	if err != nil {
		fmt.Println("Error writing tasks file:", err)
		return
	}

	fmt.Println("Task", taskNum, "marked as done.")
}
