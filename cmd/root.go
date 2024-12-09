package cmd

import (
	"fmt"
	"github.com/spf13/cobra"
)

var rootCmd = &cobra.Command{
	Use:   "gomind",
	Short: "GoMind is a CLI tool for task management and reminders.",
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("Running GoMind CLI")
	},
}

func Execute() {
	cobra.CheckErr(rootCmd.Execute())
}
