package main

import "fmt"

// NOTE: Generic functions with function parameters
// * A generic function can accept another function. All and Any work on a
// * slice of any element type; the caller supplies the check.
// * All reports whether every value passes the check. Any reports whether at
// * least one value passes. On an empty slice, All is true and Any is false.

func All[T any](values []T, check func(T) bool) bool {
	for _, value := range values {
		if !check(value) {
			return false
		}
	}
	return true
}

func Any[T any](values []T, check func(T) bool) bool {
	for _, value := range values {
		if check(value) {
			return true
		}
	}
	return false
}

func main() {
	numbers := []int{2, 4, 6}

	fmt.Println(All(numbers, func(n int) bool {
		return n%2 == 0
	}))
	fmt.Println(Any(numbers, func(n int) bool {
		return n > 5
	}))
	fmt.Println(Any(numbers, func(n int) bool {
		return n > 10
	}))

	words := []string{"go", "generics"}

	fmt.Println(All(words, func(word string) bool {
		return len(word) > 1
	}))
	fmt.Println(Any(words, func(word string) bool {
		return word == "go"
	}))

	// NOTE: These two follow the usual rules for empty slices.
	fmt.Println(All([]int{}, func(n int) bool { return n > 0 }))
	fmt.Println(Any([]int{}, func(n int) bool { return n > 0 }))
}
