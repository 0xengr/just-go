// NOTE: structs are values, so assigning one copies all of it, and == only
// compiles when every field is comparable. Struct tags are the one thing
// conversion ignores, so a tagged and an untagged struct convert freely.
package main

import "fmt"

type Point struct {
	X, Y int
}

type Labeled struct {
	Point // embedded, so X and Y are promoted to Labeled
	Label string
}

type Weighted struct {
	Point
	Weight float64 // tags are ignored by conversions
}

func main() {
	p := Point{X: 1, Y: 2}
	q := p // value copy
	q.X = 99
	fmt.Println("p:", p, "q:", q)

	l := Labeled{Point: p, Label: "start"}
	l.Y = 7 // promoted field
	fmt.Println("labeled:", l, "X:", l.X, "Y:", l.Y)

	// == compares every field
	fmt.Println("p == Point{1, 2}:", p == Point{1, 2})
	fmt.Println("l == equal literal:", l == Labeled{Point: Point{X: 1, Y: 7}, Label: "start"})

	// BUG: a struct holding a slice has no == at all.
	// _ = Bag{} == Bag{} would not compile: non-comparable field

	anon := struct {
		Name string
		Age  int
	}{Name: "zay", Age: 30}
	fmt.Println("anonymous struct:", anon)

	// convert between struct types with identical fields
	same := struct{ X, Y int }{1, 2}
	fmt.Println("anonymous -> named:", Point(same) == p)
	fmt.Println("named -> weighted:", Weighted{Point: p, Weight: 2.5})
}
