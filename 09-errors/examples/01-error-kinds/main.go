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

// BUG: this line does not compile:
// fmt.println("Enter your city:")
// undefined: fmt.println (the package has Println, not println)

// NOTE: Runtime errors appear while the program runs, not while it compiles.
// * Reading past the end of a slice, dereferencing nil, dividing an integer
// * by zero, or opening a missing file all fail during execution.
// * Go reports most of these as a panic, which stops the program unless it is
// * recovered.

// BUG: this loop reads past the end of nums and panics:
// for i := 0; i <= 10; i++ {
//     total += nums[i]
// }
// panic: runtime error: index out of range [4] with length 4

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
	// NOTE: The syntax error above stays commented out, so this line must be
	// * Println with a capital P; a syntax error would stop the build.
	fmt.Println("Enter your city:")

	// NOTE: The runtime error above stays commented out, so this loop stops
	// * at the end of the slice instead of panicking.
	nums := []int{2, 4, 6, 8}
	total := 0
	for i := range nums {
		total += nums[i]
	}
	fmt.Println("Total:", total)

	// NOTE: The semantic error runs quietly and prints the wrong answer.
	fmt.Println("buggy:", shouldTakeCarBuggy(2))
	fmt.Println("fixed:", shouldTakeCar(2))
}
