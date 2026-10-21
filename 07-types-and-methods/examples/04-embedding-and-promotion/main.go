package main

import "fmt"

// NOTE: Go encourages composition over inheritance via embedding.
type Employee struct {
	Name string
	ID   string
}

func (e Employee) Description() string {
	return fmt.Sprintf("%s (%s)", e.Name, e.ID)
}

type Manager struct {
	Employee // embedded field
	Reports  []Employee
}

func main() {
	m := Manager{
		Employee: Employee{Name: "Bob Bobson", ID: "12345"},
		Reports:  []Employee{},
	}

	// NOTE: Fields and methods are promoted.
	fmt.Println(m.ID)
	fmt.Println(m.Description())

	// NOTE: The containing type is not the embedded type.
	var e Employee = m.Employee // OK
	fmt.Println(e.Description())
	_ = e
}
