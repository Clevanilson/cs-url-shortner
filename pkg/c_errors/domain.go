package c_errors

import "fmt"

type DomainError struct {
	message string
}

func NewDomainError(message string) *DomainError {
	return &DomainError{message}
}

func (e *DomainError) Error() string {
	return fmt.Sprintf("[DomainError]: %v", e.message)
}

func (e *DomainError) Message() string {
	return e.message
}
