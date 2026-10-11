// NOTE: a slice header is just (pointer, len, cap), so every alias points at the
// same memory. The three-index form caps capacity, which forces the next append
// to copy rather than write over data someone else is holding.
package main

import "fmt"

func main() {
	base := make([]int, 3, 5) // room to spare: cap 5
	view := base[:2]
	view[0] = 100
	fmt.Println("base after writing through view:", base)

	grown := append(view, 7) // cap 5 is enough: same backing array
	grown[2] = 200
	fmt.Println("base after writing through grown:", base)

	// WARNING: a[i:j:k] caps capacity, so the next append must copy and your
	// caller can no longer be surprised by writes through the returned slice.
	fixed := base[2:3:5]     // cap is exactly 1
	over := append(fixed, 9) // cap exhausted: new backing array
	over[0] = -1
	fmt.Println("base after a capped append:", base)
	fmt.Println("over (separate array):", over)
}
