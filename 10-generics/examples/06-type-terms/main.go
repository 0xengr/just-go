package main

import "fmt"

// NOTE: Type terms and the ~ prefix
// * A constraint can list concrete types joined with |. The operators allowed
// * inside the function are the ones valid for every listed type, so + works
// * here but % would restrict the list to integers.
// * The ~ prefix also matches a named type whose underlying type is listed.
// * That is what lets Dollars use the same Sum function as int64.

type Number interface {
	~int | ~int64 | ~float64
}

func Sum[T Number](values []T) T {
	var total T
	for _, value := range values {
		total += value
	}
	return total
}

type Dollars int64

func main() {
	fmt.Println(Sum([]int{1, 2, 3}))
	fmt.Println(Sum([]int64{10, 20}))
	fmt.Println(Sum([]float64{1.5, 2.5}))
	fmt.Println(Sum([]Dollars{100, 250}))

	// BUG: a string is not in Number, so this does not compile:
	// fmt.Println(Sum([]string{"a", "b"}))
	// string does not satisfy Number (string missing in ~int | ~int64 | ~float64)
	//
	// BUG: the type set also limits the type of literals. A constant must fit
	// every type in the set. Here Number includes only int, int64, and float64,
	// so a literal that fits all three is fine; add ~int8 and the large ones
	// stop compiling.
}
