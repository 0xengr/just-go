# 04 — Blocks, Shadows, and Control Structures

The big idea: Go has only three control keywords — `if`, `for`, `switch`
(plus a rarely-used `goto`). Everything else in this chapter is about
*where* your variables live (blocks) and *how* you accidentally hide them
(shadowing). Think of it as learning the rules of the race control tower:
who can see which screen, and which button actually stops which race.

## How to run

- `go build ./...` type-checks all nine subexamples at once.
- `go run .` is only an orientation; every topic is its own program.
- Or `make run 01-blocks`, `make run 06-for-range`, … — the chapter Makefile also gives `make list` and `make check`.
- Suggested order for the first pass:

| Order | Run it | What it teaches you |
| ----- | ------ | ------------------- |
| 1 | `go run ./examples/01-blocks` | Blocks nest; inner sees outer, not the reverse |
| 2 | `go run ./examples/02-shadowing` | `:=` can silently create a lookalike variable |
| 3 | `go run ./examples/03-if-scope` | `if x := ...; cond` keeps temporaries local |
| 4 | `go run ./examples/04-for-complete` | The classic `for i := 0; i < n; i++` and its variations |
| 5 | `go run ./examples/05-for-while` | The `while` shape, infinite loops, `break`/`continue` |
| 6 | `go run ./examples/06-for-range` | The everyday loop over slices, maps, and strings |
| 7 | `go run ./examples/07-for-labels` | Steering an *outer* loop from an inner one |
| 8 | `go run ./examples/08-switch-cases` | Matching values, blank switches, no fallthrough |
| 9 | `go run ./examples/09-goto` | The one rare case where jumping forward reads better |

Topics are numbered folders under `examples/` (`examples/01-`, `examples/02-`, …) in reading order, so `ls examples/` alone
tells a newcomer where to begin. The numbers are the curriculum, not a version.
Here the numbering simply encodes the table above.

## Notes

Full chapter notes live in [notes.md](./notes.md).

## Exercises

- `exercises/01-fizzbuzz-switch` — Topics solved: 08-switch-cases
- `exercises/02-shadow-fix` — Topics solved: 02-shadowing
- `exercises/03-rune-walk` — Topics solved: 06-for-range
