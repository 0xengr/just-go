// NOTE: integer division truncates toward zero, so an average has to be worked
// out in float64. Literals like 5 and 58.4 carry no type and take the one nearby.
package main

import "fmt"

func main() {
	finished, total := 7, 10
	// WARNING: int division drops the fraction, so 7/10 is 0, not 0.7.
	fmt.Println("truncated:", finished/total)

	rate := float64(finished) / float64(total) // convert one side first
	fmt.Printf("completion: %.1f%%\n", rate*100)

	var laps int = 5       // untyped literal 5 fits int
	var avg float64 = 58.4 // untyped literal fits float64
	fmt.Println("laps:", laps, "avg:", avg)
}
