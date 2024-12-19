package cmd

import (
	"fmt"
	"log"
	"strings"

	"github.com/D-CetinEren/backend-projects/go/task-tracker/internal/task"
	"github.com/spf13/cobra"
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
			log.Printf("Error reading tasks: %v", err)
			return
		}
		newTask := task.NewTask(description)
		newTask.ID = len(tasks) + 1
		tasks = append(tasks, newTask)
		if err := task.WriteTasks(tasks); err != nil {
			log.Printf("Error writing tasks: %v", err)
			return
		}
		log.Printf("Task added: %+v", newTask)
		fmt.Printf("Task added successfully (ID: %d)\n", newTask.ID)
	},
}

func init() {
	rootCmd.AddCommand(addCmd)
}
