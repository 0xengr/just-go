package main

import (
	"database/sql"
	"errors"
	"fmt"
)

// NOTE: errors.Is
// * errors.Is reports whether an error chain contains a given error value,
// * walking through wrapping until it finds the match. Plain == cannot do
// * that: every fmt.Errorf("...: %w", err) layer returns a new error value,
// * so == misses the original error underneath the wrappers.
//
// Example: three layers of wrapping, matched at the top with errors.Is

// 1. Repository layer
func findUserByID(id int) error {
	return fmt.Errorf("find user: %w", sql.ErrNoRows)
}

// 2. Service layer
func getUser(id int) error {
	err := findUserByID(id)
	if err != nil {
		return fmt.Errorf("get user: %w", err)
	}
	return nil
}

// 3. Handler layer
func getUserHandler(id int) {
	err := getUser(id)

	// NOTE: == fails here. By the time the error reaches the handler, both
	// layers above have wrapped it, so it is a different value from
	// sql.ErrNoRows. This version would fall through and print
	// "internal error: ..." every time.
	// if err == sql.ErrNoRows {
	// 	fmt.Println("user not found")
	// 	return
	// }
	// fmt.Println("internal error:", err)

	if errors.Is(err, sql.ErrNoRows) {
		fmt.Println("user not found")
		return
	}
	fmt.Println("internal error:", err)

	// output: user not found
}

// Another example
var ErrNotFound = errors.New("not found")

func repository() error {
	return ErrNotFound
}

func service() error {
	err := repository()
	if err != nil {
		return fmt.Errorf("service failed: %w", err)
	}

	return nil
}

func main() {

	getUserHandler(1234)

	err := service()

	fmt.Println(err == ErrNotFound)          // false
	fmt.Println(errors.Is(err, ErrNotFound)) // true

}
