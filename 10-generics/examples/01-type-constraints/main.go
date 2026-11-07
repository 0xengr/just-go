package main

import "fmt"

// NOTE: A type parameter allows us to write a function that works with
// different types without duplicating the function for each type.
// A constraint defines which types are allowed for a type parameter.

// Example: an ordinary function fixed to one type
func print(value int) {
	fmt.Println(value)
}

// NOTE: A type constraint looks like this. This version accepts one value of
// any one of the listed types:
// func printValue[T int | string | bool | float32 | float64](value T) {
// 	fmt.Println(value)
// }

// A variadic any parameter is the alternative. It accepts several arguments,
// but it gives up compile-time type checking: the values do not have to share
// a type.
func printValue(value ...any) {
	fmt.Println(value...)
}

// NOTE: Type constraints are usually written as interfaces.
type Number interface {
	int | int64 | float64
}

func findGenericMin[Num Number](nums []Num) Num {
	if len(nums) == 0 {
		return -1
	}

	min := nums[0]
	for _, num := range nums {
		if num < min {
			min = num
		}
	}
	return min
}

func main() {
	print(123)
	printValue(123)
	printValue(123.3)
	printValue(false)
	printValue("Hello")

	fmt.Println(findGenericMin([]int{123, 4, 5, 56, 344}))
	fmt.Println(findGenericMin([]float64{123.34, 4.34, 5.43, 56.34, 344.33}))

}
