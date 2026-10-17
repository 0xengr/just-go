// NOTE: len 0 and nil are different, and both range fine. Re-slicing to [:0] is
// the cheapest reset but leaves the old values reachable; drop the slice
// entirely when they hold anything you would rather the collector reclaim.
package main

import "fmt"

func main() {
	items := []string{"a", "b", "c"}

	// WARNING: [:0] reuses the array, so the old values stay reachable
	// through the backing array until the slice is overwritten or dropped.
	reused := items[:0] // len 0, same backing array: cheapest
	reused = append(reused, "x")
	fmt.Println("reused:", reused, "cap:", cap(reused), "still aliases items:", len(items) == cap(items))

	fresh := make([]string, 0, len(items)) // new array, keeps the capacity
	fresh = append(fresh, "y")
	fmt.Println("fresh:", fresh)

	items = nil // drop the reference so the memory can be collected
	fmt.Printf("nil slice: %v len=%d isNil=%t\n", items, len(items), items == nil)
}
