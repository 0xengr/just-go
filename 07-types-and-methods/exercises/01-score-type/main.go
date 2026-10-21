package main

import "fmt"

type Score int

func (s Score) Pass() bool {
	return s >= 60
}

func main() {
	var s Score = 75
	fmt.Println(s, s.Pass())
}
