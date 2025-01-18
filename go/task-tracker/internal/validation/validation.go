package validation

import (
	"fmt"
	"strings"
)

type ValidationError struct {
	Field   string
	Message string
}

func ValidateTask(description string) []ValidationError {
	var errors []ValidationError

	if strings.TrimSpace(description) == "" {
		errors = append(errors, ValidationError{
			Field:   "description",
			Message: "Description cannot be empty",
		})
	}

	if len(description) > 500 {
		errors = append(errors, ValidationError{
			Field:   "description",
			Message: "Description cannot be longer than 500 characters",
		})
	}

	return errors
}

func ValidateStatus(status string) error {
	validStatuses := map[string]bool{
		"todo":        true,
		"in-progress": true,
		"done":        true,
	}

	if !validStatuses[status] {
		return fmt.Errorf("invalid status: %s", status)
	}

	return nil
}
