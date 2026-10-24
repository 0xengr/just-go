package main

import (
	"embed"
	"encoding/json"
	"fmt"
	"io"
	"strings"
)

// NOTE: example.json is embedded so the example runs from any directory.
//
//go:embed example.json
var exampleFile embed.FS

// NOTE: Accept interfaces, return concrete types.
// * The key idea is to give callers flexibility in what they provide,
// * while giving them a concrete, understandable result.

type Person struct {
	Name string `json:"name"`
	Age  int    `json:"age"`
}

// Example: A function that only accepts a string
func loadPerson2(s string) (Person, error) {
	var p Person

	err := json.NewDecoder(
		strings.NewReader(s),
	).Decode(&p)

	return p, err
}

// NOTE: The problem in the above code is:
// * it only accepts a string, so we call `strings.NewReader` to turn that
// * string into a reader
// * what if the data comes from a file, HTTP, or somewhere else?
// * we need a parameter that accepts all of those

func loadPerson(r io.Reader) (Person, error) {
	var p Person
	err := json.NewDecoder(r).Decode(&p)
	return p, err
}

func main() {

	// using loadPerson
	// source 1. A string in memory
	p1, err := loadPerson(strings.NewReader(`{"name":"john", "age":23}`))
	if err != nil {
		fmt.Println("Error:", err)
		return
	}
	fmt.Println(p1)

	// source 2. A file
	file, err := exampleFile.Open("example.json")
	if err != nil {
		fmt.Println("Error:", err)
		return
	}
	defer file.Close()

	p2, err := loadPerson(file)
	if err != nil {
		fmt.Println("Error:", err)
		return
	}
	fmt.Println(p2)

}
