package cmd

import (
	"fmt"
	"log"
	"time"

	"github.com/D-CetinEren/backend-projects/go/task-tracker/internal/task"
	"github.com/spf13/cobra"
)

var (
	listStatus string
)

var listCmd = &cobra.Command{
	Use:   "list",
	Short: "List tasks",
	Long:  "List tasks with options to filter by status or creation date range.",
	Run: func(cmd *cobra.Command, args []string) {
		tasks, err := task.ReadTasks()
		if err != nil {
			log.Printf("Error reading tasks: %v", err)
			return
		}

		// Filter by status if specified
		var filteredTasks []task.Task
		if listStatus != "" {
			for _, t := range tasks {
				if t.Status == listStatus {
					filteredTasks = append(filteredTasks, t)
				}
			}
		} else {
			filteredTasks = tasks
		}

		// Display tasks
		if len(filteredTasks) == 0 {
			fmt.Println("No tasks found.")
			return
		}

		fmt.Println("Tasks:")
		for _, t := range filteredTasks {
			fmt.Printf("ID: %d, Description: %s, Status: %s, CreatedAt: %s, UpdatedAt: %s\n",
				t.ID, t.Description, t.Status, t.CreatedAt.Format(time.RFC3339), t.UpdatedAt.Format(time.RFC3339))
		}
	},
}

func init() {
	rootCmd.AddCommand(listCmd)
	listCmd.Flags().StringVar(&listStatus, "status", "", "Filter tasks by status (todo, in-progress, done)")
}
