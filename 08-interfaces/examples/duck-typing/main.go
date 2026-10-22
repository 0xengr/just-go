package main

import "fmt"

// NOTE: DUCK TYPING
// * Duck typing takes its name from a test in computer science. It says:
// * If it looks like a duck, swims like a duck, and quacks like a duck, then it might as well be a duck.
// * In Go there is no "implements" keyword: a type satisfies an interface by having its methods.

// Let's understand with an example.

type Quacker interface {
	Quack()
}

type RealDuck struct{}
type ToyDuck struct{}
type RobotDuck struct{}

func (d RealDuck) Quack() {
	fmt.Println("RealDuck: Quack, Quack, Quack")
}
func (t ToyDuck) Quack() {
	fmt.Println("ToyDuck: Quack, Quack, Quack")
}
func (r RobotDuck) Quack() {
	fmt.Println("RobotDuck: Quack, Quack, Quack")
}

// using the Quacker interface: any value with a Quack method works here
func MakeItQuack(q Quacker) {
	q.Quack()
}

// or an even shorter, variadic form: pass as many Quackers as you like
func makeItQuack(quack ...Quacker) {
	for _, q := range quack {
		q.Quack()
	}
}

// How this code uses duck typing:
// None of RealDuck, ToyDuck, or RobotDuck says it implements Quacker.
// Each only defines Quack(), which is enough for Go to treat it as a Quacker.
// MakeItQuack and makeItQuack then accept any of them through the interface.

func main() {
	d := RealDuck{}
	t := ToyDuck{}
	r := RobotDuck{}

	MakeItQuack(d)
	MakeItQuack(t)
	MakeItQuack(r)

	// quack for all three at once
	makeItQuack(d, t, r)

}
