package cmd

import (
	"fmt"
	"strings"

	"github.com/D-CetinEren/backend-projects/go/task-tracker/internal/task"

	"github.com/spf13/cobra"
)

var (
	listPriority string   // Priority filter for list command
	listTags     []string // Tags filter for list command
)

var listCmd = &cobra.Command{
	Use:   "list",
	Short: "List tasks",
	Run: func(cmd *cobra.Command, args []string) {
		tasks, err := task.ReadTasks()
		if err != nil {
			fmt.Printf("Error reading tasks: %v\n", err)
			return
		}

		var filteredTasks []task.Task
		for _, t := range tasks {
			// Filter by priority
			if listPriority != "" && t.Priority != listPriority {
				continue
			}

			// Filter by tags
			if len(listTags) > 0 && !hasMatchingTags(t.Tags, listTags) {
				continue
			}

			filteredTasks = append(filteredTasks, t)
		}

		if len(filteredTasks) == 0 {
			fmt.Println("No tasks found.")
			return
		}

		for _, t := range filteredTasks {
			fmt.Printf("ID: %d, Description: %s, Priority: %s, Tags: %v, Status: %s\n",
				t.ID, t.Description, t.Priority, t.Tags, t.Status)
		}
	},
}

func hasMatchingTags(taskTags, filterTags []string) bool {
	for _, filterTag := range filterTags {
		for _, taskTag := range taskTags {
			if strings.EqualFold(filterTag, taskTag) {
				return true
			}
		}
	}
	return false
}

func init() {
	rootCmd.AddCommand(listCmd)
	listCmd.Flags().StringVar(&listPriority, "priority", "", "Filter tasks by priority (e.g., high, medium, low)")
	listCmd.Flags().StringSliceVar(&listTags, "tags", []string{}, "Filter tasks by tags (comma-separated)")
}
