package main

import (
	"fmt"
	"gomind/db"
	"gomind/reminder"
	"gomind/task"
	"log"
	"os"
	"sync"
	"time"
)

func main() {
	if _, err := db.Connect(); err != nil {
		log.Fatal("Error connecting to the database: ", err)
		os.Exit(1)
	}

	notifyCh := make(chan string)

	// Wait group to synchronize concurrent operations
	var wg sync.WaitGroup

	// Adding a task concurrently
	wg.Add(1)
	go func() {
		defer wg.Done()
		newTask, err := task.AddTaskv2("Test Task", time.Now().Add(24*time.Hour))
		if err != nil {
			log.Println("Error adding task: ", err)
			return
		}
		fmt.Printf("Task added: ID %d, Description %s, Due %s\n", newTask.ID, newTask.Description, newTask.Due)

		go reminder.StartReminder(newTask.ID, newTask.Description, newTask.Due, notifyCh)
	}()

	// Fetching tasks concurrently
	wg.Add(1)
	go func() {
		defer wg.Done()
		tasks, err := task.GetAllTasks()
		if err != nil {
			log.Fatal("Error getting all tasks: ", err)
		}
		for _, t := range tasks {
			fmt.Printf("Task ID: %d, Description %s, Due %s\n", t.ID, t.Description, t.Due)
		}
	}()

	// Listen for notifications (reminder notifications)
	go func() {
		for msg := range notifyCh {
			fmt.Println("Reminder:", msg)
		}
	}()

	wg.Wait()
}
