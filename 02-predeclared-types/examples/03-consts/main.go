// NOTE: an untyped constant takes whatever type the context asks for, which is
// how one limit serves both an int and a float64. Giving it a type ends that
// freedom.
package main

import "fmt"

const (
	qualifyingMs = 90_000 // untyped: fits int, float64, etc.
	raceLaps     = 5
)

const trackName string = "Harbor Loop" // typed: always a string

func main() {
	var cutoff int = qualifyingMs
	var cutoffF float64 = qualifyingMs // same constant, both work (untyped)
	fmt.Println(trackName, "laps:", raceLaps, "cutoff:", cutoff, cutoffF)
}
