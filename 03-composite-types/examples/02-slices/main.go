// NOTE: len is what you may index, cap is the room left in the backing array.
// Only make sets the capacity, and only append grows it.
package main

import "fmt"

func main() {
	var nilScores []int // nil slice: len 0, cap 0, prints []
	fmt.Printf("nil:     len=%d cap=%d %v\n", len(nilScores), cap(nilScores), nilScores)

	empty := []int{} // non-nil but still empty
	fmt.Printf("empty:   len=%d cap=%d %v\n", len(empty), cap(empty), empty)

	literal := []int{3, 1, 2} // len == cap == 3
	fmt.Printf("literal: len=%d cap=%d %v\n", len(literal), cap(literal), literal)

	made := make([]int, 2, 4) // len 2 zeroed, cap 4
	made[1] = 5
	fmt.Printf("make:    len=%d cap=%d %v\n", len(made), cap(made), made)

	// PERF: watch cap double as append runs out of room.
	grown := make([]int, 0, 2)
	for i := range 6 {
		grown = append(grown, i)
		fmt.Printf("append %d -> len=%d cap=%d\n", i, len(grown), cap(grown))
	}
}
