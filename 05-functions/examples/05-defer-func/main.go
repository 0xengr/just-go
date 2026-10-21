package main

import (
	"errors"
	"fmt"
	"log"
	"os"
)

// NOTE: Defer function in Go
// * When we write programs, we often create temporary resources like files or
// * network connections.
// * And yes, those need to be cleaned up.
// * No matter how many exit points a function has, or whether it completed
// * successfully.
// * In Go, we use the `defer` keyword to clean up.
// * `defer` is not attached only to cleanup. It runs when all the code inside
// * the function has run to completion.

// NOTE: defer function runs in LIFO order
// Example
func runDefer() {
	defer fmt.Println("defer func 1")
	defer fmt.Println("defer func 2")
	defer fmt.Println("defer func 3")
	defer fmt.Println("defer func 4")
	defer fmt.Println("defer func 5")
}

// Example: a function read the file and returns file length
func fileLen(fileName string) (int, error) {
	if fileName == "" {
		return 0, errors.New("missing fileName to read")
	}

	f, err := os.Open(fileName)
	if err != nil {
		return 0, err
	}
	defer f.Close()

	info, err := f.Stat()
	if err != nil {
		return 0, err
	}

	return int(info.Size()), nil
}

// Example: a cat function: Unix utility for printing the contents of a file.

func main() {
	// if len(os.Args) < 2 {
	// 	log.Fatal("no file specified")
	// }

	// f, err := os.Open(os.Args[1])
	// if err != nil {
	// 	log.Fatal(err)
	// }
	// defer f.Close() // <--- defer here

	// data := make([]byte, 2048)

	// for {
	// 	// Read some bytes from f and put them into this data slice
	// 	count, err := f.Read(data)
	// 	os.Stdout.Write(data[:count])
	// 	if err != nil {
	// 		if err != io.EOF {
	// 			log.Fatal(err)
	// 		}
	// 		break
	// 	}
	// }

	runDefer()
	// output:
	// defer func 5
	// defer func 4
	// defer func 3
	// defer func 2
	// defer func 1

	size, err := fileLen("./examples/05-defer-func/hello.txt")
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("size of ./examples/05-defer-func/hello.txt is: %d bytes", size)

}
