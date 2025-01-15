package task

import (
	"errors"
	"os"
	"testing"
	"time"
)

// TestNewTask validates the creation of a new task.
func TestNewTask(t *testing.T) {
	description := "Test Task"
	task := NewTask(description)

	if task.Description != description {
		t.Errorf("Expected description %s, got %s", description, task.Description)
	}
	if task.Status != "todo" {
		t.Errorf("Expected status 'todo', got %s", task.Status)
	}
	if task.CreatedAt.IsZero() || task.UpdatedAt.IsZero() {
		t.Errorf("Timestamps not set correctly.")
	}
}

// TestReadTasks validates the reading of tasks from a JSON file.
func TestReadTasks(t *testing.T) {
	tempFile := "temp_tasks.json"
	defer os.Remove(tempFile)

	tasks := []Task{
		{ID: 1, Description: "Task 1", Status: "todo", CreatedAt: time.Now(), UpdatedAt: time.Now()},
	}
	WriteTasks(tasks)

	readTasks, err := ReadTasks()
	if err != nil {
		t.Fatalf("Failed to read tasks: %v", err)
	}
	if len(readTasks) != len(tasks) {
		t.Errorf("Expected %d tasks, got %d", len(tasks), len(readTasks))
	}
}

// TestWriteTasks ensures tasks are written to the file correctly.
func TestWriteTasks(t *testing.T) {
	tempFile := "temp_tasks.json"
	defer os.Remove(tempFile)

	tasks := []Task{
		{ID: 1, Description: "Task 1", Status: "todo", CreatedAt: time.Now(), UpdatedAt: time.Now()},
	}
	err := WriteTasks(tasks)
	if err != nil {
		t.Fatalf("Failed to write tasks: %v", err)
	}
}

// TestReadTasksWithInvalidFile handles cases where the JSON file is invalid.
func TestReadTasksWithInvalidFile(t *testing.T) {
	tempFile := "temp_tasks.json"
	os.WriteFile(tempFile, []byte("invalid json"), 0644)
	defer os.Remove(tempFile)

	_, err := ReadTasks()
	if !errors.Is(err, errors.New("invalid character")) {
		t.Errorf("Expected JSON parsing error, got %v", err)
	}
}
