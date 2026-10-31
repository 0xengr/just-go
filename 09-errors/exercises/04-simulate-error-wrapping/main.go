package main

import (
	"database/sql"
	"fmt"
)

// Exercise: Illustrate error wrapping in a real-world example by using
// layers: Repository --> Service --> Handler

// 1. Repository layer
func findUserByID(ID string) error {
	// simulate a database query where the user is not found
	err := sql.ErrNoRows

	if err != nil {
		return fmt.Errorf("find user by ID %s: %w", ID, err)
	}
	return nil
}

// 2. Service layer
func getUser(ID string) error {
	err := findUserByID(ID)

	if err != nil {
		return fmt.Errorf("get user: %w", err)
	}
	return nil
}

// 3. User handler layer
func getUserHandler(ID string) {
	err := getUser(ID)

	if err != nil {
		fmt.Println("internal error:", err)
	}
}
func main() {

	getUserHandler("123")
	// output:
	// internal error: get user: find user by ID 123: sql: no rows in result set
	//
	// Every layer added useful context, but the original database error survived.

}
