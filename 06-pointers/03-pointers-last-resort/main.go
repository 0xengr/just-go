package main

import "fmt"

// NOTE: The Pointers Last Resort
// * Don't use pointers by default, use them
// * when you have a specific reason to.
//
// NOTE: Why don't use Pointers everywhere ?
// - Nil values: a pointer can be nil, so dereferencing it may panic.

// - Mutation: other code can change the pointed-to value, making behavior harder to reason about.

// PERF: - Memory management: pointers can affect object lifetimes and
// garbage-collection work.
//
// A pointer is not automatically more efficient than a value. Copying a small struct may be simpler and cheaper than introducing pointer indirection.

// NOTE: prefers to use value when possible.
func updateAge(age int) {
	age = 30
}

func updateAgeTwo(age *int) {
	*age = 30
}

func main() {
	age := 25
	updateAge(age)

	fmt.Println(age) // 25

	ageTwo := 25
	updateAgeTwo(&ageTwo)

	fmt.Println(ageTwo) // 30
}
