package main

import "fmt"

// NOTE: Generics in Go
// * Let's not memorize 300 lines about what generics are. The meaning is simple:
// * Generics let you write reusable, type-safe code that works with different data types without duplicating the same logic for each type.

// NOTE: Example without generics
// We need to find the max integer and the max float, so we have to write two separate functions.

func findMaxInt(nums []int) int {
	if len(nums) == 0 {
		return -1
	}
	max := nums[0]

	for _, num := range nums {
		if num > max {
			max = num
		}
	}

	return max
}

func findMaxFloat(nums []float64) float64 {
	if len(nums) == 0 {
		return -1
	}
	max := nums[0]

	for _, num := range nums {
		if num > max {
			max = num
		}
	}

	return max
}

// NOTE: Now the generic version
func findMaxGeneric[Num int | float64](nums []Num) Num {
	if len(nums) == 0 {
		return -1
	}
	max := nums[0]

	for _, num := range nums {
		if num > max {
			max = num
		}
	}

	return max
}

func main() {
	numSlice := []int{1, 2, 3, 4, 8, 5, 6}
	maxInt := findMaxInt(numSlice)
	fmt.Printf("Max Int value: %d\n", maxInt)

	floatSlice := []float64{1.3, 2.2, 3.3, 4.6, 8.6, 5.32, 6.21}
	maxFloat := findMaxFloat(floatSlice)
	fmt.Printf("Max Float value: %.2f\n", maxFloat)

	// NOTE: The generic function handles both int and float64 slices.
	fmt.Println(findMaxGeneric(numSlice))
	fmt.Println(findMaxGeneric(floatSlice))
}
