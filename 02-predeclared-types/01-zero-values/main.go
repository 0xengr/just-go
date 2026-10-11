// NOTE: every type has a zero value worth using, so a struct is best designed
// so that `var p player` is already correct and no constructor is needed.
package main

import "fmt"

type player struct {
	name  string
	score int
	ready bool
}

func main() {
	var p player
	fmt.Printf("%q %d %t\n", p.name, p.score, p.ready)

	var inbox []string
	fmt.Println("nil slice len:", len(inbox), "is nil:", inbox == nil)
	inbox = append(inbox, "welcome") // append works on nil slices
	fmt.Println(inbox)
}
