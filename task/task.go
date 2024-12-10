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
	rows, err := db.DB.Query("SELECT id, description, due, completed FROM tasks")
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var tasks []Task
	for rows.Next() {
		var task Task
		if err := rows.Scan(&task.ID, &task.Description, &task.Due, &task.Completed); err != nil {
			log.Fatal(err)
		}
	}
	if err := rows.Err(); err != nil {
		log.Fatal(err)
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

func AddTaskv2(description string, due time.Time) (Task, error) {
	var task Task
	err := db.DB.QueryRow(
		"INSERT INTO tasks (description, due) VALUES ($1, $2) RETURNING id",
		description, due).Scan(&task.ID)
	if err != nil {
		return task, err
	}

	task.Description = description
	task.Due = due
	task.Completed = false
	return task, nil
}

func (tm *TaskManager) ListTasks() []Task {
	return tm.Tasks
}
