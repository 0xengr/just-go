package main

import "fmt"

// NOTE: panic with recover
// * recover is a built-in function that can stop an active panic from
// * continuing to propagate. It must be called directly inside a deferred
// * function: during unwinding, recover captures the panic value and lets the
// * surrounding function return normally.
// * Called when no panic is in progress, recover returns nil and does nothing.

func riskyOperation() {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered from panic:", r)
		}
	}()

	fmt.Println("Operation started")
	panic("something went wrong")
	// NOTE: Once panic runs, the rest of this function is abandoned. The
	// deferred function above still runs while the stack unwinds.
}

func main() {
	fmt.Println("Before operation")
	riskyOperation()
	fmt.Println("After operation")
}
