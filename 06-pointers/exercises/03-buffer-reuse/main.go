package main

import (
	"bytes"
	"fmt"
)

func main() {
	var buf bytes.Buffer
	buf.WriteString("Hello ")
	buf.WriteString("Go")
	fmt.Println(buf.String())
	buf.Reset()
	buf.WriteString("reused")
	fmt.Println(buf.String())
}
