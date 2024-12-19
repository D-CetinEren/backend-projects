package cmd

import (
	"fmt"
	"log"
	"strconv"
	"time"

	"github.com/D-CetinEren/backend-projects/go/task-tracker/internal/task"
	"github.com/spf13/cobra"
)

var updateCmd = &cobra.Command{
	Use:   "update",
	Short: "Update a task",
	Run: func(cmd *cobra.Command, args []string) {
		if len(args) < 2 {
			fmt.Println("Usage: task-tracker update <id> <new description>")
			return
		}

		id, err := strconv.Atoi(args[0])
		if err != nil {
			fmt.Println("Enter a valid ID")
			return
		}

		newDescription := args[1]

		tasks, err := task.ReadTasks()
		if err != nil {
			fmt.Printf("Error reading tasks: %v\n", err)
			return
		}

		updated := false
		for i := range tasks {
			if tasks[i].ID == id {
				tasks[i].Description = newDescription
				tasks[i].UpdatedAt = time.Now()
				updated = true
				break
			}
		}

		if !updated {
			fmt.Printf("Task with ID %d couldn't find\n", id)
			return
		}

		err = task.WriteTasks(tasks)
		if err != nil {
			fmt.Printf("Error saving tasks: %v\n", err)
			return
		}
		log.Printf("Task with ID %d has been updated", id)
		fmt.Printf("Task with ID %d has been updated\n", id)
	},
}

func init() {
	rootCmd.AddCommand(updateCmd)
}
