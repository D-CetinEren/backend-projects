package cmd

import (
	"fmt"
	"log"
	"time"

	"github.com/D-CetinEren/backend-projects/go/task-tracker/internal/task"
	"github.com/spf13/cobra"
)

var (
	listStatus    string
	startDateFlag string
	endDateFlag   string
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

		// Filter by date range if specified
		if startDateFlag != "" || endDateFlag != "" {
			startDate, endDate, dateErr := parseDateRange(startDateFlag, endDateFlag)
			if dateErr != nil {
				log.Printf("Invalid date range: %v", dateErr)
				fmt.Println("Please provide valid date formats: YYYY-MM-DD")
				return
			}

			filteredTasks = filterTasksByDate(filteredTasks, startDate, endDate)
		}

		// Display tasks
		if len(filteredTasks) == 0 {
			fmt.Println("No tasks found.")
			return
		}

		log.Printf("Tasks have been listed")
		fmt.Println("Tasks:")
		for _, t := range filteredTasks {
			fmt.Printf("ID: %d, Description: %s, Status: %s, CreatedAt: %s, UpdatedAt: %s\n",
				t.ID, t.Description, t.Status, t.CreatedAt.Format(time.RFC3339), t.UpdatedAt.Format(time.RFC3339))
		}
	},
}

// Parse date range flags
func parseDateRange(startDateStr, endDateStr string) (time.Time, time.Time, error) {
	var startDate, endDate time.Time
	var err error

	if startDateStr != "" {
		startDate, err = time.Parse("2006-01-02", startDateStr)
		if err != nil {
			return time.Time{}, time.Time{}, fmt.Errorf("invalid start date format")
		}
	}

	if endDateStr != "" {
		endDate, err = time.Parse("2006-01-02", endDateStr)
		if err != nil {
			return time.Time{}, time.Time{}, fmt.Errorf("invalid end date format")
		}
	}

	return startDate, endDate, nil
}

// Filter tasks by date range
func filterTasksByDate(tasks []task.Task, startDate, endDate time.Time) []task.Task {
	var filtered []task.Task

	for _, t := range tasks {
		if (startDate.IsZero() || t.CreatedAt.After(startDate) || t.CreatedAt.Equal(startDate)) &&
			(endDate.IsZero() || t.CreatedAt.Before(endDate) || t.CreatedAt.Equal(endDate)) {
			filtered = append(filtered, t)
		}
	}

	return filtered
}

func init() {
	rootCmd.AddCommand(listCmd)

	// Add flags for filtering
	listCmd.Flags().StringVar(&listStatus, "status", "", "Filter tasks by status (todo, in-progress, done)")
	listCmd.Flags().StringVar(&startDateFlag, "start-date", "", "Filter tasks created after this date (YYYY-MM-DD)")
	listCmd.Flags().StringVar(&endDateFlag, "end-date", "", "Filter tasks created before this date (YYYY-MM-DD)")
}
