package errors

import "fmt"

type TaskError struct {
	Code    string
	Message string
	Err     error
}

func (e *TaskError) Error() string {
	if e.Err != nil {
		return fmt.Sprintf("%s: %s (%v)", e.Code, e.Message, e.Err)
	}
	return fmt.Sprintf("%s: %s", e.Code, e.Message)
}

func NewTaskNotFoundError(id int) *TaskError {
	return &TaskError{
		Code:    "TASK_NOT_FOUND",
		Message: fmt.Sprintf("Task with ID %d not found", id),
	}
}

func NewStorageError(err error) *TaskError {
	return &TaskError{
		Code:    "STORAGE_ERROR",
		Message: "Failed to perform storage operation",
		Err:     err,
	}
}
