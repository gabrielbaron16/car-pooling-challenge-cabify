package utils

import "errors"

var (
	// ErrNotFound is returned when the car is not found
	ErrNotFound = errors.New("not found")

	// ErrDuplicatedID is returned when the ID is duplicated
	ErrDuplicatedID = errors.New("duplicated ID")
)
