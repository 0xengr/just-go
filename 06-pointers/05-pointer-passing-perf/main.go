package main

import "fmt"

// NOTE: The pointer passing performance.
// * passing a pointer can improve performance
// * when struct is large enough, because a pointer
// * copies a fixed-size address rather than copying the entire struct

// NOTE: prefer to use pass by value when struct is small
// * Use pointer when large struct.
// * Pointers comes with cost: it can introduce indirection, aliasing, and
// * additional garbage-collection work.

// NOTE: Example: passing a value vs passing a pointer
// The benchmark for this lives in pointer_perf_test.go.

type Data struct {
	// NOTE: large data (1 million bytes)
	Buffer [1_000_000]byte
}

func processCopy(d Data) {
	// NOTE: Receives a copy of data, which is heavy operation.
}

func processPointerPassing(d *Data) {
	// NOTE: Receives a copy of pointer.
}

// NOTE: the other end of the story is what a pointer holds before anything is
// put there. A *Data can be a real address, and it can also be nothing at all.

func zeroValueVersusNoValue() {
	// WARNING: the zero value of a pointer is nil. It is not a pointer to a zero
	// Data, it is the absence of any Data, so there is no buffer to read and
	// dereferencing it panics.
	var nothing *Data

	fmt.Println("var p *Data is nil:", nothing == nil)

	// BUG: nothing.Buffer[0] would panic, so it stays commented out:
	// fmt.Println("reading through nil:", nothing.Buffer[0])

	// The guard that makes the dereference in the else branch safe.
	// if nothing != nil {
	// 	fmt.Println("reading through a real pointer:", nothing.Buffer[0])
	// }

	// NOTE: new(Data) is the other side of the same coin. It hands back a real
	// address holding a full 1 MB zero struct, so it is never nil and reading
	// through it is safe without a guard.
	fresh := new(Data)

	fmt.Println("new(Data) is never nil:", fresh != nil)
	fmt.Println("its zero value byte:", fresh.Buffer[0])

	// PERF: new(Data) is never free, though. It allocates the whole 1 MB on the
	// heap (`go build -gcflags=-m` reports "new(Data) escapes to heap"), which is
	// the very cost that &d below sidesteps: taking the address of an existing
	// value copies 8 bytes and allocates nothing.
}

func main() {

	d := Data{}
	processCopy(d)

	// NOTE: Now GO copies the pointer, not the entire struct.
	processPointerPassing(&d)

	zeroValueVersusNoValue()

}
