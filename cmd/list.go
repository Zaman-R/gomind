package cmd

import (
	"fmt"
	"github.com/spf13/cobra"
)

var listCmd = &cobra.Command{
	Use:   "list",
	Short: "List all tasks",
	Run: func(cmd *cobra.Command, args []string) {
		if taskManager == nil {
			fmt.Println("Error: Task manager is not initialized.")
			return
		}

		tasks := taskManager.ListTasks()

		for _, task := range tasks {
			fmt.Printf("ID: %d, Description: %s, Due: %s, Completed: %t\n",
				task.ID, task.Description, task.Due.Format("2006-01-02 15:04:05"), task.Completed)
		}
	},
}

func init() {
	rootCmd.AddCommand(listCmd)
}
