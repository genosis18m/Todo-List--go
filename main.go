package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

type Task struct {
	TaskName  string
	completed bool
}

var tasks []Task // array that stores the tasks

// Function to add a task
func addTask(task string) {
	newTask := Task{TaskName: task, completed: false} // Create a new task
	tasks = append(tasks, newTask)                    // Add it to the slice
	fmt.Println("Task added successfully.")
	listTasks() // Automatically list tasks after adding
}

// Function to list all tasks
func listTasks() {
	if len(tasks) == 0 {
		fmt.Println("No tasks available.")
		return
	}
	fmt.Println("\nCurrent Tasks:")
	for i, task := range tasks {
		status := "not completed"
		if task.completed {
			status = "completed"
		}
		fmt.Printf("%d. %s [%s]\n", i+1, task.TaskName, status)
	}
}

// Function to mark a task as completed
func markCompleted(index int) {
	if index >= 1 && index <= len(tasks) {
		tasks[index-1].completed = true // index starts from 0
		fmt.Println("Task marked as completed.")
		listTasks() // Automatically list tasks after completion
	} else {
		fmt.Println("Invalid index.")
	}
}

// Function to edit a task
func editTask(index int, newTask string) {
	if index >= 1 && index <= len(tasks) {
		tasks[index-1].TaskName = newTask
		fmt.Println("Task updated successfully.")
		listTasks() // Automatically list tasks after editing
	} else {
		fmt.Println("Invalid index.")
	}
}

// Function to delete a task
func deleteTask(index int) {
	if index >= 1 && index <= len(tasks) {
		tasks = append(tasks[:index-1], tasks[index:]...)
		fmt.Println("Task deleted.")
		listTasks() // Automatically list tasks after deletion
	} else {
		fmt.Println("Invalid index.")
	}
}

func main() {

	var indexInput int
	var taskInput, newTaskInput string
	scanner := bufio.NewScanner(os.Stdin)

	for {
		fmt.Println("\nOptions:")
		fmt.Println("1. Add Task")
		fmt.Println("2. List Tasks")
		fmt.Println("3. Mark Task as Completed")
		fmt.Println("4. Edit Task")
		fmt.Println("5. Delete Task")
		fmt.Println("6. Exit")
		fmt.Print("Enter your choice (1-6): ")

		scanner.Scan()
		input := scanner.Text()
		choice, err := strconv.Atoi(input)
		if err != nil {
			fmt.Println("Invalid input. Please enter a number.")
			continue
		}

		switch choice {
		case 1:
			fmt.Print("Enter task: ")
			scanner.Scan()
			taskInput = scanner.Text()
			addTask(taskInput)
		case 2:
			listTasks() // Explicitly list tasks on demand
		case 3:
			fmt.Print("Enter the task number to mark as completed: ")
			scanner.Scan()
			input = scanner.Text()
			indexInput, err = strconv.Atoi(input)
			if err != nil {
				fmt.Println("Invalid input.")
				continue
			}
			markCompleted(indexInput)
		case 4:
			fmt.Print("Enter the task number to edit: ")
			scanner.Scan()
			input = scanner.Text()
			indexInput, err = strconv.Atoi(input)
			if err != nil {
				fmt.Println("Invalid input.")
				continue
			}
			fmt.Print("Enter the new task: ")
			scanner.Scan()
			newTaskInput = scanner.Text()
			editTask(indexInput, newTaskInput)
		case 5:
			fmt.Print("Enter the task number to delete: ")
			scanner.Scan()
			input = scanner.Text()
			indexInput, err = strconv.Atoi(input)
			if err != nil {
				fmt.Println("Invalid input.")
				continue
			}
			deleteTask(indexInput)
		case 6:
			fmt.Println("Exiting.")
			return
		default:
			fmt.Println("Invalid choice. Please choose between 1-6.")
		}
	}
}
