// NOTE: drop init and post and for becomes while; drop the condition too and
// it loops forever until break or return. Keep bodies flat with continue.
package main

import "fmt"

func main() {
	// Condition-only: the while shape, for values computed elsewhere.
	speed := 1
	for speed < 8 {
		fmt.Print(speed, " ")
		speed *= 2
	}
	fmt.Println()

	// Infinite loop with an explicit exit: the idiomatic do-while shape.
	// Java's do { work } while (cond) becomes for { work; if !cond break }.
	attempts := 0
	// WARNING: an infinite for needs a break or return in its body, or the
	// program never finishes.
	for {
		attempts++
		if attempts >= 3 {
			break
		}
	}
	fmt.Println("attempts:", attempts)

	// continue keeps the happy path left-aligned instead of nested.
	for lap := 1; lap <= 12; lap++ {
		if lap%2 == 0 && lap%3 == 0 {
			fmt.Println(lap, "pit-charge")
			continue
		}
		if lap%2 == 0 {
			fmt.Println(lap, "pit")
			continue
		}
		if lap%3 == 0 {
			fmt.Println(lap, "charge")
			continue
		}
		fmt.Println(lap, "green")
	}
}
