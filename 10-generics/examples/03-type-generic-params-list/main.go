package main

import "fmt"

// NOTE: Understanding K comparable and V any with maps
// * Imagine you're building an API where you need to store different kinds of data using maps:
// * Products: product ID → product details.
// * Users: user ID → user details.
// * Orders: order ID → order details.
// * Instead of writing separate functions for each map, we can write one generic function.

// Example:
func getValue[K comparable, V any](data map[K]V, key K) V {
	return data[key]
}

type Product struct {
	Name     string
	Price    int
	Discount int
}

type User struct {
	Name  string
	Email string
}

func main() {
	product := map[int]Product{
		1: {Name: "samsung s21", Price: 1200, Discount: 10},
		2: {Name: "iphone 14", Price: 1500, Discount: 5},
		3: {Name: "redmi note 12", Price: 300, Discount: 15},
		4: {Name: "oneplus 11", Price: 900, Discount: 10},
		5: {Name: "google pixel 8", Price: 800, Discount: 20},
	}

	user := map[string]User{
		"u1": {Name: "Mohammad", Email: "mohammad@example.com"},
		"u2": {Name: "Ali", Email: "ali@example.com"},
		"u3": {Name: "John", Email: "john@example.com"},
		"u4": {Name: "Sara", Email: "sara@example.com"},
		"u5": {Name: "Ayesha", Email: "ayesha@example.com"},
	}

	fmt.Println(getValue(product, 1))
	fmt.Println(getValue(user, "u2"))
	fmt.Println(getValue(user, "u3"))
	fmt.Println(getValue(product, 5))

}
