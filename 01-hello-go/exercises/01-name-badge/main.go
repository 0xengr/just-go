package main

import (
	"fmt"
	"os"
	"runtime"
)

func main() {
	name := "Gopher"
	if len(os.Args) > 1 && os.Args[1] != "" {
		name = os.Args[1]
	}
	fmt.Printf("Hello, %s! (built with %s)\n", name, runtime.Version())
}
