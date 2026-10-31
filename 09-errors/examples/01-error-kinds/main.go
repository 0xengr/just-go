package main

import "fmt"

// NOTE: Go has three broad kinds of errors:
// *
// *  1. Syntax errors   - the compiler rejects the code before it runs.
// *  2. Runtime errors  - the program panics while running.
// *  3. Semantic errors - the code builds and runs, but the logic is wrong.

// NOTE: Syntax errors come from writing code the language does not allow.
// * The compiler catches them before the program runs, and editors that use
// * gopls underline them as you type.
// * Go is case-sensitive: `println` is not `Println`, and `fmt.println` does
// * not exist.

// NOTE: Runtime errors appear while the program runs, not while it compiles.
// * Reading past the end of a slice, dereferencing nil, dividing an integer
// * by zero, or opening a missing file all fail during execution.
// * Go reports most of these as a panic, which stops the program unless it is
// * recovered.

// NOTE: Semantic errors compile and run, but the logic produces the wrong
// * result. The compiler cannot catch them; tests and careful reading do.

// NOTE: shouldTakeCarBuggy gets the boundary wrong: at exactly 2 km it says
// * to walk, although the car should be taken.
func shouldTakeCarBuggy(km int) string {
	if km > 2 {
		return "Take the car"
	}
	return "Going to walk today"
}

// NOTE: shouldTakeCar treats 2 km as the boundary.
func shouldTakeCar(km int) string {
	if km >= 2 {
		return "Take the car"
	}
	return "Going to walk today"
}

func main() {
	// BUG: uncommenting the next line does not compile:
	// fmt.println("Enter your city:")
	// undefined: fmt.println (the package has Println, not println)
	fmt.Println("Enter your city:")

	nums := []int{2, 4, 6, 8}
	total := 0
	for i := range nums {
		total += nums[i]
	}

	// BUG: uncommenting this loop reads past the end of nums and panics:
	// for i := 0; i <= 10; i++ {
	//     total += nums[i]
	// }
	// panic: runtime error: index out of range [4] with length 4

	// NOTE: The safe loop above stops at the end of the slice; the commented
	// * one would keep going past the end and panic at index 4.
	fmt.Println("Total:", total)

	// NOTE: The semantic error runs quietly and prints the wrong answer.
	fmt.Println("buggy:", shouldTakeCarBuggy(2))
	fmt.Println("fixed:", shouldTakeCar(2))
}
