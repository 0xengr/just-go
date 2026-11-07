# 10 — Generics

The big idea: a type parameter is a placeholder for a type, filled in by the
caller and checked by the compiler. It lets one function or one data structure
work with many types without duplicating the body and without falling back to
`any` and type assertions. A constraint says which types are allowed: `any`
means anything, `comparable` means values that support `==` and `!=`, and an
interface with type terms such as `int | float64` lists the exact types a
parameter accepts. The compiler still checks every call, so a mistake is a build
error instead of a runtime surprise.

## How to run

- `go build ./...` type-checks the overview, the examples, and the exercises at once.
- Every topic is its own program; `main.go` at the chapter root is the overview.
- Or `make list` lists the topics and `make check` runs the chapter gate.
- Suggested order for the first pass:

| Order | Run it | What it teaches you |
| ----- | ------ | ------------------- |
| 1 | `go run .` | The duplication problem and the first generic function |
| 2 | `go run ./examples/01-type-constraints` | Type parameters, constraints, and a `Number` interface |
| 3 | `go run ./examples/02-type-inference` | When the compiler infers types and when you must state them |
| 4 | `go run ./examples/03-type-generic-params-list` | `K comparable, V any` with maps of products and users |
| 5 | `go run ./examples/04-generic-types` | A generic `Box[T]` with methods and a useful zero value |
| 6 | `go run ./examples/05-generic-functions` | `All` and `Any`, generic functions that take a check function |
| 7 | `go run ./examples/06-type-terms` | Type terms, `~` underlying types, and operators |
| 8 | `go run ./examples/07-interface-constraints` | Using an interface such as `Describer` as a constraint |
| 9 | `go run ./examples/08-generic-interfaces` | An interface with its own type parameter |
| 10 | `go run ./examples/09-comparable` | The `comparable` and interface-value pitfall |
| 11 | `go run ./examples/10-standard-library` | `slices`, `maps`, and `cmp` generic helpers |
| 12 | `go run ./exercises/01-last-value` | Exercise: last element of any slice |
| 13 | `go run ./exercises/02-map-keys` | Exercise: collect map keys with two type parameters |
| 14 | `go run ./exercises/03-generic-queue` | Exercise: a generic FIFO queue |
| 15 | `go run ./exercises/04-unique-values` | Exercise: remove duplicates from a slice |

A few examples keep a failing line commented out so the program still builds.
Those lines are marked `BUG:` and the comment quotes the real compiler error.
Read them; deleting the `//` reproduces the failure.

`examples/09-comparable` has one more commented block that compiles and then
panics at run time, because an interface value holds a slice. That is the
lesson: `comparable` checks the type of the interface, not what is inside it.

Topics are numbered folders (`01-`, `02-`, …) in reading order, so `ls` alone
tells a newcomer where to begin. The numbers are the curriculum, not a version.

## Cheat sheet

| Idea | Syntax |
| ---- | ------ |
| Any type | `func Print[T any](v T)` |
| Comparable type | `func Equal[T comparable](a, b T) bool` |
| A constraint interface | `type Number interface { int \| int64 \| float64 }` |
| Use the constraint | `func Sum[T Number](values []T) T` |
| Underlying types | `type Integer interface { ~int \| ~int64 }` |
| Union with methods | `type Printable interface { fmt.Stringer; ~int \| ~float64 }` |
| Generic type | `type Box[T any] struct { value T }` |
| Methods on a generic type | `func (b *Box[T]) Put(v T)` |
| Zero value of `T` | `var zero T` |
| Two type parameters | `func Keys[K comparable, V any](m map[K]V) []K` |
| Explicit instantiation | `Keys[string, int](m)` |
| Inferred instantiation | `Keys(m)` |
| Generic function as a value | `f := Sum[int]` |
| Ordering helper | `cmp.Compare(a, b)` |
| Sort a slice | `slices.Sort(values)` |
| Sort with a function | `slices.SortFunc(values, cmp.Compare)` |
| Copy a map | `maps.Clone(m)` |

## Notes

Full chapter notes live in [notes.md](./notes.md).

## Exercises

- `exercises/01-last-value` — Topics solved: a generic function over any slice,
  the `(T, bool)` result, and the zero value of `T`.
- `exercises/02-map-keys` — Topics solved: two type parameters, `K comparable`,
  `V any`, and returning `[]K` instead of `[]any`.
- `exercises/03-generic-queue` — Topics solved: a generic type with several
  methods, a pointer receiver, FIFO order, and an empty-queue result.
- `exercises/04-unique-values` — Topics solved: `comparable`, using a value as
  a map key, and building a result slice without duplicates.
