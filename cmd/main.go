package main

import (
	"encoding/json"
	"flag"
	"fmt"
	"os"
	"time"
)

// Data File
const dataFile = "data.json"

// Declare Task Struct
type Task struct {
	Id          int       `json:"id"`
	Description string    `json:"description"`
	Status      string    `json:"status"`
	CreatedAt   time.Time `json:"created_at"`
	UpdatedAt   time.Time `json:"updated_at"`
	NextId      int       `json:"nextid"` // -> TaskList 단위로 관리 필요
}

// Buffered Task List
type TaskList []Task

// Read Json file to TaskList
func loadTasks(filename string) (TaskList, error) {
	// Return empty list if file doesn't exist
	data, err := os.ReadFile(filename)
	if err != nil {
		if os.IsNotExist(err) {
			return TaskList{}, nil
		}
		return nil, err
	}

	var tasks TaskList
	err = json.Unmarshal(data, &tasks)
	if err != nil {
		return nil, err
	}
	return tasks, nil
}

// Save TaskList to json file
func saveTasks(filename string, tl TaskList) error {
	data, err := json.MarshalIndent(tl, "", "  ")
	if err != nil {
		return err
	}
	return os.WriteFile(filename, data, 0644)
}

// Add new task with incremented ID
func (tl *TaskList) addTask(description string) {
	now := time.Now()
	var id, nextId int
	if len(*tl) == 0 {
		id = 1
		nextId = 2
	} else {
		id = (*tl)[len(*tl)-1].NextId
		nextId = (*tl)[len(*tl)-1].NextId + 1
	}
	task := Task{
		Id:          id,
		Description: description,
		Status:      "todo",
		CreatedAt:   now,
		UpdatedAt:   now,
		NextId:      nextId,
	}
	*tl = append(*tl, task)
}

// list
func (tl *TaskList) listTasks() {
	for _, task := range *tl {
		fmt.Printf("ID: %d, Description: %s, Status: %s, CreatedAt: %s, UpdatedAt: %s\n", task.Id, task.Description, task.Status, task.CreatedAt, task.UpdatedAt)
	}
}

// update
func (tl *TaskList) updateTask(id int, description string, status string) {
	// Implementation for updating a task
}

// delete
func (tl *TaskList) deleteTask(id int) {
	// Implementation for deleting a task
}

func main() {
	// positional argument parsing
	flag.Parse()
	var operation string = flag.Args()[0]
	var description string
	if operation == "add" && len(flag.Args()) > 1 {
		description = flag.Args()[1]
	}

	// load existing tasks from file
	tempTaskList, err := loadTasks(dataFile)
	if err != nil {
		fmt.Println("Error loading tasks:", err)
		return
	}

	// Perform operation based on the command
	switch operation {
	case "add":
		// Add task
		tempTaskList.addTask(description)
		err := saveTasks(dataFile, tempTaskList)
		if err != nil {
			fmt.Println("Error saving tasks:", err)
			return
		}
		fmt.Printf("Output: Task added successfully (ID: %d)\n", len(tempTaskList))
	case "list":
		// List tasks
		tempTaskList.listTasks()
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
}
