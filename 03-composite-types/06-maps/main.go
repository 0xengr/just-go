// NOTE: a nil map reads fine and panics the moment you write to it. A missing
// key hands back the zero value, so comma ok is the only way to tell "absent"
// from "present and zero".
package main

import (
	"fmt"
	"sort"
)

func main() {
	var nilMap map[string]int
	fmt.Printf("nil map: len=%d isNil=%t read=%d\n", len(nilMap), nilMap == nil, nilMap["ana"])
	// BUG: writing to a nil map panics.
	// nilMap["ana"] = 31 would panic: assignment to entry in nil map

	ages := map[string]int{"ana": 31, "bo": 27}
	ages["cy"] = 24

	age, ok := ages["zz"]
	fmt.Println("missing key ->", age, ok)
	age, ok = ages["bo"]
	fmt.Println("present key ->", age, ok)

	delete(ages, "bo") // deleting a missing key is a no-op
	fmt.Println("after delete:", len(ages), ages)

	clear(ages) // Go 1.21+: empty it without rebuilding
	fmt.Println("after clear:", len(ages), "isNil:", ages == nil)

	// PERF: map iteration order is randomized per run (a hash seed plus jitter,
	// anti Hash-DoS armor), so sort keys for deterministic output. fmt.Println
	// sorts map keys for you when printing a map directly.
	scores := map[string]int{"c": 3, "a": 1, "b": 2}
	keys := make([]string, 0, len(scores))
	for k := range scores {
		keys = append(keys, k)
	}
	sort.Strings(keys)
	for _, k := range keys {
		fmt.Printf("%s=%d ", k, scores[k])
	}
	fmt.Println()
}
