// NOTE: the length is part of an array's type, and assignment copies every
// element, so arrays show up as fixed buffers and in [N]byte style APIs.
package main

import "fmt"

func total(vals [5]int) int {
	var sum int
	for _, v := range vals {
		sum += v
	}
	return sum
}

func main() {
	var grid [3][2]int // 3 rows of 2 columns
	grid[1][0] = 7
	fmt.Println("grid:", grid, "rows:", len(grid), "cols:", len(grid[0]))

	scores := [5]int{9, 8, 7, 6, 5}
	backup := scores // full element copy, not a reference
	backup[0] = 0
	fmt.Println("scores:", scores, "backup:", backup, "sum:", total(scores))
}
