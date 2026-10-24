package main

import (
	"fmt"
	"math"
)

// NOTE: Polymorphism in Go
// * Polymorphism is the ability to appear in various forms. For example, a shape
// * can appear as a square, rectangle, or triangle.
// * Polymorphism is not specific to Go; it is a paradigm for writing code.
// * It helps decouple code, and in Go, with the power of interfaces,
// * we can achieve polymorphism easily.

// Shape interface
type Shape interface {
	Area() float64
}

type Circle struct {
	Radius float64
}
type Rectangle struct {
	Length float64
	Width  float64
}
type Square struct {
	Length float64
}

// Example: Let's calculate the area of a few shapes without interfaces
func areaOfCircle(c Circle) {
	area := math.Pi * c.Radius * c.Radius
	fmt.Printf("Area of circle: %.2f\n", area)
}

func areaOfRectangle(r Rectangle) {
	area := r.Length * r.Width
	fmt.Printf("Area of rectangle: %.2f\n", area)
}
func areaOfSquare(s Square) {
	area := s.Length * s.Length
	fmt.Printf("Area of square: %.2f\n", area)
}

// NOTE: Notice the difference in the area functions above.
// * They are good and solve our problem, but what if we have to write
// * the area of more shapes? Then the functions never end.
// * So let's use an interface here: you just add a method receiver to a struct
// * and it will implement the Shape interface.
//
// Adding a method receiver to each struct
func (c Circle) Area() float64 {
	return math.Pi * c.Radius * c.Radius
}
func (r Rectangle) Area() float64 {
	return r.Length * r.Width
}
func (s Square) Area() float64 {
	return s.Length * s.Length
}

func areaPrint(s Shape) {
	fmt.Println(s.Area())
}

func main() {
	circle := Circle{Radius: 12.2}
	rectangle := Rectangle{Length: 10, Width: 6}
	square := Square{Length: 4}

	areaOfCircle(circle)
	areaOfRectangle(rectangle)
	areaOfSquare(square)

	areaPrint(circle)
	areaPrint(rectangle)
	areaPrint(square)
}
