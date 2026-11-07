package main

import "fmt"

// NOTE: Exercise: remove duplicate values.
// * Unique needs T comparable because the seen map uses the value as a key.
// * The result keeps the first occurrence of each value and its original order.

func Unique[T comparable](values []T) []T {
	seen := make(map[T]struct{}, len(values))
	result := make([]T, 0, len(values))

	for _, value := range values {
		if _, ok := seen[value]; ok {
			continue
		}
		seen[value] = struct{}{}
		result = append(result, value)
	}
	return result
}

func main() {
	fmt.Println(Unique([]int{1, 2, 1, 3, 2}))
	fmt.Println(Unique([]string{"go", "is", "go", "fun"}))
	fmt.Println(Unique([]string{}))
}
