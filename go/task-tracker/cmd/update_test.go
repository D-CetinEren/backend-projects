package cmd

import (
	"bytes"
	"testing"

	"github.com/D-CetinEren/backend-projects/go/task-tracker/internal/task"
	"github.com/spf13/cobra"
)

func MockUpdateCommand(args []string) (*cobra.Command, *bytes.Buffer) {
	output := new(bytes.Buffer)
	cmd := updateCmd
	cmd.SetOut(output)
	cmd.SetArgs(args)
	return cmd, output
}

func TestUpdateCommandSuccess(t *testing.T) {
	mockTasks := []task.Task{
		{ID: 1, Description: "Old Task", Status: "todo"},
	}
	task.WriteTasks(mockTasks)

	mockArgs := []string{"1", "Updated Task"}
	cmd, output := MockUpdateCommand(mockArgs)

	err := cmd.Execute()
	if err != nil {
		t.Fatalf("Failed to execute update command: %v", err)
	}

	tasks, _ := task.ReadTasks()
	if tasks[0].Description != "Updated Task" {
		t.Errorf("Expected updated description 'Updated Task', but got '%s'", tasks[0].Description)
	}

	expectedOutput := "Task updated successfully"
	if !bytes.Contains(output.Bytes(), []byte(expectedOutput)) {
		t.Errorf("Expected output to contain '%s', but got '%s'", expectedOutput, output.String())
	}
}

func TestUpdateCommandInvalidID(t *testing.T) {
	mockArgs := []string{"999", "Updated Task"}
	cmd, output := MockUpdateCommand(mockArgs)

	err := cmd.Execute()
	if err == nil {
		t.Fatal("Expected an error when updating a non-existent task, but got none")
	}

	expectedOutput := "Task with ID 999 not found."
	if !bytes.Contains(output.Bytes(), []byte(expectedOutput)) {
		t.Errorf("Expected output to contain '%s', but got '%s'", expectedOutput, output.String())
	}
}
