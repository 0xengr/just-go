package main

import (
	"errors"
	"fmt"
)

// NOTE: In the standard library, the errors package has a function we can use
// * to create errors.

// NOTE: From the standard library source:
// https://golang.org/src/errors/errors.go
// New returns an error that formats as the given text.
// func New(text string) error {
// 	return &errorString{text}
// }

// NOTE: Go's HTTP standard library uses errors.New to create error values.
// * Here are a few real ones from net/http:
var (
	ErrBodyNotAllowed  = errors.New("http: request method or response status code does not allow body")
	ErrHijacked        = errors.New("http: connection has been hijacked")
	ErrContentLength   = errors.New("http: wrote more than the declared Content-Length")
	ErrWriteAfterFlush = errors.New("unused")
)

func main() {

	// Example: creating custom error values
	errBadData := errors.New("some bad data")
	fmt.Printf("errBadData type: %T\n", errBadData) // *errors.errorString
}

// NOTE: Visit ./exercises/01-calculate-pay for a real use case of error values.
