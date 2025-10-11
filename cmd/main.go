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

func main() {
	newTask := Task{}
	description := flag.String("add", "", "Add a new task to the list")
	flag.Parse()
	newTask.description = *description
	newTask.id++
	newTask.status = "todo"
	newTask.createdAt = time.Now()
	newTask.updatedAt = time.Now()

	// For demonstration purposes, print the new task details
	fmt.Println("Id: ", newTask.id)
	fmt.Println("Description: ", newTask.description)
	fmt.Println("Status: ", newTask.status)
	fmt.Println("CreateAt: ", newTask.createdAt)
	fmt.Println("UpdatedAt: ", newTask.updatedAt)

}
