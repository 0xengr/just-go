package main

import (
	"fmt"
	"time"
)

// NOTE: Methods are functions with a receiver declared after func.
type Person struct {
	FirstName string
	LastName  string
	Age       int
}

func (p Person) String() string {
	return fmt.Sprintf("%s %s, age %d", p.FirstName, p.LastName, p.Age)
}

// NOTE: Use a pointer receiver when the method modifies the receiver.
type Counter struct {
	total       int
	lastUpdated time.Time
}

func (c *Counter) Increment() {
	c.total++
	c.lastUpdated = time.Now()
}

func (c Counter) String() string {
	return fmt.Sprintf("total: %d, last updated: %v", c.total, c.lastUpdated)
}

// NOTE: Go handles address and dereference automatically for local variables.
func doUpdateWrong(c Counter) {
	c.Increment()
	fmt.Println("wrong:", c.String())
}

func doUpdateRight(c *Counter) {
	c.Increment()
	fmt.Println("right:", c.String())
}

func main() {
	p := Person{FirstName: "Fred", LastName: "Fredson", Age: 52}
	fmt.Println(p.String())

	var c Counter
	fmt.Println(c.String())
	c.Increment()
	fmt.Println(c.String())

	// NOTE: A pointer instance can call both pointer and value receiver methods.
	pc := &Counter{}
	fmt.Println(pc.String())
	pc.Increment()

	// NOTE: Passing a value copies it, so the original is unchanged.
	doUpdateWrong(c)
	fmt.Println("after wrong:", c.String())
	doUpdateRight(&c)
	fmt.Println("after right:", c.String())
	_ = pc
}
