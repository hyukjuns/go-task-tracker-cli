package main

import (
	"encoding/json"
	"errors"
	"flag"
	"fmt"
	"os"
	"strconv"
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

// list all tasks
func (tl *TaskList) listTasks(filter string) error {
	if len(*tl) == 0 {
		fmt.Println("No tasks found.")
		return nil
	}
	if filter == "all" {
		for _, task := range *tl {
			fmt.Printf("ID: %d, Description: %s, Status: %s, CreatedAt: %s, UpdatedAt: %s\n", task.Id, task.Description, task.Status, task.CreatedAt, task.UpdatedAt)
		}
	} else if filter == "todo" || filter == "in-progress" || filter == "done" {
		for _, task := range *tl {
			if task.Status == filter {
				fmt.Printf("ID: %d, Description: %s, Status: %s, CreatedAt: %s, UpdatedAt: %s\n", task.Id, task.Description, task.Status, task.CreatedAt, task.UpdatedAt)
			}
		}
	} else {
		return errors.New("Invalid filter. Use 'todo', 'in-progress', or 'done'.")
	}
	return nil
}

// Add new task with incremented ID
func (tl *TaskList) addTask(description string) error {
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
	err := saveTasks(dataFile, *tl)
	if err != nil {
		fmt.Println("Error saving tasks:", err)
		return err
	}
	return nil
}

// update
func (tl *TaskList) updateTask(id int, description string, status string) error {
	// Implementation for updating a task
	for idx, task := range *tl {
		if task.Id == id {
			task.Description = description
			task.Status = status
			task.UpdatedAt = time.Now()
			(*tl)[idx] = task
			err := saveTasks(dataFile, *tl)
			if err != nil {
				fmt.Println("Error saving tasks:", err)
				return err
			}
			return nil
		}
	}
	return errors.New("No such task")
}

// delete
func (tl *TaskList) deleteTask(id int) error {
	// Implementation for deleting a task
	for i, task := range *tl {
		if task.Id == id {
			*tl = append((*tl)[:i], (*tl)[i+1:]...)
			saveTasks(dataFile, *tl)
			return nil
		}
	}
	return errors.New("No such task")
}

func (tl *TaskList) markTask(id int, status string) error {
	// Implementation for marking a task
	for idx, task := range *tl {
		if task.Id == id {
			task.Status = status
			task.UpdatedAt = time.Now()
			(*tl)[idx] = task
			err := saveTasks(dataFile, *tl)
			if err != nil {
				fmt.Println("Error saving tasks:", err)
				return err
			}
			return nil
		}
	}
	return errors.New("No such task")
}

func main() {
	// positional argument parsing
	flag.Parse()
	var operation string = flag.Args()[0]
	var id int
	var description string
	var status string
	var filter string

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
		description = flag.Args()[1]
		err := tempTaskList.addTask(description)
		if err != nil {
			fmt.Println("Error adding task:", err)
			return
		}
		fmt.Printf("Output: Task added successfully (ID: %d)\n", len(tempTaskList))
	case "list":
		// List tasks
		if len(flag.Args()) == 1 {
			filter = "all"
		} else if len(flag.Args()) > 1 {
			filter = flag.Args()[1]
		}
		err := tempTaskList.listTasks(filter)
		if err != nil {
			fmt.Println("Error listing tasks:", err)
			return
		}
	case "update":
		// Update task
		id, _ = strconv.Atoi(flag.Args()[1])
		description = flag.Args()[2]
		err := tempTaskList.updateTask(id, description, status)
		if err != nil {
			fmt.Println("Error updating task:", err)
			return
		}
		fmt.Println("Output: Task updated successfully")
	case "delete":
		// Delete task
		id, _ = strconv.Atoi(flag.Args()[1])
		err := tempTaskList.deleteTask(id)
		if err != nil {
			fmt.Println("Error deleting task:", err)
			return
		}
	case "mark-in-progress":
		// Mark task as in-progress
		id, _ = strconv.Atoi(flag.Args()[1])
		status = "in-progress"
		err := tempTaskList.markTask(id, status)
		if err != nil {
			fmt.Println("Error marking task:", err)
			return
		}
		fmt.Println("Output: Task marked as in-progress successfully")
	case "mark-done":
		// Mark task as done
		id, _ = strconv.Atoi(flag.Args()[1])
		status = "done"
		err := tempTaskList.markTask(id, status)
		if err != nil {
			fmt.Println("Error marking task:", err)
			return
		}
		fmt.Println("Output: Task marked as done successfully")
	default:
		// Invalid operation
		fmt.Println("Invaild operation. Use add, list, update, or delete, mark-in-progress, mark-done.")
	}
}
