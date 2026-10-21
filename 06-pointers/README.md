# 06 — Pointers

The big idea: Go passes everything by value, so a function can never reach back
and change the variable you handed it. Pointers are how you hand over an address
instead of a copy — and, because Go has a garbage collector, that address is
just a note in a notebook, not a promise to free anything by hand. Think of it
as writing an address on a sticky note: the note is not the house, but you can
send someone to the house with it.

## How to run

- `go build ./...` type-checks the overview and the subexample at once.
- `go run .` is only an orientation; every topic is its own program.
- Or `make list` lists the topics and `make check` runs the chapter gate.
- Suggested order for the first pass:

| Order | Run it | What it teaches you |
| ----- | ------ | ------------------- |
| 1 | `go run .` | `&` takes an address, `*` reads the value back out |
| 2 | `go run ./examples/01-call-by-value` | Every argument is a copy — and why slices and maps still surprise you |
| 3 | `go run ./examples/02-pointers-mutable-params` | Passing a pointer so the caller sees the change — and why you still cannot swap the caller's pointer out |
| 4 | `go run ./examples/03-pointers-last-resort` | nil, aliasing, and GC work: the reasons to reach for a pointer only on purpose |
| 5 | `go run ./examples/04-map-slice-diffs` | Why a map or a slice reaches the caller's data with no pointer, and when a struct is the better contract |
| 6 | `go run ./examples/05-pointer-passing-perf` | What copying a 1 MB struct actually costs, measured |
| 7 | `go run ./examples/06-slices-as-buffers` | Reusing one byte slice across calls, and the stale-bytes trap |

Topics are numbered folders under `examples/` (`examples/01-`, `examples/02-`, …) in reading order, so `ls examples/` alone
tells a newcomer where to begin. The numbers are the curriculum, not a version.

## Notes

Full chapter notes live in [notes.md](./notes.md).

## Exercises

- `exercises/01-swap-values` — Topics solved: 02-pointers-mutable-params
- `exercises/02-scale-slice` — Topics solved: 04-map-slice-diffs
- `exercises/03-buffer-reuse` — Topics solved: 06-slices-as-buffers
