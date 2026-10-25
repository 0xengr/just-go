# 09 — Errors

The big idea: in Go an error is an ordinary value, not an exception. A function
that can fail returns an `error` next to its result, and the caller checks it
with `if err != nil`. Nothing is thrown and nothing is caught. The `error`
interface is tiny — any type with `Error() string` satisfies it — so packages
report failures without agreeing on a concrete type. Panic is the escape hatch
for unrecoverable problems, and `defer` is the code that still runs while a
panic unwinds.

## How to run

- `go build ./...` type-checks the overview, the examples, and the exercises at once.
- Every topic is its own program; `main.go` at the chapter root is the overview.
- Or `make list` lists the topics and `make check` runs the chapter gate.
- Suggested order for the first pass:

| Order | Run it | What it teaches you |
| ----- | ------ | ------------------- |
| 1 | `go run .` | Returning a result and an error, and checking it against `nil` |
| 2 | `go run ./examples/01-error-kinds` | Syntax, runtime, and semantic errors |
| 3 | `go run ./examples/02-error-interface` | The `error` interface and implicit satisfaction |
| 4 | `go run ./examples/03-creating-error-values` | `errors.New` and package-level sentinel errors |
| 5 | `go run ./examples/04-panic` | Panic from the developer and from the runtime |
| 6 | `go run ./examples/05-defer-with-panic` | Deferred cleanup while a panic unwinds |
| 7 | `go run ./exercises/01-calculate-pay` | Exercise: returning sentinel errors |
| 8 | `go run ./exercises/02-calculate-pay-panic` | Exercise: panicking with the same errors |

`examples/05-defer-with-panic` ends in a panic on purpose. Its non-zero exit
status and stack trace are the lesson; the program is not broken.

`examples/04-panic` prints nothing as committed, because its panic calls are
commented out to keep the chapter gate green. Uncomment one to see the crash
(add `"fmt"` to the imports for the runtime loop).

Topics are numbered folders (`01-`, `02-`, …) in reading order, so `ls` alone
tells a newcomer where to begin. The numbers are the curriculum, not a version.

## Cheat sheet

| Idea | Syntax |
| ---- | ------ |
| Create an error | `errors.New("denominator is 0")` |
| Return it with a result | `return 0, 0, errors.New("...")` |
| Check it | `if err != nil { ... }` |
| The interface | `type error interface { Error() string }` |
| Custom error type | `func (e *MyError) Error() string { ... }` |
| Compile-time check | `var _ error = (*MyError)(nil)` |
| Print an error | `fmt.Println(err)` — calls `Error()` |
| Stop the program | `panic("missing PORT")` |
| Clean up on panic | `defer file.Close()` |
| Exit with a status | `os.Exit(1)` |

## Notes

Full chapter notes live in [notes.md](./notes.md).

## Exercises

- `exercises/01-calculate-pay` — Topics solved: `errors.New` sentinels,
  validation, returning and checking errors.
- `exercises/02-calculate-pay-panic` — Topics solved: `panic` with error values,
  deferred reporting while a panic unwinds.
