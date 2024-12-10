package task

import (
	"fmt"
	"gomind/db"
	"log"
	"time"
)

type Task struct {
	ID          int
	Description string
	Due         time.Time
	Completed   bool
}

func GetAllTasks() ([]Task, error) {
	var tasks []Task
	rows, err := db.Conn.Query("SELECT id, description, due, completed FROM tasks")
	if err != nil {
		return nil, fmt.Errorf("could not fetch tasks: %v", err)
	}
	defer rows.Close()

	for rows.Next() {
		var task Task
		if err := rows.Scan(&task.ID, &task.Description, &task.Due, &task.Completed); err != nil {
			return nil, fmt.Errorf("could not scan task: %v", err)
		}
		tasks = append(tasks, task)
	}

	return tasks, nil
}

type TaskManager struct {
	Tasks  []Task
	NextID int
}

func NewTaskManager() *TaskManager {
	return &TaskManager{Tasks: []Task{}, NextID: 1}
}

func (tm *TaskManager) AddTask(description string, due time.Time) Task {
	fmt.Println("Adding task:", description, due)

	task := Task{
		ID:          tm.NextID,
		Description: description,
		Due:         due,
		Completed:   false,
	}
	tm.NextID++
	tm.Tasks = append(tm.Tasks, task)
	return task
}

func AddTaskv2(description string, due time.Time) (*Task, error) {
	// Prepare SQL statement to insert a new task
	query := `INSERT INTO tasks (description, due, completed) VALUES ($1, $2, $3) RETURNING id, description, due`
	var task Task

	// Execute the query
	err := db.Conn.QueryRow(query, description, due, false).Scan(&task.ID, &task.Description, &task.Due)
	if err != nil {
		log.Println("Error inserting task: ", err)
		return nil, fmt.Errorf("could not insert task: %v", err)
	}

	return &task, nil
}

func (tm *TaskManager) ListTasks() []Task {
	return tm.Tasks
}
