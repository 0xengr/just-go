package main

import "fmt"

// NOTE: Generic types
// * A type parameter works on a struct too, not just on a function. Box[T]
// * holds one value of whatever type the caller chooses, and every method keeps
// * that same type.
// * The zero value of a generic type is still useful: an empty Box is a Box
// * that has not been filled yet.

type Box[T any] struct {
	value  T
	filled bool
}

func (b *Box[T]) Put(value T) {
	b.value = value
	b.filled = true
}

func (b Box[T]) Get() (T, bool) {
	return b.value, b.filled
}

func (b Box[T]) IsEmpty() bool {
	return !b.filled
}

func main() {
	var numberBox Box[int]
	numberBox.Put(42)
	value, ok := numberBox.Get()
	fmt.Println(value, ok)
	fmt.Println(numberBox.IsEmpty())

	var wordBox Box[string]
	fmt.Println(wordBox.IsEmpty())
	wordBox.Put("hello")
	fmt.Println(wordBox.Get())

	// BUG: uncommenting the next line does not compile:
	// numberBox.Put("not a number")
	// cannot use "not a number" (untyped string constant) as int value
	// in argument to numberBox.Put
	//
	// The compiler knows numberBox is a Box[int], so a string is rejected at
	// build time instead of failing later at run time.
}
