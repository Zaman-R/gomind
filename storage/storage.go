package storage

import (
	"encoding/json"
	"fmt"
	"gomind/task"
	"io"
	"log"
	"os"
	"time"
)

const layout = "2006-01-02 15:04:05 -0700 MST"

func LoadTasks() ([]task.Task, error) {
	file, err := os.Open("tasks.json")
	if err != nil {
		fmt.Println("Error opening file:", err)
		return nil, err
	}
	defer file.Close()

	var tasks []task.Task
	byteValue, err := io.ReadAll(file)
	if err != nil {
		log.Fatal(err)
	}
	err = json.Unmarshal(byteValue, &tasks)
	if err != nil {
		log.Fatal(err)
	}

	for i, t := range tasks {
		parsedTime, err := time.Parse(layout, t.Due.Format(layout))
		if err != nil {
			fmt.Printf("Error parsing task due date: %v\n", err)
		} else {
			tasks[i].Due = parsedTime
		}
	}

	return tasks, nil
}

func SaveTasks(tasks []task.Task) error {
	file, err := os.Create("tasks.json")
	if err != nil {
		fmt.Println("Error creating file:", err)
		return err
	}
	defer file.Close()

	// Marshal tasks into JSON
	data, err := json.MarshalIndent(tasks, "", "  ")
	if err != nil {
		fmt.Println("Error marshaling tasks to JSON:", err)
		return err
	}

	// Write JSON data to the file
	_, err = file.Write(data)
	if err != nil {
		fmt.Println("Error writing tasks to file:", err)
		return err
	}

	fmt.Println("Tasks saved successfully.")
	return nil
}
