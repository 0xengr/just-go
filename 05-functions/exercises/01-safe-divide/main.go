package main

import (
	"errors"
	"fmt"
)

func divide(a, b int) (int, error) {
	if b == 0 {
		return 0, errors.New("division by zero")
	}
	return a / b, nil
}

func main() {
	q, err := divide(12, 3)
	fmt.Println(q, err)
	q, err = divide(12, 0)
	fmt.Println(q, err)
}
