package cmd

import (
	"fmt"
	"github.com/spf13/cobra"
	"gomind/reminder"
	"gomind/task"
	"time"
)

var taskManager *task.TaskManager
var description string
var due string

var addCmd = &cobra.Command{
	Use:   "add",
	Short: "Add a task",
	Run: func(cmd *cobra.Command, args []string) {
		if taskManager == nil {
			fmt.Println("Error: Task manager is not initialized.")
			return
		}

		dueTime, err := time.Parse("2006-01-02 15:04:05", due)
		if err != nil {
			fmt.Println("Error parsing due date:", err)
			return
		}

		task := taskManager.AddTask(description, dueTime)
		fmt.Printf("Task added: ID=%d, Description=%s, Due=%s\n", task.ID, task.Description, task.Due)

		go reminder.StartReminder(task.ID, task.Description, task.Due, notifyCh)
	},
}

func init() {
	addCmd.Flags().StringVarP(&description, "description", "d", "", "Task description")
	addCmd.Flags().StringVarP(&due, "due", "u", "", "Due time (format: 2006-01-02 15:04:05)")
	addCmd.MarkFlagRequired("description")
	addCmd.MarkFlagRequired("due")
	rootCmd.AddCommand(addCmd)
}
