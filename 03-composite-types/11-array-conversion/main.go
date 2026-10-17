// NOTE: arr[:] shares memory with the array behind it, so copy is the only route
// to an independent array. Turning a slice into an array needs the slice to be
// at least that long, and panics when it is not.
package main

import "fmt"

func main() {
	arr := [4]int{1, 2, 3, 4}
	sl := arr[:] // array to slice: shares the same memory
	sl[0] = 100
	fmt.Println("array after writing through the slice:", arr)

	var dup [4]int
	copy(dup[:], sl) // the only way to truly copy an array
	fmt.Println("independent copy:", dup)

	// slice to array: the slice must be at least as long as the array type
	back := [3]int(sl[:3])
	fmt.Println("slice to array:", back)

	defer func() { // recover so the demo does not crash
		if r := recover(); r != nil {
			fmt.Println("recovered from:", r)
		}
	}()
	tooBig := [8]int(sl) // WARNING: len(sl) is 4, so this panics
	fmt.Println("never reached:", tooBig)
}
