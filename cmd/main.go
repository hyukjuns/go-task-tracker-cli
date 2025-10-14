package main

import (
	"encoding/json"
	"flag"
	"fmt"
	"os"
	"time"
)

// 데이터 파일
const dataFile = "data.json"

// Task 구조체 선언 (json 인코딩으로 인해 앞에 대문자)
type Task struct {
	Id          int       `json:"id"`
	Description string    `json:"description"`
	Status      string    `json:"status"`
	CreatedAt   time.Time `json:"created_at"`
	UpdatedAt   time.Time `json:"updated_at"`
	NextId      int       `json:"nextid"` // -> TaskList 단위로 관리 필요
}

type TaskList []Task //작업 버퍼용 배열

// 파일에서 데이터 읽어들임 json file -> TaskList
// 파일이 존재하지 않으면 새로운 빈 배열 반환
func loadTasks(filename string) (TaskList, error) {
	data, err := os.ReadFile(filename)
	if err != nil {
		if os.IsNotExist(err) {
			return TaskList{}, nil // Return empty list if file doesn't exist
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

// TaskList -> json file, 파일 없으면 새로 생성
func saveTasks(filename string, tl TaskList) error {
	data, err := json.MarshalIndent(tl, "", "  ")
	if err != nil {
		return err
	}
	return os.WriteFile(filename, data, 0644)
}

// add/list/update/delete 구현 함수
// add
func (tl *TaskList) addTask(description string) {
	now := time.Now()
	if len(*tl) == 0 {
		// If the list is empty, start IDs from 1
		task := Task{
			Id:          1,
			Description: description,
			Status:      "todo",
			CreatedAt:   now,
			UpdatedAt:   now,
			NextId:      2,
		}
		*tl = append(*tl, task)
		return
	}
	// Add new task with incremented ID
	task := Task{
		Id:          (*tl)[len(*tl)-1].NextId,
		Description: description,
		Status:      "todo",
		CreatedAt:   now,
		UpdatedAt:   now,
		NextId:      (*tl)[len(*tl)-1].NextId + 1,
	}
	*tl = append(*tl, task)
}

// list
func listTasks(filename string) {
	loaded, err := loadTasks(filename)
	if err != nil {
		fmt.Println("Error loading tasks:", err)
		return
	}
	for _, task := range loaded {
		fmt.Printf("ID: %d, Description: %s, Status: %s, CreatedAt: %s, UpdatedAt: %s\n", task.Id, task.Description, task.Status, task.CreatedAt, task.UpdatedAt)
	}
}

// update
// delete

func main() {
	// positional argument 사용
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
		listTasks(dataFile)
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
