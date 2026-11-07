package main

import "fmt"

// NOTE: Type inference
// * In type inference, the compiler figures out the type instead of requiring you to specify it explicitly.

// Example:
func multiply[Num int](a, b Num) Num {
	return a * b
}

// Example: type inference with multiple type parameters
func pair[K any, V any](key K, value V) {
	fmt.Println(key, value)
}

// NOTE: But in some situations we must pass the type explicitly.
// For example:
func create[T any]() T {
	var value T
	return value
}

func main() {

	fmt.Println(multiply(12, 23))
	// NOTE: Here we didn't pass the type when passing arguments to the multiply function.
	// The Go compiler infers it like this:

	fmt.Println(multiply[int](30, 45))

	// BUG: multiply is constrained to int, so this does not compile:
	// fmt.Println(multiply(1.5, 2.5))
	// cannot use 1.5 (untyped float constant) as int value in argument to multiply

	pair("name", "Mohammad")
	pair(1, true)

	val := create[string]()
	fmt.Println(val)

}
