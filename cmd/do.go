package cmd

import (
	"fmt"
	"strconv"
	"task/db"

	"github.com/spf13/cobra"
)

var doCmd = &cobra.Command{
	Use:   "do",
	Short: "Marks a task as completed",
	Run: func(cmd *cobra.Command, args []string) {
		for _, arg := range args {
			id, err := strconv.Atoi(arg)
			if err != nil {
				fmt.Println("Invalid task ID:", arg)
				continue
			}
			err = db.CompleteTask(id)
			if err != nil {
				fmt.Printf("Failed to complete task %d: %s\n", id, err.Error())
			} else {
				fmt.Printf("Marked task %d as completed.\n", id)
			}
		}
	},
}

func init() {
	RootCmd.AddCommand(doCmd)
}
