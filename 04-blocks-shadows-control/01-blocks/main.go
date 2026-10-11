// NOTE: blocks nest. Each pair of braces sees its outer names, but names made
// inside vanish when the closing brace ends their block.
package main

import "fmt"

var season = "spring series" // package block: visible to the whole file

func main() {
	laps := 3 // function block: visible everywhere below in main

	if laps > 0 {
		// The if body is its own block: it reads outer names freely.
		fmt.Println(season, "laps:", laps)

		{
			// A bare block is a smaller scope still. Useful to fence off
			// a temporary variable so it cannot leak out.
			best := 58.4
			fmt.Println("best this stint:", best, "over", laps, "laps")
		}
		// BUG: best is gone once its block ends, so this would not compile:
		// fmt.Println(best)
	}

	for i := 0; i < 1; i++ {
		// Loop bodies are blocks too; i lives only for the statement.
		fmt.Println("loop i:", i, "still sees laps:", laps)
	}
	// BUG: i is scoped to the for statement, so this would not compile:
	// fmt.Println(i)
}
