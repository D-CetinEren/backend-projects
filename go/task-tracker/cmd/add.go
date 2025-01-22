package cmd

import (
	"fmt"
	"strings"

	"github.com/D-CetinEren/backend-projects/go/task-tracker/internal/task"

	"github.com/google/uuid"
	"github.com/spf13/cobra"
)

var (
	addPriority string   // Priority flag for add command
	addTags     []string // Tags flag for add command
)

var addCmd = &cobra.Command{
	Use:   "add",
	Short: "Add a new task",
	Run: func(cmd *cobra.Command, args []string) {
		if len(args) < 1 {
			fmt.Println("Please provide a task description.")
			return
		}

		description := strings.Join(args, " ")
		tasks, err := task.ReadTasks()
		if err != nil {
			fmt.Printf("Error reading tasks: %v\n", err)
			return
		}

		newTask := task.NewTask(description, addPriority, addTags)
		newID := uuid.New().ID()
		newTask.ID = int(newID)
		tasks = append(tasks, newTask)

		if err := task.WriteTasks(tasks); err != nil {
			fmt.Printf("Error writing tasks: %v\n", err)
			return
		}

		fmt.Printf("Task added successfully (ID: %d)\n", newTask.ID)
	},
}

func init() {
	rootCmd.AddCommand(addCmd)
	addCmd.Flags().StringVar(&addPriority, "priority", "medium", "Set task priority (e.g., high, medium, low)")
	addCmd.Flags().StringSliceVar(&addTags, "tags", []string{}, "Set task tags (comma-separated)")
}
