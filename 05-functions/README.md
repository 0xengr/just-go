# 05 — Functions

The big idea: functions are Go's only unit of behavior, and in this chapter they
wear every hat at once — several return values, named results, a stand-in for
optional parameters, and first-class values that can be passed around and that
remember things (closures). `defer` is the cleanup hook riding along the whole
time. Picture the chapter as the pit wall of a race team: every helper has a job,
hands back a report, and cleans up its station before the next stint.

## How to run

- `go build ./...` type-checks all five subexamples at once.
- `go run .` is only an orientation; every topic is its own program.
- Or use the chapter Makefile: `make run <topic>` (e.g. `make run 01-multiple-returns`); `make list` lists the topics and `make check` runs the gate.
- Suggested order for the first pass:

| Order | Run it | What it teaches you |
| ----- | ------ | ------------------- |
| 1 | `go run ./examples/01-multiple-returns` | Returning `(value, ok)` instead of throwing |
| 2 | `go run ./examples/02-named-return-shadowing` | Named results, blank `return`, and `:=` shadows |
| 3 | `go run ./examples/03-options-variadic` | No named/optional params: struct options + `...T` |
| 4 | `go run ./examples/04-anonymous-closure-func` | Functions as values; closures keep state |
| 5 | `go run ./examples/05-defer-func` | `defer` cleanup, last-in-first-out |

Topics are numbered folders under `examples/` (`examples/01-`, `examples/02-`, …) in reading order, so `ls examples/` alone
tells a newcomer where to begin. The numbers are the curriculum, not a version.

## Notes

Full chapter notes live in [notes.md](./notes.md).

## Exercises

- `exercises/01-safe-divide` — Topics solved: 01-multiple-returns
- `exercises/02-counter-closure` — Topics solved: 04-anonymous-closure-func
- `exercises/03-defer-cleanup` — Topics solved: 05-defer-func
