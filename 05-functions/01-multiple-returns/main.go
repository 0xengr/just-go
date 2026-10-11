package main

import "fmt"

// NOTE: Multiple Return values in GO.
// * One of the best feature of Go is Multiple return values
// * You can return more than one values from Go functions
// * which is not possible in most of the programming languages

// Example: lookup user

type user struct {
	Name  string
	Email string
}

var users = map[string]user{
	"001": {Name: "Alice", Email: "alice@example.com"},
	"002": {Name: "John", Email: "john@example.com"},
	"003": {Name: "Jane", Email: "jane@example.com"},
}

func lookupUser(id string) (user, bool) {
	user, found := users[id]
	return user, found
}

func main() {
	u, ok := lookupUser("001")

	if !ok {
		fmt.Println("Not found")
	} else {
		fmt.Printf("%s <%s>\n", u.Name, u.Email)
	}
}
