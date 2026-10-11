// NOTE: a bare break or continue steers only its innermost loop. Label the
// outer loop when a whole heat must be skipped or abandoned from inside.
package main

import "fmt"

func main() {
	heats := []string{"q1", "q2"}
	lanes := []int{1, 2, 3}

outer:
	for _, heat := range heats {
		for _, lane := range lanes {
			if heat == "q1" && lane == 2 {
				fmt.Println(heat, "lane", lane, "false start: skipping heat")
				continue outer
			}
			fmt.Println(heat, "lane", lane, "clean")
		}
		fmt.Println(heat, "complete: all lanes ran")
	}

	// Same tool, other direction: break out of everything at once.
laps:
	for lap := 1; lap <= 5; lap++ {
		for stall := 1; stall <= 5; stall++ {
			if lap == 3 && stall == 2 {
				fmt.Println("red flag on lap", lap)
				break laps
			}
		}
	}
	fmt.Println("session halted")
}
