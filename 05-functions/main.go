package main

import (
	"errors"
	"fmt"
	"os"
)

// NOTE: Functions in Go
// * 1. Functions are first class values,
// which means a function can be assigned to a variable,
// passed as an args or returned from another function.

// Example: add
func add(a, b int) int {
	return a + b
}

// * 2. Multiple return values are built into the language
// Example: divide
func divide(a, b int) (int, error) {
	if b == 0 {
		return 0, errors.New("division by zero")
	}
	return a / b, nil
}

// * 3. Parameters and return types are explicitly typed
// Example: greet
func greet(name string, age int) string {
	return "Hello" + name
}

// * 4. Variadic functions are built in: We can accept an arbitrary number of arguments:
// Example: Sum
func sum(nums ...int) int {
	total := 0
	for _, n := range nums {
		total += n
	}
	return total
}

// * 5. Functions can return functions
// Example: Multiplier
func multiplier(x int) func(int) int {
	return func(y int) int {
		return x * y
	}
}

// * 6. closures captures variables from their surrounding scopes
// Example: counter
func counter() func() int {
	n := 0
	return func() int {
		n++
		return n
	}
}

// * 7. `defer` functions: is tightly integrated with functions
// * defer schedules a function call when the surrounding function returns
// Example: process
func process() {
	defer fmt.Println("defer cleanup: <runs at last>")
	fmt.Println("processing before running defer func")
}

// defer function is heavily used in cleanup:
func cleanup() {
	file, err := os.Open("something.txt")
	if err != nil {
		fmt.Println("error.Open.File<Something.txt> ", err)
	}
	defer file.Close() // cleanup here
}

func main() {
	// assigning function to a variable
	f := add
	fmt.Println(f(1, 2))

	result, err := divide(12, 3)
	fmt.Println(result, err)

	// Multiplier
	double := multiplier(10)
	fmt.Println(double(5))

	// counter
	c := counter()
	fmt.Println(c()) // 1
	fmt.Println(c()) // 2
	fmt.Println(c()) // 3

	process()
	cleanup()
}
