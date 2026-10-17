# 07 — Types and Methods

The big idea: types are executable documentation. A defined type gives a concept
its own name and identity, and methods let values of that type carry their own
behavior. Go has no inheritance; it favors composition with embedding.

## How to run

- `go build ./...` type-checks the overview and all subexamples at once.
- `go run .` is only an orientation; every topic is its own program.
- Or `make list` lists the topics and `make check` runs the chapter gate.
- Suggested order for the first pass:

| Order | Run it | What it teaches you |
| ----- | ------ | ------------------- |
| 1 | `go run ./01-user-defined-types` | Defined types, aliases, underlying types, and conversions |
| 2 | `go run ./02-methods-and-receivers` | Methods, value vs pointer receivers, and update semantics |
| 3 | `go run ./03-nil-receivers` | Writing methods that handle nil receivers |
| 4 | `go run ./04-embedding-and-promotion` | Composition through embedding and promotion |

Topics are numbered folders (`01-`, `02-`, …) in reading order, so `ls` alone
tells a newcomer where to begin. The numbers are the curriculum, not a version.

## Notes

- A defined type such as `type Score int` has its own type identity, even when
  another defined type shares the same underlying type.
- A type alias such as `type ID = int` is just another name for an existing
  type and gives no extra type safety.
- The underlying type of `type HighScore Score` is `int`, not `Score`.
- Untyped constants can initialize compatible defined types directly; typed values
  need explicit conversion.
- Methods are declared at package level with a receiver. Pointer receivers modify
  the receiver and can handle nil; value receivers work on copies.
- Go automatically takes addresses and dereferences pointers for method calls on
  local variables.
- Embedding promotes fields and methods to the containing struct, but it is not
  inheritance: the containing type cannot be assigned to the embedded type.

## Observe

```text
$ go run ./03-nil-receivers
true
false
```
```