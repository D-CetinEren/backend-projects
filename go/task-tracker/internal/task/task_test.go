package task

import (
	"os"
	"testing"
	"time"
)

// TestNewTask validates the creation of a new task, including Priority and Tags fields.
func TestNewTask(t *testing.T) {
	description := "Test Task"
	priority := "high"
	tags := []string{"urgent", "work"}
	task := NewTask(description, priority, tags)

	if task.Description != description {
		t.Errorf("Expected description '%s', got '%s'", description, task.Description)
	}
	if task.Priority != priority {
		t.Errorf("Expected priority '%s', got '%s'", priority, task.Priority)
	}
	if len(task.Tags) != len(tags) {
		t.Errorf("Expected %d tags, got %d", len(tags), len(task.Tags))
	}
	if task.Status != "todo" {
		t.Errorf("Expected status 'todo', got '%s'", task.Status)
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
		{ID: 1, Description: "Task 1", Priority: "medium", Tags: []string{"home"}, Status: "todo", CreatedAt: time.Now(), UpdatedAt: time.Now()},
		{ID: 2, Description: "Task 2", Priority: "high", Tags: []string{"work", "urgent"}, Status: "done", CreatedAt: time.Now(), UpdatedAt: time.Now()},
	}
	err := WriteTasks(tasks)
	if err != nil {
		t.Fatalf("Failed to write tasks: %v", err)
	}

	readTasks, err := ReadTasks()
	if err != nil {
		t.Fatalf("Failed to read tasks: %v", err)
	}

	if len(readTasks) != len(tasks) {
		t.Errorf("Expected %d tasks, got %d", len(tasks), len(readTasks))
	}

	for i, task := range tasks {
		if readTasks[i].Description != task.Description {
			t.Errorf("Expected description '%s', got '%s'", task.Description, readTasks[i].Description)
		}
		if readTasks[i].Priority != task.Priority {
			t.Errorf("Expected priority '%s', got '%s'", task.Priority, readTasks[i].Priority)
		}
		if len(readTasks[i].Tags) != len(task.Tags) {
			t.Errorf("Expected %d tags, got %d", len(task.Tags), len(readTasks[i].Tags))
		}
	}
}

// TestWriteTasks ensures tasks are written to the file correctly, including new fields.
func TestWriteTasks(t *testing.T) {
	tempFile := "temp_tasks.json"
	defer os.Remove(tempFile)

	tasks := []Task{
		{ID: 1, Description: "Task 1", Priority: "medium", Tags: []string{"home"}, Status: "todo", CreatedAt: time.Now(), UpdatedAt: time.Now()},
	}
	err := WriteTasks(tasks)
	if err != nil {
		t.Fatalf("Failed to write tasks: %v", err)
	}

	data, err := os.ReadFile(tempFile)
	if err != nil {
		t.Fatalf("Failed to read written tasks file: %v", err)
	}

	if len(data) == 0 {
		t.Errorf("Expected non-empty file, but got empty")
	}
}

// TestReadTasksWithInvalidFile handles cases where the JSON file is invalid.
func TestReadTasksWithInvalidFile(t *testing.T) {
	tempFile := "temp_tasks.json"
	defer os.Remove(tempFile)

	invalidData := []byte("invalid json data")
	err := os.WriteFile(tempFile, invalidData, 0644)
	if err != nil {
		t.Fatalf("Failed to create invalid JSON file: %v", err)
	}

	_, err = ReadTasks()
	if err == nil {
		t.Errorf("Expected an error when reading invalid JSON, but got none")
	}
}
