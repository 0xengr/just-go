// NOTE: GOOS and GOARCH are baked in when you build, not when you run, so the
// same source reports whatever you targeted:
// GOOS=windows GOARCH=amd64 go build ./...
package main

import (
	"fmt"
	"runtime"
)

func main() {
	fmt.Println("go version:", runtime.Version())
	fmt.Println("os/arch:  ", runtime.GOOS+"/"+runtime.GOARCH)
	fmt.Println("compilers:", runtime.NumCPU(), "logical CPUs visible")
}
