package main

import (
	"errors"
	"fmt"
)

// NOTE: Error wrapping
// * Error wrapping preserves the original error while adding context to it.
// * The mental model is an error inside another error: each layer wraps the
// * one below it, so the chain records where the failure happened and what
// * caused it, making errors easier to understand and handle.

// Example: Error wrapping

func getUsers() error {
	err := errors.New("connection refused")

	return fmt.Errorf("failed to get users: %w", err)
}

func main() {

	err := getUsers()
	fmt.Println(err)
	// output: failed to get users: connection refused
}
