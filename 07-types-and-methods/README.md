# 07 — Types and Methods

The big idea: types are executable documentation. A defined type gives a concept
its own name and identity, and methods let values of that type carry their own
behavior. Go has no inheritance; it favors composition with embedding.

## How to run

- `go build ./...` type-checks all four subexamples at once.
- Every topic is its own program; this chapter has no overview `main.go`.
- Or `make list` lists the topics and `make check` runs the chapter gate.
- Suggested order for the first pass:

| Order | Run it | What it teaches you |
| ----- | ------ | ------------------- |
| 1 | `go run ./examples/01-user-defined-types` | Defined types, aliases, underlying types, and conversions |
| 2 | `go run ./examples/02-methods-and-receivers` | Methods, value vs pointer receivers, and update semantics |
| 3 | `go run ./examples/03-nil-receivers` | Writing methods that handle nil receivers |
| 4 | `go run ./examples/04-embedding-and-promotion` | Composition through embedding and promotion |

Topics are numbered folders under `examples/` (`examples/01-`, `examples/02-`, …) in reading order, so `ls examples/` alone
tells a newcomer where to begin. The numbers are the curriculum, not a version.

## Notes

Full chapter notes live in [notes.md](./notes.md).

## Exercises

- `exercises/01-score-type` — Topics solved: 01-user-defined-types, 02-methods-and-receivers
- `exercises/02-counter-pointer` — Topics solved: 02-methods-and-receivers
- `exercises/03-tree-insert` — Topics solved: 03-nil-receivers
