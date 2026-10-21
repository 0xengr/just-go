package main

import "fmt"

// NOTE: Go is call by value: Which means
// when we supply a variable for a parameter to a
// function, Go always makes a copy of the variable.
//
// Example:

type person struct {
	age  int
	name string
}

// modifyBasic receives copies of i, s and p.
// Important:
//
// # Go is always call by value
//
// That means:
//
//	i -> copied into the parameter i
//	s -> copied into the parameter s
//	p -> copied into the parameter p
//
// changing these parameters does NOT change the original variables.
func modifyBasic(i int, s string, p person) {
	i = 1 * 2
	s = "Goodbye"
	p.name = "Bob"
}

// modSlice receives a COPY of the slice value.
//
// A slice value contains information about underlying array
// (Conceptually: pointer + length + capacity)
//
// Therefore:
//
//	original slice ──┐
//	                 ├──> [1 2 3]
//	copied slice ────┘
//
// Both slice values refer to the same underlying array.
//
// SO modifying an element through the copied slice modifies
// the same underlying array seen by the original slice.
func modSlice(s []int) {
	// This modifies the underlying array.
	//
	// The caller will see these changes
	for i := range s {
		s[i] = s[1] * 2
	}

	// append returns a NEW slice value.
	//
	// We assign that new slice value only to the local parameter `s`.
	//
	// The caller's slice variable is NOT changed.
	s = append(s, 10)

	fmt.Println("inside modSlice:", s)
}

// modMap receives a COPY of the map value.
//
// The copied map value still refers to the same underlying map data.
//
// Therefore:
//
//	original map ──┐
//	               ├──> map data
//	copied map ────┘
//
// Modifying the map through the copied value changes the same
// underlying map that the caller sees.
func modMap(m map[int]string) {
	m[2] = "hello"
	m[3] = "goodbye"

	delete(m, 1)

	fmt.Println("inside modMap:", m)
}

// reassignSlice demonstrates an important distinction.
//
// The parameter `s` is a COPY of the caller's slice value.
//
// Reassigning the parameter changes only the local copy
func reassignSlice(s []int) {
	s = []int{100, 200, 300}
	fmt.Println("inside reassignSlice:", s)
}

// reassignMap demonstrate the same rule for maps.
//
// The map itself can be modifies through the parameter,
// but reassigning the parameter doesn't change the caller's map.
func reassignMap(m map[string]int) {
	m = map[string]int{
		"new": 100,
	}

	fmt.Println("inside reassignMap:", m)
}

func main() {
	fmt.Println("========== 1. BASIC VALUES ==========")

	i := 2
	s := "Hello"

	p := person{
		age:  20,
		name: "Alice",
	}

	modifyBasic(i, s, p)

	fmt.Println("outside:", i, s, p)

	/*
		Output:

		outside: 2 Hello {20 Alice}

		Why?

		The function received copies:

		    i -> copy
		    s -> copy
		    p -> copy

		Changing those copies doesn't affect the originals.
	*/

	fmt.Println("\n========== 2. SLICE ==========")

	numbers := []int{1, 2, 3}

	fmt.Println("before:", numbers)

	modSlice(numbers)

	fmt.Println("after:", numbers)

	/*
		Output:

		before: [1 2 3]
		inside modSlice: [2 4 6 10]
		after: [2 4 6]

		Notice something interesting:

		    The elements changed.
		    The length did NOT change.

		Why?

		The slice itself was copied:

		    numbers
		       │
		       ▼
		    ┌───────────────┐
		    │ ptr │ len │ cap│
		    └─────┬─────────┘
		          │
		          ▼
		        [2 4 6]

		        ▲
		        │
		    parameter s
		    is another slice value

		Both slice values refer to the same underlying array.

		But:

		    s = append(s, 10)

		only changes the LOCAL slice variable `s`.

		The caller's `numbers` still has length 3.
	*/

	fmt.Println("\n========== 3. SLICE REASSIGNMENT ==========")

	numbers = []int{1, 2, 3}

	reassignSlice(numbers)

	fmt.Println("outside:", numbers)

	/*
		Output:

		inside reassignSlice: [100 200 300]
		outside: [1 2 3]

		Again:

		    s = []int{100, 200, 300}

		changes only the local copy of the slice value.
	*/

	fmt.Println("\n========== 4. MAP ==========")

	users := map[int]string{
		1: "Alice",
		2: "John",
	}

	fmt.Println("before:", users)

	modMap(users)

	fmt.Println("after:", users)

	/*
		Output:

		before: map[1:Alice 2:John]
		inside modMap: map[2:hello 3:goodbye]
		after: map[2:hello 3:goodbye]

		The map value is copied when passed to modMap.

		But both map values refer to the same underlying map data.

		So:

		    m[2] = "hello"

		modifies the same map that `users` refers to.
	*/

	fmt.Println("\n========== 5. MAP REASSIGNMENT ==========")

	usersByName := map[string]int{
		"Alice": 20,
	}

	reassignMap(usersByName)

	fmt.Println("outside:", usersByName)

	fmt.Println("outside:", users)

	/*
		Output:

		inside reassignMap: map[new:100]
		outside: map[Alice:20]

		Why?

		Because:

		    m = map[string]int{"new": 100}

		only replaces the LOCAL copy of the map value.

		It does not replace the caller's `users` variable.
	*/
}
