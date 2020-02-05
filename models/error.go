package models

import "fmt"

type NotFoundError struct {
	error uint32
}

func (e *NotFoundError) Error() string {
	e.error = 400
	return fmt.Sprintf("Item not found")
}

type InternalError struct {
	error uint32
}

func (e *InternalError) Error() string {
	e.error = 500
	return fmt.Sprintf("Something went wrong")
}