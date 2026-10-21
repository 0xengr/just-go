// NOTE: the complete for has init; condition; post. Init must use :=, the
// condition is checked before every lap, and the post runs after each one.
package main

import "fmt"

func main() {
	for i := 0; i < 5; i++ {
		fmt.Print(i, " ")
	}
	fmt.Println()

	// Any part can be left out. Hoist the init when it comes from earlier work:
	i := 0
	for ; i < 3; i++ {
		fmt.Print(i, " ")
	}
	fmt.Println()

	// Or drop the post when the step rule needs its own logic inside:
	for j := 0; j < 10; {
		fmt.Print(j, " ")
		if j%2 == 0 {
			j++
		} else {
			j += 2
		}
	}
	fmt.Println()

	// NOTE: prefer this form when you iterate a slice window rather than every
	// element; the bounds state the intent better than break/continue can.
	scores := []int{9, 7, 8, 6, 10}
	for k := 1; k < len(scores)-1; k++ {
		fmt.Print(scores[k], " ")
	}
	fmt.Println()
}
