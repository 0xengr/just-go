// NOTE: nothing converts on its own, and int -> string is the conversion that
// bites: it produces a rune, not digits, so strconv.Itoa is what you want.
package main

import (
	"fmt"
	"strconv"
)

func main() {
	celsius := 21.7
	fmt.Println("truncated to int:", int(celsius)) // drops .7, no rounding

	bib := 42
	fmt.Println("bib number:", strconv.Itoa(bib)) // int -> decimal string

	raw := []byte("pit")
	fmt.Println("bytes:", raw, "back to string:", string(raw))
}
