package cmd

import (
	"bytes"
	"testing"

	"github.com/D-CetinEren/backend-projects/go/task-tracker/internal/task"
	"github.com/spf13/cobra"
)

func MockMarkCommand(cmd *cobra.Command, args []string) (*cobra.Command, *bytes.Buffer) {
	output := new(bytes.Buffer)
	cmd.SetOut(output)
	cmd.SetArgs(args)
	return cmd, output
}

func TestMarkInProgressCommand(t *testing.T) {
	mockTasks := []task.Task{
		{ID: 1, Description: "Task to Mark", Status: "todo"},
	}
	task.WriteTasks(mockTasks)

	cmd, output := MockMarkCommand(markInProgressCmd, []string{"1"})
	err := cmd.Execute()
	if err != nil {
		t.Fatalf("Failed to execute mark-in-progress command: %v", err)
	}

	tasks, _ := task.ReadTasks()
	if tasks[0].Status != "in-progress" {
		t.Errorf("Expected status 'in-progress', but got '%s'", tasks[0].Status)
	}

	expectedOutput := "Task marked as in-progress"
	if !bytes.Contains(output.Bytes(), []byte(expectedOutput)) {
		t.Errorf("Expected output to contain '%s', but got '%s'", expectedOutput, output.String())
	}
}
