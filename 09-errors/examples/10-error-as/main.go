package main

import (
	"errors"
	"fmt"
)

// NOTE: errors.As
// * errors.As finds the first error in the chain whose concrete type matches
// * the target and stores it there, so the caller can read its fields.

// Example: extracting a custom db error

type DatabaseError struct {
	Operation string
	Message   string
}

func (e *DatabaseError) Error() string {
	return e.Operation + ": " + e.Message
}

func queryDatabase() error {
	dbErr := &DatabaseError{
		Operation: "SELECT",
		Message:   "connection timeout",
	}
	return fmt.Errorf("repository failed: %w", dbErr)
}

func main() {

	// Extract it
	err := queryDatabase()

	var dbErr *DatabaseError

	if errors.As(err, &dbErr) {
		fmt.Println("Operation:", dbErr.Operation)
		fmt.Println("Message:", dbErr.Message)
	}

}

// NOTE: Notice the difference:
// Function	Purpose
// errors.Is()	Checks whether an error matches a specific value
// errors.As()	Finds an error of a particular type and extracts it

// A useful rule:
// - Use errors.Is() for sentinel errors, such as sql.ErrNoRows.
// - Use errors.As() for structured custom errors, such as *DatabaseError.
