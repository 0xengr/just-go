package main

import "fmt"

func main() {
	for i, r := range "pit_π!" {
		fmt.Println(i, r, string(r))
	}
}
