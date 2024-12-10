package task

import (
	"fmt"
	"time"
)

type Task struct {
	ID          int
	Description string
	Due         time.Time
	Completed   bool
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

func (tm *TaskManager) ListTasks() []Task {
	return tm.Tasks
}
