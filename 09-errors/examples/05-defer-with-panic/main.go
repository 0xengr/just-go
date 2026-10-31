package main

import (
	"fmt"
	"os"
)

// NOTE: How defer works with panic
// * Deferred functions run in LIFO order.
// * A panic does not skip them: while the panic unwinds the stack, every
// * deferred call still runs, and only then does the program crash.
// * That is why Go uses defer for cleanup: it runs in almost every situation.
// * To learn more about defer order, visit:
// * 05-functions/examples/05-defer-func/main.go

func main() {
	defer fmt.Println("defer first")
	defer fmt.Println("defer second")

	fmt.Println("BEFORE PANIC")

	// NOTE: readFile defers a close and then panics. Its deferred call runs
	// before the ones below because defers unwind LIFO across functions:
	// the last defer registered is the first one to run.
	readFile()

	fmt.Println("AFTER PANIC") // this won't run

	// output:
	// BEFORE PANIC
	// closing file
	// defer second
	// defer first
	// panic: unexpected failure
}

// NOTE: Opening a file
// * If the file opens successfully, the deferred cleanup runs even when the
// * function panics later, because the panic unwinds through the defer.
// * If the file does not open, err itself becomes the panic value.
func readFile() {
	file, err := os.Open("./examples/05-defer-with-panic/hello.txt")
	if err != nil {
		panic(err)
	}
	defer func() {
		fmt.Println("closing file")
		file.Close()
	}()

	panic("unexpected failure")
}
