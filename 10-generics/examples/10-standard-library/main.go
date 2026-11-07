package main

import (
	"cmp"
	"fmt"
	"maps"
	"slices"
)

// NOTE: Generics in the standard library
// * Go 1.21 added generic helpers for slices, maps, and comparison. Reach for
// * these before writing another loop by hand.
// * slices.Sort, slices.Contains, slices.Max, and slices.Min work on any
// * ordered element type. cmp.Compare returns -1, 0, or 1, which is the shape
// * a sort comparison needs. maps.Clone and maps.Equal take their key and value
// * types from the map argument.

func main() {
	numbers := []int{5, 3, 8, 1}

	slices.Sort(numbers)
	fmt.Println(numbers)
	fmt.Println(slices.Contains(numbers, 8))
	fmt.Println(slices.Max(numbers), slices.Min(numbers))

	words := []string{"go", "is", "fun"}
	slices.SortStableFunc(words, func(a, b string) int {
		return cmp.Compare(len(a), len(b))
	})
	fmt.Println(words)

	first := map[string]int{"a": 1}
	second := maps.Clone(first)
	second["a"] = 2

	fmt.Println(first["a"], second["a"])
	fmt.Println(maps.Equal(first, second))
}
