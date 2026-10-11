// NOTE: go run ./greet [name]
//
// os.Args[0] holds the program name, so the first real argument is at index 1.
package main

import (
	"fmt"
	"os"
)

func main() {
	name := "Gopher"
	if len(os.Args) > 1 && os.Args[1] != "" {
		name = os.Args[1]
	}
	fmt.Printf("Hello, %s!\n", name)
}
