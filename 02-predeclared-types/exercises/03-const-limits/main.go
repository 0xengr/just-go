package main

import "fmt"

const limit = 90_000

func main() {
	var asInt int = limit
	var asFloat float64 = limit
	fmt.Println("int:", asInt, "float64:", asFloat)
}
