package task

import "time"

// Task represents a single task in the task tracker.
type Task struct {
	ID          int       `json:"id"`          // Unique identifier for the task
	Description string    `json:"description"` // Brief description of the task
	Status      string    `json:"status"`      // Status: "todo", "in-progress", or "done"
	CreatedAt   time.Time `json:"createdAt"`   // Timestamp when the task was created
	UpdatedAt   time.Time `json:"updatedAt"`   // Timestamp when the task was last updated
}

// NewTask initializes a new task with the given description.
func NewTask(description string) Task {
	now := time.Now()
	return Task{
		Description: description,
		Status:      "todo",
		CreatedAt:   now,
		UpdatedAt:   now,
	}
}
