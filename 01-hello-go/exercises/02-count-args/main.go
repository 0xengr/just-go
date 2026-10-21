package main

import (
	"fmt"
	"os"
)

func main() {
	fmt.Println("argument count:", len(os.Args))
	for i, arg := range os.Args {
		fmt.Printf("%d: %s\n", i, arg)
	}
}
