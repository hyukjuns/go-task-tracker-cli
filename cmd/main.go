package main

import (
	"flag"
	"fmt"
	"time"
)

type Task struct {
	id          int
	description string
	status      string
	createdAt   time.Time
	updatedAt   time.Time
}

func (t *Task) addTask(task_list *[]Task, description string) {
	// Add a new task to the list
	t.id = len(*task_list) + 1
	t.description = description
	t.status = "todo"
	t.createdAt = time.Now()
	t.updatedAt = time.Now()

	*task_list = append(*task_list, *t)
}

func main() {
	newTaskList := make([]Task, 0) // -> json 파일로 저장해야함
	flag.Parse()
	var operation string = flag.Args()[0]
	switch operation {
	case "add":
		fmt.Println("add")
		var newTask Task
		description := flag.Args()[1]
		fmt.Println(description)
		newTask.addTask(&newTaskList, description)
		fmt.Println(newTaskList)
	case "list":
		// List tasks
		fmt.Println("list")
	case "update":
		// Update task
		fmt.Println("update")
	case "delete":
		// Delete task
		fmt.Println("delete")
	default:
		// Invalid operation
		fmt.Println("default")
	}

	// // For demonstration purposes, print the new task details
	// fmt.Println("Id: ", newTask.id)
	// fmt.Println("Description: ", newTask.description)
	// fmt.Println("Status: ", newTask.status)
	// fmt.Println("CreateAt: ", newTask.createdAt)
	// fmt.Println("UpdatedAt: ", newTask.updatedAt)

}
