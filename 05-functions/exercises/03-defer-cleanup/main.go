package main

import (
	"fmt"
	"os"
)

func fileSize(name string) (int, error) {
	f, err := os.Open(name)
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

func main() {
	n, err := fileSize("./examples/05-defer-func/hello.txt")
	fmt.Println(n, err)
}
