# 02 — Predeclared Types and Declarations

The big idea: Go's predeclared types (`int`, `float64`, `string`, `bool`, `rune`,
`byte`, and friends) are the vocabulary every program is written in, and the
rules for binding names to values are deliberately strict. There is no implicit
conversion between types, every variable starts at a known and useful zero value,
and `:=` is a function-local convenience rather than a second way to declare
things. Learn these rules once and most of the confusing compile errors in the
chapters that follow stop being confusing.

## How to run

- `go build ./...` type-checks the overview and all subexamples at once.
- `go run .` is only an orientation; every topic is its own program.
- Or `make list` lists the topics and `make check` runs the chapter gate.
- Suggested order for the first pass:

| Order | Run it | What it teaches you |
| ----- | ------ | ------------------- |
| 1 | `go run ./examples/01-zero-values` | Every type has a useful zero value, and `append` works on a nil slice |
| 2 | `go run ./examples/02-var-short` | `var` everywhere, `:=` only inside a function |
| 3 | `go run ./examples/03-consts` | Compile-time constants, typed and untyped |
| 4 | `go run ./examples/04-numeric` | Integer division truncates; underscores in literals |
| 5 | `go run ./examples/05-strings` | `len` counts bytes, `range` counts runes |
| 6 | `go run ./examples/06-conversions` | No implicit conversion anywhere |

Topics are numbered folders under `examples/` (`examples/01-`, `examples/02-`, …) in reading order, so `ls examples/` alone
tells a newcomer where to begin. The numbers are the curriculum, not a version.
`01-zero-values` and `06-conversions` are the pair worth doing first and last:
between them they cover the two rules that trip up everything in the chapters
after this one.

## Notes

Full chapter notes live in [notes.md](./notes.md).

## Exercises

- `exercises/01-lap-stats` — Topics solved: 04-numeric, 06-conversions
- `exercises/02-callsign-bytes` — Topics solved: 05-strings
- `exercises/03-const-limits` — Topics solved: 03-consts, 04-numeric
