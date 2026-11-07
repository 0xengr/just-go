package main

import "fmt"

// NOTE: comparable and interface values
// * comparable accepts types whose values support == and !=. Slices, maps,
// * and functions do not. Interface types themselves are comparable, though,
// * so Same[Item] compiles even when the value inside the interface is a slice.
// * The comparison then panics at run time.

type Item interface {
	Name() string
}

type IntItem int

func (i IntItem) Name() string {
	return fmt.Sprintf("int:%d", i)
}

type ListItem []int

func (l ListItem) Name() string {
	return fmt.Sprintf("list:%v", l)
}

func Same[T comparable](a, b T) bool {
	return a == b
}

func main() {
	fmt.Println(Same(7, 7))
	fmt.Println(Same("go", "go"))
	fmt.Println(Same(IntItem(7), IntItem(7)))

	// BUG: a slice itself does not satisfy comparable, so this does not
	// compile:
	// fmt.Println(Same(ListItem{1, 2}, ListItem{1, 2}))
	// ListItem does not satisfy comparable

	// NOTE: Item is an interface, and an interface is comparable, so this
	// * compiles. It is safe only while the value inside is comparable.
	var a Item = IntItem(7)
	var b Item = IntItem(7)
	fmt.Println(Same(a, b))

	// BUG: uncommenting the next three lines compiles and then panics at run
	// time, because the interface values hold slices:
	// var c Item = ListItem{1, 2}
	// var d Item = ListItem{1, 2}
	// fmt.Println(Same(c, d))
	// panic: runtime error: comparing uncomparable type main.ListItem
}
