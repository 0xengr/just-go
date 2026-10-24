# 08 — Interfaces

The big idea: an interface describes behavior, not data. It is a list of method
signatures, and any type that has those methods satisfies it — there is no
`implements` keyword and nothing to register. That is duck typing, and it is how
Go gets polymorphism without inheritance.

## How to run

- `go build ./...` type-checks the overview, the examples, and the exercise at once.
- Every topic is its own program; `main.go` at the chapter root is the overview.
- Or `make list` lists the topics and `make check` runs the chapter gate.
- Suggested order for the first pass:

| Order | Run it | What it teaches you |
| ----- | ------ | ------------------- |
| 1 | `go run .` | Interfaces as a set of methods, interface parameters, variadic interfaces |
| 2 | `go run ./examples/duck-typing` | Implicit satisfaction with a single method |
| 3 | `go run ./examples/polymorphism` | One interface-typed function replacing three concrete ones |
| 4 | `go run ./examples/interface-api-design` | Accepting `io.Reader` instead of a string, and returning a concrete type |
| 5 | `go run ./examples/type-assertion-switches` | Recovering concrete event types with a type switch |
| 6 | `go run ./exercises/notification-system` | Exercise: switch dispatch vs interface dispatch, plus `error` returns |

The topic folders are not numbered; this table is the reading order.

## Cheat sheet

| Idea | Syntax |
| ---- | ------ |
| Declare an interface | `type Shape interface { Area() float64 }` |
| Satisfy an interface | `func (c Circle) Area() float64 { ... }` — no keyword |
| Take any implementation | `func printArea(s Shape)` |
| Take many | `func sayAll(say ...speaker)` |
| Empty interface / any value | `any`, `interface{}` |
| Assert the dynamic type | `v, ok := s.(Circle)` |
| Branch on the dynamic type | `switch v := s.(type) { ... }` |
| Errors are interfaces | `type error interface { Error() string }` |
| Nil interface | `var s Shape; s == nil` is true |

## Notes

Full chapter notes live in [notes.md](./notes.md).

## Exercises

- `exercises/notification-system` — Topics solved: interface declaration and
  satisfaction, interface parameters, `error` returns, switch vs interface
  dispatch.
