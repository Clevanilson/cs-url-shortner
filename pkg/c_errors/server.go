package c_errors

import "fmt"

type InternalServerError struct {
	message string
}

func NewInternalServerError(message string) *InternalServerError {
	return &InternalServerError{message}
}

func (e *InternalServerError) Error() string {
	return fmt.Sprintf("[InternalServerError]: %v", e.message)
}

func (e *InternalServerError) Message() string {
	return e.message
}
