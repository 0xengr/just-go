package main

import "fmt"

func newCounter() func() int {
	n := 0
	return func() int {
		n++
		return n
	}
}

func main() {
	count := newCounter()
	fmt.Println(count(), count(), count())
}
