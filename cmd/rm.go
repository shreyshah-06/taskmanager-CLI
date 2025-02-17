package cmd

import (
	"fmt"
	"strconv"
	"task/db"

	"github.com/spf13/cobra"
)

var rmCmd = &cobra.Command{
	Use:   "rm",
	Short: "Deletes a task permanently",
	Run: func(cmd *cobra.Command, args []string) {
		tasks, err := db.AllTasks()
		if err != nil {
			fmt.Println("Something went wrong:", err.Error())
			return
		}

		if len(tasks) == 0 {
			fmt.Println("No tasks to delete.")
			return
		}

		for _, arg := range args {
			index, err := strconv.Atoi(arg)
			if err != nil || index <= 0 || index > len(tasks) {
				fmt.Println("Invalid task number:", arg)
				continue
			}

			taskKey := tasks[index-1].Key

			err = db.DeleteTask(taskKey)
			if err != nil {
				fmt.Printf("Failed to delete task %d: %s\n", index, err.Error())
			} else {
				fmt.Printf("You have deleted task %d.\n", index)
			}
		}
	},
}

func init() {
	RootCmd.AddCommand(rmCmd)
}
