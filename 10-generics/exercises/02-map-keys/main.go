package main

import (
	"fmt"
	"slices"
)

// NOTE: Exercise: collect the keys of a map.
// * Keys has two type parameters: K is the key type and V is the value type.
// * K must be comparable because map keys always are. V can be anything because
// * the function never looks at the values.
// * The result has type []K, so the caller keeps the real key type.

func Keys[K comparable, V any](m map[K]V) []K {
	keys := make([]K, 0, len(m))
	for key := range m {
		keys = append(keys, key)
	}
	return keys
}

func main() {
	ages := map[string]int{
		"Ada": 36,
		"Bob": 41,
	}
	names := map[int]string{
		1: "one",
		2: "two",
	}

	stringKeys := Keys(ages)
	slices.Sort(stringKeys)
	fmt.Println(stringKeys)

	intKeys := Keys(names)
	slices.Sort(intKeys)
	fmt.Println(intKeys)

	fmt.Println(Keys(map[string]int{}))
}
