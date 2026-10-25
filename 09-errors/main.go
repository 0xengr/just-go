package main

import (
	"errors"
	"fmt"
	"os"
)

// NOTE: Errors in Go
// * Errors in Go are very different from other languages.
// * In Go, errors are values, and you can handle them like any other value.
// * There are no try/catch/finally blocks in Go like there are in JavaScript,
// * Python, Ruby, and other languages.
// * Go says: if there is an error, handle it explicitly; don't panic on errors.
// * Handling errors is simple in Go — boring and beautiful too.
// * One of the best features in Go is that a function returns the error
// * alongside its result, so the caller sees both at once.

// Example:
func calcQuotientAndRemainder(numerator, denominator int) (int, int, error) {
	if denominator == 0 {
		return 0, 0, errors.New("denominator is 0")
	}
	return numerator / denominator, numerator % denominator, nil
}

func main() {

	// NOTE: Unlike languages with exceptions, Go has no special construct
	// * for detecting that an error was returned.
	// * When a function returns an error, an if statement checks it against nil.

	numerator, denominator := 30, 3
	quotient, remainder, err := calcQuotientAndRemainder(numerator, denominator)

	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
	fmt.Printf("quotient: %d\nremainder: %d\n", quotient, remainder)
	// output:
	// quotient: 10
	// remainder: 0

}
