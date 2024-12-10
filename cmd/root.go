package cmd

import (
	"fmt"
	"github.com/spf13/cobra"
	"gomind/task"
	"os"
)

var rootCmd = &cobra.Command{
	Use:   "gomind",
	Short: "GoMind is a CLI tool for task management and reminders.",
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("Running GoMind CLI")
	},
}

var notifyCh chan string

func Execute(tm *task.TaskManager, ch chan string) {
	taskManager = tm
	notifyCh = ch
	if err := rootCmd.Execute(); err != nil {
		os.Exit(1)
	}
}
