# 03 — Composite Types

The big idea: Go gives you four ways to group data — arrays, slices, maps,
and structs — and they are all **values**. Assigning one copies it; nothing is
shared unless you explicitly take an address or pass a slice header. Slices
are the workhorse: a tiny header (pointer, length, capacity) sitting on top
of a backing array. Think of an array as an egg carton with a fixed number
of slots, a slice as a sticky note saying "look at slots 2 through 5 of that
carton", a map as a drawer label that still opens the same drawer, and a
struct as a photocopy of a form.

## How to run

- `go build ./...` type-checks the overview and all eleven subexamples at once.
- `go run .` is only an orientation; every topic is its own program.
- Or use the chapter Makefile: `make run <topic>` (e.g. `make run 02-slices`); `make list` lists the topics and `make check` runs the gate.
- Suggested order for the first pass:

| Order | Run it | What it teaches you |
| ----- | ------ | ------------------- |
| 1 | `go run ./examples/01-arrays` | Fixed length is part of the type; assignment copies everything |
| 2 | `go run ./examples/02-slices` | `len` is what you may index, `cap` is the room left |
| 3 | `go run ./examples/03-slice-aliasing` | A slice header is a pointer, so every alias shares data |
| 4 | `go run ./examples/04-slicing` | `a[i:j]` versus `a[i:j:k]`, and what bounds panic |
| 5 | `go run ./examples/05-slice-emptying` | `len 0`, `nil`, and `[:0]` are three different things |
| 6 | `go run ./examples/06-maps` | nil maps read but panic on write; order is randomized |
| 7 | `go run ./examples/07-map-sets` | `map[T]struct{}` occupies no memory |
| 8 | `go run ./examples/08-map-comparing` | Maps are not comparable; use `maps.Equal` |
| 9 | `go run ./examples/09-structs` | Structs are values: assignment copies, `==` is structural |
| 10 | `go run ./examples/10-strings-bytes` | Converting to `[]byte`/`[]rune` always copies |
| 11 | `go run ./examples/11-array-conversion` | Converting an array to a slice allocates and copies |

Topics are numbered folders under `examples/` (`examples/01-`, `examples/02-`, …) in reading order, so `ls examples/` alone
tells a newcomer where to begin. The numbers are the curriculum, not a version.

## Notes

Full chapter notes live in [notes.md](./notes.md).

## Exercises

- `exercises/01-slice-stats` — Topics solved: 02-slices, 04-slicing, 05-slice-emptying
- `exercises/02-word-frequency` — Topics solved: 06-maps
- `exercises/03-points` — Topics solved: 09-structs
