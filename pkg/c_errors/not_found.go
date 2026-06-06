package c_errors

import "fmt"

type NotFoundError struct {
	message string
}

func NewNotFoundError(message string) *NotFoundError {
	return &NotFoundError{message}
}

func (e *NotFoundError) Error() string {
	return fmt.Sprintf("[NotFound]: %v", e.message)
}

func (e *NotFoundError) Message() string {
	return e.message
}
