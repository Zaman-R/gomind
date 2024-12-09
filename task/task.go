package task

import "time"

type Tasks struct {
	ID          int
	Description string
	Due         time.Time
	Completed   bool
}

type TaskManager struct {
	Tasks  []Tasks
	NextID int
}

func NewTaskManager() *TaskManager {
	return &TaskManager{
		Tasks:  []Tasks{},
		NextID: 1,
	}
}

func (tm *TaskManager) AddTask(description string, due time.Time) Task {
	task := Tasks{
		ID:          tm.NextID,
		Description: description,
		Due:         due,
		Completed:   false,
	}
	tm.Tasks = append(tm.Tasks, task)
	tm.NextID++
	return task
}

func (tm *TaskManager) ListTasks() []Tasks {
	return tm.Tasks
}
