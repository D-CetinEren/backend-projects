package cmd

import (
	"fmt"
	"strconv"
	"time"

	"github.com/D-CetinEren/backend-projects/go/task-tracker/internal/task"
	"github.com/spf13/cobra"
)

var markInProgressCmd = &cobra.Command{
	Use:   "mark-in-progress",
	Short: "Mark a task as in progress",
	Run: func(cmd *cobra.Command, args []string) {
		markTaskStatus(args, "in-progress")
	},
}

var markDoneCmd = &cobra.Command{
	Use:   "mark-done",
	Short: "Mark a task as done",
	Run: func(cmd *cobra.Command, args []string) {
		markTaskStatus(args, "done")
	},
}

func markTaskStatus(args []string, newStatus string) {
	if len(args) < 1 {
		fmt.Println("Usage: task-tracker mark-" + newStatus + " <id>")
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

	updated := false
	for i := range tasks {
		if tasks[i].ID == id {
			tasks[i].Status = newStatus
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

	fmt.Printf("Task with ID %d status has been changed as %s\n", id, newStatus)
}

func init() {
	rootCmd.AddCommand(markInProgressCmd)
	rootCmd.AddCommand(markDoneCmd)
}
