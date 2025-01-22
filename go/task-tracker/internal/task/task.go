package task

import "time"

type Task struct {
	ID          int       `json:"id"`
	Description string    `json:"description"`
	Status      string    `json:"status"`
	Priority    string    `json:"priority"` // New field: Priority (e.g., high, medium, low)
	Tags        []string  `json:"tags"`     // New field: Tags for categorization
	CreatedAt   time.Time `json:"createdAt"`
	UpdatedAt   time.Time `json:"updatedAt"`
}

// NewTask creates a new task with default values for priority and tags.
func NewTask(description, priority string, tags []string) Task {
	now := time.Now()
	return Task{
		Description: description,
		Status:      "todo",
		Priority:    priority,
		Tags:        tags,
		CreatedAt:   now,
		UpdatedAt:   now,
	}
}
