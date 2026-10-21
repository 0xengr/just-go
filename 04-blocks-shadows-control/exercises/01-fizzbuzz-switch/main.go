package main

import "fmt"

func main() {
	for n := 1; n <= 15; n++ {
		switch {
		case n%15 == 0:
			fmt.Println(n, "FizzBuzz")
		case n%3 == 0:
			fmt.Println(n, "Fizz")
		case n%5 == 0:
			fmt.Println(n, "Buzz")
		default:
			fmt.Println(n)
		}
	}
}
