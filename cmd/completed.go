package cmd

import (
	"fmt"
	"task/db"

	"github.com/spf13/cobra"
)

var completedCmd = &cobra.Command{
	Use:   "completed",
	Short: "Lists completed tasks from today",
	Run: func(cmd *cobra.Command, args []string) {
		tasks, err := db.CompletedTasks()
		if err != nil {
			fmt.Println("Something went wrong:", err.Error())
			return
		}
		if len(tasks) == 0 {
			fmt.Println("No tasks completed today.")
			return
		}
		fmt.Println("You have finished the following tasks today:")
		for i, task := range tasks {
			fmt.Printf("%d. %s\n", i+1, task.Value)
		}
	},
}

func init() {
	RootCmd.AddCommand(completedCmd)
}
