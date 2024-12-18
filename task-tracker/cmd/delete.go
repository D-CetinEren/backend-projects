package cmd

import (
	"fmt"
	"strconv"

	"github.com/D-CetinEren/backend-projects/go/task-tracker/internal/task"
	"github.com/spf13/cobra"
)

var deleteCmd = &cobra.Command{
	Use:   "delete",
	Short: "Delete a task",
	Run: func(cmd *cobra.Command, args []string) {
		if len(args) < 1 {
			fmt.Println("Usage: task-tracker delete <id>")
			return
		}

		id, err := strconv.Atoi(args[0])
		if err != nil {
			fmt.Println("Enter a valid ID")
			return
		}

		tasks, err := task.ReadTasks()
		if err != nil {
			fmt.Printf("Error reading tasks: %v\n", err)
			return
		}

		for i := range tasks {
			if tasks[i].ID == id {
				tasks = append(tasks[:i], tasks[i+1:]...)

				err = task.WriteTasks(tasks)
				if err != nil {
					fmt.Printf("Error saving tasks: %v\n", err)
					return
				}

				fmt.Printf("Task with ID %d has deleted\n", id)
				return
			}
		}

		fmt.Printf("Task with ID %d couldn't find\n", id)
	},
}

func init() {
	rootCmd.AddCommand(deleteCmd)
}
