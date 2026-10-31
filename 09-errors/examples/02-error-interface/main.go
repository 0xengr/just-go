package main

import "fmt"

// NOTE: An error in Go is anything that implements the error interface.
// * error itself is an interface, declared once in the builtin package:
// *
// *     type error interface {
// *         Error() string
// *     }
// *
// * A type satisfies it implicitly; there is no "implements" keyword.
// * If a type has an Error() string method, it is an error.

// NOTE: Any type with an Error method that returns a string satisfies
// * the error interface. A method with a pointer receiver belongs to the
// * pointer's method set only, so *NetworkError is an error but the plain
// * NetworkError value is not. That is why main passes &networkError.

type NetworkError struct{}

func (ne *NetworkError) Error() string {
	return "network unavailable"
}

type FileError struct{}

func (fe *FileError) Error() string {
	return "failed opening file"
}

type ValidationError struct{}

func (ve *ValidationError) Error() string {
	return "validation error"
}

// NOTE: Compile-time interface checks. Assigning each pointer to the blank
// * identifier makes the compiler verify that the type really implements
// * error; a broken Error signature then fails the build instead of waiting
// * for the first call.
var (
	_ error = (*NetworkError)(nil)
	_ error = (*FileError)(nil)
	_ error = (*ValidationError)(nil)
)

// NOTE: printError accepts the interface, not the concrete types, so it can
// * take any of them and cannot see their struct fields.
func printError(err error) {
	// fmt calls Error() for us, so Println(err) prints the same text.
	fmt.Println(err)
}

func main() {
	networkError := NetworkError{}
	fileError := FileError{}
	validationError := ValidationError{}

	printError(&networkError)
	printError(&fileError)
	printError(&validationError)
}
