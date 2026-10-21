package main

import "fmt"

type Point struct {
	X, Y int
}

func main() {
	p := Point{X: 1, Y: 2}
	q := p
	q.X = 99
	fmt.Println("original:", p)
	fmt.Println("copy:", q)
}
