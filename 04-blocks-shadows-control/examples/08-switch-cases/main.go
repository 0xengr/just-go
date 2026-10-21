// NOTE: switch matches without fallthrough. List several values per case,
// scope a variable to every branch, and reach for a blank switch when the
// cases are boolean tests rather than equality checks.
package main

import "fmt"

func main() {
	callsigns := []string{"al", "bea", "rex", "pitlane", "chronometer"}

	for _, name := range callsigns {
		switch size := len(name); size {
		case 1, 2, 3:
			fmt.Println(name, "is short")
		case 4:
			fmt.Println(name, "fits exactly:", size)
		case 5, 6, 7:
			// An empty case means do nothing: mid-length names stay quiet.
		default:
			fmt.Println(name, "is long")
		}
	}

	// Blank switch: each case is its own boolean expression. Favor it over
	// an if/else chain when the cases are a related set of checks on one idea.
	for lap := 1; lap <= 15; lap++ {
		switch {
		case lap%2 == 0 && lap%3 == 0:
			fmt.Println(lap, "pit-charge")
		case lap%2 == 0:
			fmt.Println(lap, "pit")
		case lap%3 == 0:
			fmt.Println(lap, "charge")
		default:
			fmt.Println(lap, "green")
		}
	}

	// WARNING: break inside a case exits the switch, not the loop. Label the
	// loop (break laps) to stop everything from inside a case.
laps:
	for lap := 1; lap <= 10; lap++ {
		switch lap {
		case 2, 4, 6:
			fmt.Println(lap, "even")
		case 7:
			fmt.Println("red flag: leaving the session")
			break laps
		default:
			fmt.Println(lap, "running")
		}
	}
}
