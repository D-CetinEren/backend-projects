package cmd

import (
	"bytes"
	"testing"

	"github.com/D-CetinEren/backend-projects/go/task-tracker/internal/task"
	"github.com/spf13/cobra"
)

func MockDeleteCommand(args []string) (*cobra.Command, *bytes.Buffer) {
	output := new(bytes.Buffer)
	cmd := deleteCmd
	cmd.SetOut(output)
	cmd.SetArgs(args)
	return cmd, output
}

func TestDeleteCommandSuccess(t *testing.T) {
	mockTasks := []task.Task{
		{ID: 1, Description: "Task to Delete", Status: "todo"},
	}
	task.WriteTasks(mockTasks)

	mockArgs := []string{"1"}
	cmd, output := MockDeleteCommand(mockArgs)

	err := cmd.Execute()
	if err != nil {
		t.Fatalf("Failed to execute delete command: %v", err)
	}

	tasks, _ := task.ReadTasks()
	if len(tasks) != 0 {
		t.Errorf("Expected 0 tasks, but got %d", len(tasks))
	}

	expectedOutput := "Task deleted successfully"
	if !bytes.Contains(output.Bytes(), []byte(expectedOutput)) {
		t.Errorf("Expected output to contain '%s', but got '%s'", expectedOutput, output.String())
	}
}

func TestDeleteCommandInvalidID(t *testing.T) {
	mockArgs := []string{"999"}
	cmd, output := MockDeleteCommand(mockArgs)

	err := cmd.Execute()
	if err == nil {
		t.Fatal("Expected an error when deleting a non-existent task, but got none")
	}

	expectedOutput := "Task with ID 999 not found."
	if !bytes.Contains(output.Bytes(), []byte(expectedOutput)) {
		t.Errorf("Expected output to contain '%s', but got '%s'", expectedOutput, output.String())
	}
}
