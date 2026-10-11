package main

import (
	"errors"
	"fmt"
)

// NOTE: Named return values in GO.
// The names are predeclared variables, zero-valued like any other, usable
// anywhere in the body.

// Normally we return from a func something like this:
func divAndRemainderDemo1(num, denom int) (int, int, error) {
	if denom == 0 {
		return 0, 0, errors.New("cannot divide by zero")
	}
	return num / denom, num % denom, nil
}

// Now with named return values. `result`, `remainder` and `err` are inside the
// function now, so we can do something like this:
func divAndRemainderWithNamedReturns(num, denom int) (
	result int,
	remainder int,
	err error,
) {
	if denom == 0 {
		return 0, 0, errors.New("cannot divide by zero")
	}

	result = num / denom
	remainder = num % denom

	// WARNING: never write a blank return. This one means "return whatever
	// result, remainder and err hold right now", so the reader has to scan back
	// up to learn what is being returned. And when the input is wrong we return
	// before assigning anything, so the caller gets 0, 0 for the two results.
	// Always write the values out: `return result, remainder, nil`.
	return
}

// NOTE: naming the values doesn't mean you have to return them. The compiler
// copies what an explicit return hands back into the named variables, so these
// two are dead assignments. It compiles, and it confuses everybody.
func divAndRemainderWithNamedReturns2(num, denom int) (
	result int,
	remainder int,
	err error,
) {
	result, remainder = 20, 30

	if denom == 0 {
		return 0, 0, errors.New("cannot divide by zero")
	}

	return num / denom, num % denom, nil
}

// WARNING: you can shadow a named return value. `=` assigns to it, so this one
// works.
func divide(num, denom int) (result int, err error) {
	if denom == 0 {
		return 0, errors.New("cannot divide by zero")
	}

	result = num / denom
	return
}

// One character changed: `:=` declares a NEW result scoped to that if, so the
// named one is never assigned and the blank return hands back 0. No error, no
// panic, just a wrong answer. Be sure you are assigning to the return value and
// not to a shadow of it.
func divideShadowed(num, denom int) (result int, err error) {
	if denom == 0 {
		return 0, errors.New("cannot divide by zero")
	}

	if result := num / denom; result > 0 {
		fmt.Println("     shadowed result =", result)
	}
	return
}

func main() {
	q, r, err := divAndRemainderDemo1(5, 2)
	fmt.Println("divAndRemainderDemo1(5, 2)   ->", q, r, err)

	q, r, err = divAndRemainderWithNamedReturns(5, 2)
	fmt.Println("withNamedReturns(5, 2)       ->", q, r, err)

	q, r, err = divAndRemainderWithNamedReturns2(5, 2)
	fmt.Println("withNamedReturns2(5, 2)      ->", q, r, err)

	result, err := divide(10, 2)
	fmt.Println("divide(10, 2)                ->", result, err)

	result, err = divideShadowed(10, 2)
	fmt.Println("divideShadowed(10, 2)        ->", result, err)
}
