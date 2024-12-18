package task

import (
	"encoding/json"
	"os"
	"path/filepath"
)

const tasksFileName = "tasks.json"

// GetTasksFilePath returns the path to the tasks.json file in the current directory.
func GetTasksFilePath() string {
	return filepath.Join(".", tasksFileName) // "." specifies the current directory
}

// ReadTasks reads tasks from the JSON file.
func ReadTasks() ([]Task, error) {
	filePath := GetTasksFilePath()
	if _, err := os.Stat(filePath); os.IsNotExist(err) {
		return []Task{}, nil
	}
	data, err := os.ReadFile(filePath)
	if err != nil {
		return nil, err
	}
	var tasks []Task
	if err := json.Unmarshal(data, &tasks); err != nil {
		return nil, err
	}
	return tasks, nil
}

// WriteTasks writes tasks to the JSON file.
func WriteTasks(tasks []Task) error {
	filePath := GetTasksFilePath()
	data, err := json.MarshalIndent(tasks, "", "  ")
	if err != nil {
		return err
	}
	if err := os.WriteFile(filePath, data, 0644); err != nil {
		return err
	}
	return nil
}
