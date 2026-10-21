// NOTE: len counts bytes, so a multi-byte letter reports 2. Count real letters
// with utf8.RuneCountInString, or range the string and get one rune per letter.
package main

import (
	"fmt"
	"unicode/utf8"
)

func main() {
	tag := "café"
	fmt.Println("bytes:", len(tag), "letters:", utf8.RuneCountInString(tag))

	// WARNING: the index is a byte offset, so it jumps by 2 for "é",
	// while r is a rune (int32) holding the decoded character.
	for i, r := range tag {
		fmt.Printf("byte %d -> %c\n", i, r)
	}
}
