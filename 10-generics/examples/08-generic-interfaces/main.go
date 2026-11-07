package main

import "fmt"

// NOTE: An interface can have type parameters too
// * Comparator[T] compares a value with another value of the same type T.
// * Largest requires T to implement Comparator[T]. The two arguments must be
// * the same type, and that type must know how to compare itself.
// * This is the idea behind interfaces such as the standard library's
// * constraints on ordered values, written in a form a reader can see.

type Comparator[T any] interface {
	Compare(T) int
}

func Largest[T Comparator[T]](a, b T) T {
	if a.Compare(b) >= 0 {
		return a
	}
	return b
}

type Version struct {
	Major int
	Minor int
}

func (v Version) Compare(other Version) int {
	if v.Major != other.Major {
		return v.Major - other.Major
	}
	return v.Minor - other.Minor
}

func (v Version) String() string {
	return fmt.Sprintf("v%d.%d", v.Major, v.Minor)
}

func main() {
	fmt.Println(Largest(Version{Major: 1, Minor: 4}, Version{Major: 2, Minor: 0}))
	fmt.Println(Largest(Version{Major: 2, Minor: 1}, Version{Major: 2, Minor: 0}))
	fmt.Println(Largest(Version{Major: 1, Minor: 2}, Version{Major: 1, Minor: 2}))
}
