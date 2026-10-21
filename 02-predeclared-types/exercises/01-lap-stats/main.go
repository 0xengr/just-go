package main

import "fmt"

func main() {
	finished, total := 7, 10
	rate := float64(finished) / float64(total) * 100
	fmt.Printf("finished %d/%d = %.1f%%\n", finished, total, rate)
}
