package main

import (
	"fmt"
	"gomind/cmd"
	"gomind/storage"
	"gomind/task"
)

var notifyCh chan string

func main() {
	taskManager := task.NewTaskManager()
	fmt.Println("Task Manager Initialized:", taskManager)

	tasks, err := storage.LoadTasks()
	if err != nil {
		fmt.Println("Error loading tasks:", err)
		return
	}
	taskManager.Tasks = tasks

	notifyCh = make(chan string)
	go func() {
		for msg := range notifyCh {
			fmt.Println(msg)
		}
	}()

	cmd.Execute(taskManager, notifyCh)

	err = storage.SaveTasks(taskManager.Tasks)
	if err != nil {
		fmt.Println("Error saving tasks:", err)
	}
}
