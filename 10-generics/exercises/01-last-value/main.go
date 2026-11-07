package main

import "fmt"

// NOTE: Exercise: return the last value of a slice.
// * Last works for any element type. The bool result says whether a value was
// * actually found, because the zero value of T could also be a real value.
// * The zero value is written as var zero T.

func Last[T any](values []T) (T, bool) {
	if len(values) == 0 {
		var zero T
		return zero, false
	}
	return values[len(values)-1], true
}

func main() {
	fmt.Println(Last([]int{1, 2, 3}))
	fmt.Println(Last([]string{"go", "generics"}))
	fmt.Println(Last([]int{}))
	fmt.Println(Last([]bool{false, true}))
}
