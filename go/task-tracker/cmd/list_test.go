package cmd

import (
	"bytes"
	"testing"

	"github.com/D-CetinEren/backend-projects/go/task-tracker/internal/task"
	"github.com/spf13/cobra"
)

func MockListCommand(args []string) (*cobra.Command, *bytes.Buffer) {
	output := new(bytes.Buffer)
	cmd := listCmd
	cmd.SetOut(output)
	cmd.SetArgs(args)
	return cmd, output
}

func TestListCommandAllTasks(t *testing.T) {
	mockTasks := []task.Task{
		{ID: 1, Description: "Task 1", Status: "todo"},
		{ID: 2, Description: "Task 2", Status: "done"},
	}
	task.WriteTasks(mockTasks)

	cmd, output := MockListCommand([]string{})
	err := cmd.Execute()
	if err != nil {
		t.Fatalf("Failed to execute list command: %v", err)
	}

	for _, task := range mockTasks {
		if !bytes.Contains(output.Bytes(), []byte(task.Description)) {
			t.Errorf("Expected output to contain task '%s'", task.Description)
		}
	}
}
