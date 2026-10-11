package main

import (
	"errors"
	"fmt"
	"log"
	"os"
)

// NOTE: Defer function in GO
// * when we write programs, it often creates some temporary resources like file or network connections.
// * And yes that need to be cleaned up.
// * No matter how many exit points a function has or whether a function completed successfully or not.
// * In GO, we use `defer` keyword to cleanup.
// * `defer` is not only attach to cleanup. It runs when all the code inside function completely run.

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

	size, err := fileLen("./05-defer-func/hello.txt")
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("size of ./05-defer-func/hello.txt is: %d bytes", size)

}
