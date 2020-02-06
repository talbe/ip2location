package models

import "fmt"

type NotFoundError struct {
}

func (e *NotFoundError) Error() string {
	return fmt.Sprintf("Item not found")
}

type InternalError struct {
}

func (e *InternalError) Error() string {
	return fmt.Sprintf("Something went wrong")
}

type HttpError struct {
	Error uint32
}