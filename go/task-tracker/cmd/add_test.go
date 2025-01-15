package cmd

import (
	"bytes"
	"testing"

	"github.com/D-CetinEren/backend-projects/go/task-tracker/internal/task"
	"github.com/spf13/cobra"
)

func MockAddCommand(args []string) (*cobra.Command, *bytes.Buffer) {
	output := new(bytes.Buffer)
	cmd := addCmd
	cmd.SetOut(output)
	cmd.SetArgs(args)
	return cmd, output
}

func TestAddCommandSuccess(t *testing.T) {
	mockArgs := []string{"New Task"}
	cmd, output := MockAddCommand(mockArgs)

	err := cmd.Execute()
	if err != nil {
		t.Fatalf("Failed to execute add command: %v", err)
	}

	tasks, _ := task.ReadTasks()
	if len(tasks) != 1 {
		t.Fatalf("Expected 1 task, but got %d", len(tasks))
	}
	if tasks[0].Description != mockArgs[0] {
		t.Errorf("Expected task description '%s', but got '%s'", mockArgs[0], tasks[0].Description)
	}

	expectedOutput := "Task added successfully"
	if !bytes.Contains(output.Bytes(), []byte(expectedOutput)) {
		t.Errorf("Expected output to contain '%s', but got '%s'", expectedOutput, output.String())
	}
}

func TestAddCommandNoDescription(t *testing.T) {
	mockArgs := []string{}
	cmd, output := MockAddCommand(mockArgs)

	err := cmd.Execute()
	if err == nil {
		t.Fatal("Expected an error when adding a task without a description, but got none")
	}

	expectedOutput := "Please provide a task description."
	if !bytes.Contains(output.Bytes(), []byte(expectedOutput)) {
		t.Errorf("Expected output to contain '%s', but got '%s'", expectedOutput, output.String())
	}
}
