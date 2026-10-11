package main

import "fmt"

// NOTE: Pointers Indicate Mutable Params
// * As we know, GO is pass by value language.
// * So, when we pass a variable to a function,
// * GO passes a copy of its value.

// NOTE: If we want to modify the caller's variable
// * then we pass pointer to that variable.

// NOTE: Example: Passing without pointers
func updateOne(num int) {
	num = 100
}

// NOTE: Example: passing with pointers
func updateTwo(num *int) {
	*num = 200
}

// WARNING: We cannot replace the caller's pointer
// Example: failedUpdate
func failedUpdate(g *int) {
	x := 10
	g = &x
}

func main() {

	n1 := 10
	updateOne(n1)
	fmt.Println("n1: ", n1) // 10

	n2 := 20
	updateTwo(&n2)
	fmt.Println("n2: ", n2) // 200

	// NOTE: Here's
	// - `&num` -> gets the address of num
	// - `num *int` -> declares num as a pointer to an integer
	// - `*num = 200` -> dereferences the pointer and changes
	// the value at that address

	var f *int
	failedUpdate(f)
	fmt.Println("f: ", f) // f:  <nil>

	// WARNING: * This failed because `g` receives a copy of the pointer
	// * Reassigning `g` doesn't change `f` in the caller.

}
