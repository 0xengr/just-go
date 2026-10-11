package main

import "fmt"

// NOTE: Anonymous functions and closures.
//
// * An anonymous function (Go calls it a "func literal") is a function with no
//   name. `func(n int) int { return n * 2 }` is everything it needs to exist.
//
// * A closure is a func literal that also remembers the variables written
//   around it, even after the function it was written in has returned.
//
// * Picture it like this: the code is a machine, the variables it remembers
//   are its settings. makeAdder(10) hands back a machine whose setting is
//   "always add 10", so you never pass 10 in again.

// makeAdder returns a function that adds the same number to whatever it gets.
// `alwaysAdd` is the setting this machine remembers.
func makeAdder(alwaysAdd int) func(int) int {
	return func(n int) int {
		return n + alwaysAdd
	}
}

// newCounter returns a function that counts how many times you have called it.
// `count` is the memory it keeps between calls.
func newCounter() func() int {
	count := 0
	return func() int {
		count++
		return count
	}
}

// each hands every number in nums to fn and collects whatever fn returns.
// fn is just a parameter here, so any function with this shape will do.
func each(nums []int, fn func(int) int) []int {
	out := make([]int, 0, len(nums))
	for _, n := range nums {
		out = append(out, fn(n))
	}
	return out
}

func main() {
	nums := []int{1, 2, 3}

	// 1. A function is a value, so it can live in a variable like an int does.
	// A normal named function is a value too, this one just skips the name.
	double := func(n int) int { return n * 2 }
	fmt.Println("double(5)             ->", double(5))

	// 2. A function can be handed to another function, either by variable or
	// written on the spot. `each` never learns which one it was given.
	fmt.Println("each(nums, double)    ->", each(nums, double))
	fmt.Println("each(nums, +10)       ->", each(nums, func(n int) int { return n + 10 }))

	// 3. "Anonymous" does not mean unusable: you can call one right away
	// without ever putting it in a variable.
	fmt.Println("add 2 and 3 now       ->", func(a, b int) int { return a + b }(2, 3))

	// 4. Here is the closure part: addTen remembers the 10 we built it with,
	// so both calls below add 10 even though 10 is not passed in again.
	addTen := makeAdder(10)
	fmt.Println("addTen(5)             ->", addTen(5))
	fmt.Println("addTen(100)           ->", addTen(100))

	// 5. A closure can also remember what happened between calls, so it holds
	// state: 1, then 2, then 3.
	count := newCounter()
	fmt.Println("count() three times   ->", count(), count(), count())

	// A brand new counter starts over, because calling newCounter() made a
	// brand new `count` variable for this one only.
	other := newCounter()
	fmt.Println("other()               ->", other())

	// 6. Closures are values, so a slice of them is fine too. These three all
	// have the same type, they just remember different numbers.
	numbers := []int{2, 3, 4}
	adders := make([]func(int) int, 0, len(numbers))
	for _, n := range numbers {
		adders = append(adders, func(x int) int { return x + n })
	}

	// WARNING: what each closure above remembers is the loop's `n`. Go 1.22
	// and later give every iteration its own copy, so the three really do
	// add 2, 3 and 4. With an older `go 1.21` line in go.mod they would all
	// have shared the last value and printed 104 three times.
	for i, add := range adders {
		fmt.Printf("%d + 100              -> %d\n", numbers[i], add(100))
	}

	// WARNING: a func value can be compared to nil and to nothing else.
	// `f == g` does not even compile, not for two identical functions.
	var missing func(int) int
	fmt.Println("missing == nil        ->", missing == nil)

	// BUG: calling an empty func value crashes the program, which is exactly
	// what the nil check above is there to stop.
	// missing(1)
}
