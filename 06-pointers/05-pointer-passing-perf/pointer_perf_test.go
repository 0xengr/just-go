package main

import "testing"

// NOTE: run the benchmark from this folder with:
//
//	go test -bench=. -benchmem
//
// PERF: both benchmarks read a single byte from the same 1 MB struct, so the
// only thing they measure is how that struct reaches the function. Measured
// with go version 1.27.1; the absolute numbers move with the machine, so read
// the gap between the two rows rather than either number on its own:
//
//	BenchmarkByValue-4     10846     126240 ns/op   1007620 B/op   1 allocs/op
//	BenchmarkByPointer-1e9 1000000000      0.5125 ns/op         0 B/op   0 allocs/op
//
// WARNING: the pointer row sits near zero because the compiler inlines
// byPointer and then drops the unused result, so it is really measuring an
// empty loop. The honest reading is "a million bytes copied" against "nothing
// copied", not a literal two-hundred-thousand-times speedup.
//
// NOTE: keep the machine details out of the chapter notes. They describe only
// whoever happened to run this, not what a reader will see.

type data struct {
	Buffer [1_000_000]byte
}

func byValue(d data) byte {
	return d.Buffer[0]
}

func byPointer(d *data) byte {
	return d.Buffer[0]
}

func BenchmarkByValue(b *testing.B) {
	d := data{}

	// NOTE: b.ReportAllocs() adds the B/op and allocs/op columns, which is
	// where the real cost of copying shows up.
	b.ReportAllocs()
	for b.Loop() {
		_ = byValue(d)
	}
}

func BenchmarkByPointer(b *testing.B) {
	d := data{}

	// NOTE: same body, pointer instead of value, so the only difference the
	// benchmark measures is the copying itself.
	b.ReportAllocs()
	for b.Loop() {
		_ = byPointer(&d)
	}
}
