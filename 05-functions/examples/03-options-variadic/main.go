package main

import "fmt"

// NOTE: Named and optional parameters
// Go doesn't support named or optional parameters.
// For example, in TypeScript we can simulate named parameters with an object
// and declare optional parameters directly:
/*
 function greet({ name, age }: { name: string; age?: number }) {}
 greet({ name: "John" });
*/

// In Go, we simulate named and optional parameters using a struct.
// Missing fields get the zero value, which acts like "optional".

type greetOptions struct {
	Name    string
	Age     int
	Address string
}

func greetUser(opts greetOptions) {
	fmt.Printf("Hello %s, I know your age is %d\n", opts.Name, opts.Age)
}

func main() {
	greetUser(greetOptions{Name: "Mohammad", Age: 12})

	label, total := addNumbers("addNumbers", 1, 2, 3, 4, 5, 5, 6, 6, 7)
	fmt.Println(label, total)

	// Expand a slice into separate arguments with `...`.
	numbers := []int{12, 12, 13, 14, 15, 1, 15, 22, 4}
	label2, total2 := addNumbers("addNumbers2: ", numbers...)
	fmt.Println(label2, total2)
}

// NOTE: Variadic functions
// In practice, not having named and optional parameters isn't a limitation.
// A function shouldn't take too many parameters anyway.
// But when it needs to accept many values of the same type, use a variadic function:

// A variadic function accepts any number of arguments. We declare it with
// `...` before the type, and inside the function it becomes a slice.
// The variadic parameter must be the last parameter in the list.

// The opposite, `slice...` at the call site, expands a slice into separate
// arguments. See example above.

func addNumbers(label string, nums ...int) (string, int) {
	total := 0
	for _, num := range nums {
		total += num
	}
	return label, total
}
