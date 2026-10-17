// NOTE: maps can't be compared with == (only against nil). Use maps.Equal for
// comparable values, and reflect.DeepEqual when the values are slices or maps.
package main

import (
	"fmt"
	"maps"
	"reflect"
)

func main() {
	a := map[string]int{"x": 1, "y": 2}
	b := map[string]int{"x": 1, "y": 2}
	// BUG: a == b would not compile: invalid operation, maps are not comparable
	fmt.Println("a == nil:", a == nil)

	fmt.Println("maps.Equal:", maps.Equal(a, b))
	b["y"] = 99
	fmt.Println("after change:", maps.Equal(a, b))

	// WARNING: with slice values maps.Equal will not compile, so reach for
	// reflect.DeepEqual instead.
	c := map[string][]int{"x": {1, 2}}
	d := map[string][]int{"x": {1, 2}}
	fmt.Println("reflect.DeepEqual:", reflect.DeepEqual(c, d))
	d["x"][0] = 9
	fmt.Println("the two maps are independent:", c["x"], d["x"])
}
