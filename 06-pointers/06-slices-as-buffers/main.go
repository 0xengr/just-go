package main

import (
	"bytes"
	"fmt"
)

// NOTE: Slices as Buffers
// Imagine you are building an HTTP API. A client sends
// a request body, and your server must read it.
//
// A straightforward approach is to allocate a new byte slice
// every time you need a buffer:
//
// example: allocating a new byte slice every time (Not efficient)
func processRequest() {
	buf := make([]byte, 32*1024)

	// Read request data into buf
	// process data

	// PERF: this 32 KB is allocated fresh on every call, and the request only
	// ever fills the first few bytes of it. At a few thousand requests a second
	// that is a lot of short-lived garbage for the collector to chase. Printing
	// len and cap instead of the contents, since 32 KB of zeros is unreadable.
	fmt.Println("fresh buffer:", len(buf), "bytes, cap", cap(buf))
}

// NOTE: the counter-example. This is the trap that makes buffers dangerous:
// `copy` returns how many bytes it wrote, and the bytes past that are left
// over from the previous message. Forget `[:n]` and a 5-byte reply arrives
// padded with the tail of a 5-byte greeting.
func staleBytesDemo() {
	buffer := make([]byte, 8)

	for _, msg := range []string{"Hello", "World"} {
		n := copy(buffer, msg)
		fmt.Printf("%q\n", buffer[:n])
	}

	// PERF: the buffer's len is still 8 after two copies. `copy` never resizes
	// it, so the honest view of what was written is `buffer[:n]`. Printing
	// buffer directly would show the greeting's leftovers every time.
	fmt.Printf("len is still %d, and the whole array holds %q\n", len(buffer), buffer)
}

// NOTE: you do not have to hand-roll buffer management. bytes.Buffer manages
// its own storage and grows as needed, so this reads like string building but
// keeps the bytes reusable across operations.
func stdlibBuffer() {
	var buf bytes.Buffer

	buf.WriteString("Hello ")
	buf.WriteString("Go")

	fmt.Println(buf.String())

	// PERF: reuse means resetting, not reallocating. `Reset` keeps the
	// backing array and just sets len back to zero, so the next write reuses
	// the same memory. Assigning `buf = bytes.Buffer{}` would throw it away and
	// allocate again.
	buf.Reset()
	buf.WriteString("reused")
	fmt.Println("after Reset:", buf.String())
}

func main() {
	// Example: Allocating a new byte slice one time and reuse it.

	// make([]byte, 8) creates a slice with eight bytes of length and capacity.

	// copy(buf, message) copies the bytes of the string into the existing backing array.

	// buf[:n] creates a smaller slice exposing only the bytes actually written.

	buffer := make([]byte, 8)

	for _, msg := range []string{"Hello", "World"} {
		n := copy(buffer, msg)
		fmt.Printf("%q\n", buffer[:n])
	}

	// WARNING: nothing above allocated twice. The second iteration wrote into
	// the same backing array the first one used, so the 8 bytes were made once
	// and reused twice.

	staleBytesDemo()
	stdlibBuffer()

	processRequest()
}
