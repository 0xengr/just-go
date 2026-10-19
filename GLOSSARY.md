# Glossary

Plain-English definitions of the Go terms that appear in chapters 1–7. Each
entry names the chapter that introduces it.

- **alias (type alias)** — `type ID = int` makes another name for an existing
  type; no new type identity. (ch 7)
- **array** — fixed-length sequence of values of one type; length is part of the
  type. (ch 3)
- **block** — a region of code delimited by braces `{ }`; variables declared
  inside are scoped to it. (ch 4)
- **capacity** — how many elements a slice's backing array can hold before
  growing. (ch 3)
- **closure** — a function value that captures variables from its enclosing
  scope. (ch 5)
- **composite literal** — `T{...}` syntax to create a value of a composite type.
  (ch 3)
- **const** — a compile-time constant; untyped constants are flexible. (ch 2)
- **defer** — schedules a call to run when the surrounding function returns.
  (ch 5)
- **defined type** — `type Score int` creates a new type with its own identity
  and methods. (ch 7)
- **embedding** — putting one type inside a struct so its fields/methods are
  promoted to the outer type. (ch 7)
- **escape analysis** — compiler decides whether a value lives on stack or heap.
  (ch 6)
- **gofmt** — the canonical Go formatter; the style authority. (ch 1)
- **goroutine** — lightweight concurrent function (not covered here; ch 12).
- **interface** — a set of method signatures; satisfied implicitly (ch 7 quick
  intro; full coverage in book ch 7).
- **len** — built-in returning length of string/slice/array/map/channel.
- **map** — hash table mapping keys to values. (ch 3)
- **method** — a function with a receiver, attached to a type. (ch 7)
- **method set** — the methods callable on a value or pointer of a type. (ch 7)
- **module** — a directory with a `go.mod` naming it; unit of versioning. (ch 1)
- **nil** — zero value for pointers, slices, maps, channels, interfaces, funcs.
- **package** — a directory of Go files sharing a package name. (ch 10 in book)
- **pointer** — `*T` holds the address of a `T`; `&x` takes an address. (ch 6)
- **receiver** — the value a method is called on: value receiver or pointer
  receiver. (ch 7)
- **rune** — an alias for `int32`, represents a Unicode code point. (ch 2)
- **shadowing** — an inner declaration hides an outer one of the same name. (ch 4)
- **slice** — a view (pointer, len, cap) into a backing array. (ch 3)
- **struct** — composite type grouping named fields. (ch 3)
- **underlying type** — the base type a defined type is built on. (ch 7)
- **value receiver / pointer receiver** — method receiver passed by value or by
  pointer. (ch 7)
- **variadic** — a function parameter accepting zero or more trailing args. (ch 5)
- **zero value** — the default value of a type when no initializer is given. (ch 2)