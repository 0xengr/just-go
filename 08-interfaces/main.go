package main

import "fmt"

// NOTE: Interfaces in Go
// * An interface is a set of methods that describe the behavior of a data type.
// * An interface never describes the implementation.

// Example:
type speaker interface {
	speak() string
}

type cat struct{}
type dog struct{}
type person struct{ name string }

func (c cat) speak() string {
	return "Meow Meow"
}

func (d dog) speak() string {
	return "woof woof"
}

func (p person) speak() string {
	return "Hi my name is: " + p.name
}

// NOTE: Now see how interfaces help us.
// Suppose you want to print all three of these `speak` methods; without an
// interface, you would do it like this:
func catSpeak(c cat) {
	fmt.Println(c.speak())
}
func dogSpeak(d dog) {
	fmt.Println(d.speak())
}
func personSpeak(p person) {
	fmt.Println(p.speak())
}

// NOTE: The problem here.
// * All structs have the `speak` method.
// * To use it, we created three separate functions, one for each type.
// * But an interface can make this better: all three implement the
// * `speaker` interface, so we can leverage it like this.

func speaking(s speaker) {
	fmt.Println(s.speak())
}

// or
// using a variadic function
func saySomething(say ...speaker) {
	for _, s := range say {
		fmt.Println(s.speak())
	}

}

func main() {

	c := cat{}
	d := dog{}
	p := person{name: "John Doe"}

	catSpeak(c)    // Meow Meow
	dogSpeak(d)    // woof woof
	personSpeak(p) // Hi my name is: John Doe

	// NOTE: Now use the interface method here to make it better and easier.
	speaking(c)
	speaking(d)
	speaking(p)

	// using a variadic function
	saySomething(c, d, p)

}
