// NOTE: converting a string to []byte or []rune always copies, so the edits
// below leave the original alone. For building a string up in a loop, reach for
// strings.Builder rather than s += x, which reallocates every time.
package main

import (
	"fmt"
	"strings"
	"unicode/utf8"
)

func main() {
	name := "héllo"
	fmt.Println("bytes:", len(name), "runes:", utf8.RuneCountInString(name))

	bs := []byte(name) // copy of the underlying bytes
	bs[0] = 'H'
	fmt.Println("from []byte:", string(bs), "| original:", name)

	rs := []rune(name) // decoded characters
	rs[1] = 'E'
	fmt.Println("from []rune:", string(rs), "chars:", len(rs))

	var sb strings.Builder
	// PERF: Grow reserves the room once, so the loop below never reallocates.
	sb.Grow(16)
	for i := range 3 {
		fmt.Fprintf(&sb, "part%d-", i)
	}
	fmt.Println("builder:", strings.TrimSuffix(sb.String(), "-"))
}
