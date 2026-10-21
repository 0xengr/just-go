package main

import "fmt"

func main() {
	fuel := 10
	if fuel > 5 {
		fuel = 5
	}
	fmt.Println("after:", fuel)
}
