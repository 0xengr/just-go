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
| 1 | `go run ./01-zero-values` | Every type has a useful zero value, and `append` works on a nil slice |
| 2 | `go run ./02-var-short` | `var` everywhere, `:=` only inside a function |
| 3 | `go run ./03-consts` | Compile-time constants, typed and untyped |
| 4 | `go run ./04-numeric` | Integer division truncates; underscores in literals |
| 5 | `go run ./05-strings` | `len` counts bytes, `range` counts runes |
| 6 | `go run ./06-conversions` | No implicit conversion anywhere |

Topics are numbered folders (`01-`, `02-`, …) in reading order, so `ls` alone
tells a newcomer where to begin. The numbers are the curriculum, not a version.
`01-zero-values` and `06-conversions` are the pair worth doing first and last:
between them they cover the two rules that trip up everything in the chapters
after this one.

## 1. Every type has a useful zero value (`01-zero-values/`)

Plain English: declaring a variable with `var` and no initial value does not
leave it uninitialized or garbage-filled. Go sets it to the type's zero value:
`""` for strings, `0` for numbers, `false` for booleans, `nil` for pointers,
slices, maps, channels, interfaces, and functions. A struct's zero value is the
struct with every field set to its own zero value. Design types so that this
zero value is already correct and no constructor is needed.

Real world: a scorecard printed before the game starts, with every box showing
zero. The game can begin immediately because the "empty" state is a legal state.
A nil slice is the same idea: an empty list that is ready to be appended to,
with no allocation yet.

```go
package main

import "fmt"

type player struct {
	name  string
	score int
	ready bool
}

func main() {
	var p player
	fmt.Printf("%q %d %t\n", p.name, p.score, p.ready)

	var inbox []string
	fmt.Println("nil slice len:", len(inbox), "is nil:", inbox == nil)
	inbox = append(inbox, "welcome") // append works on nil slices
	fmt.Println(inbox)
}
```

Output:

```text
"" 0 false
nil slice len: 0 is nil: true
[welcome]
```

Beginner trap: assuming a nil slice cannot be appended to. It can, and `append`
allocates the backing array on the first call. The matching trap is the opposite
assumption for maps: a nil map is also a zero value, but writing to it panics.
`m[k] = v` on a nil map crashes at runtime, while reading and ranging over a
nil map are perfectly fine. Slices get the friendly treatment; maps do not.

## 2. Declaring variables: `var` and `:=` (`02-var-short/`)

Plain English: `var` declares a variable and works everywhere, including at
package level and inside functions. `:=` is shorthand that declares and
initializes in one step, but only inside a function, and the type is inferred
from the right-hand side. Once a variable exists, reassign it with plain `=`,
never `:=`. A package-level variable cannot use `:=` at all.

Real world: `var` is filling out a form where you write the type and the value.
`:=` is a sticky note you slap on a value and let the compiler read the label
off of it. Both end up as a named variable; only `var` is allowed at the top of
the file.

```go
package main

import "fmt"

// WARNING: := is not legal up here, at package level.
var serverPort = 8080

func main() {
	laps := 3               // short decl, type inferred as int
	var best float64 = 58.4 // explicit type when it matters
	fmt.Println("laps:", laps, "best:", best, "port:", serverPort)

	laps = 4 // plain = for variables that already exist
	fmt.Println("updated laps:", laps)
}
```

Output:

```text
laps: 3 best: 58.4 port: 8080
updated laps: 4
```

Beginner trap: using `:=` when you meant `=`. `laps := 4` inside `main` does
not update the existing `laps`; it declares a brand-new variable that shadows it
for the rest of the block, and the original stays `3`. The compiler only allows
`:=` to reuse a name when at least one other name on the left is genuinely new,
so this trap is easy to hit and hard to spot. When in doubt, write `=`.

## 3. Constants: typed and untyped (`03-consts/`)

Plain English: a constant is a compile-time value, not a variable. Go has two
kinds. An untyped constant such as `const limit = 90_000` has no fixed type; it
takes whatever type the context asks for, so the same constant can initialize
both an `int` and a `float64`. A typed constant such as
`const trackName string = "Harbor Loop"` is locked to that type forever.
Constants are never addresses: you cannot take `&limit`. Unused constants
compile fine, while unused local variables do not.

Real world: an untyped constant is a blank check that becomes whatever currency
the transaction needs. A typed constant is a gift card locked to one store. Both
are fixed values; only the untyped one is flexible about where it can be spent.

```go
package main

import "fmt"

const (
	qualifyingMs = 90_000 // untyped: fits int, float64, etc.
	raceLaps     = 5
)

const trackName string = "Harbor Loop" // typed: always a string

func main() {
	var cutoff int = qualifyingMs
	var cutoffF float64 = qualifyingMs // same constant, both work (untyped)
	fmt.Println(trackName, "laps:", raceLaps, "cutoff:", cutoff, cutoffF)
}
```

Output:

```text
Harbor Loop laps: 5 cutoff: 90000 90000
```

Beginner trap: assuming an untyped constant can hold any value. It is flexible
about *type*, not about *size*: `const big = 1000` used as a `byte` is a
compile error because 1000 does not fit in a byte. The flexibility only applies
when the value actually fits the target type. A second trap: trying to reassign
a constant with `=`. Constants are not variables, so `qualifyingMs = 100_000`
does not compile.

## 4. Numeric types and integer division (`04-numeric/`)

Plain English: dividing two integers truncates toward zero, so `7 / 10` is `0`,
not `0.7`. To get a fractional result, convert to `float64` *before* dividing.
Numeric literals such as `5` and `58.4` are untyped: they carry no type of
their own and take the one the surrounding code needs. Underscores in literals
(`90_000`, `1_000_000`) are visual separators and are ignored by the compiler.

Real world: splitting 7 cookies among 10 people with integer math gives 0
cookies each and 7 left over. You have to switch to a measuring cup (floats)
before dividing to get `0.7`. The underscores are just thousands separators
written into the number itself.

```go
package main

import "fmt"

func main() {
	finished, total := 7, 10
	// WARNING: int division drops the fraction, so 7/10 is 0, not 0.7.
	fmt.Println("truncated:", finished/total)

	rate := float64(finished) / float64(total) // convert one side first
	fmt.Printf("completion: %.1f%%\n", rate*100)

	var laps int = 5       // untyped literal 5 fits int
	var avg float64 = 58.4 // untyped literal fits float64
	fmt.Println("laps:", laps, "avg:", avg)
}
```

Output:

```text
truncated: 0
completion: 70.0%
laps: 5 avg: 58.4
```

Beginner trap: converting *after* dividing. `float64(finished / total)` converts
the already-truncated `0` into `0.0`, which is too late. Convert at least one
operand before the `/`. A second trap: `int(x)` truncates toward zero, it does
not round, so `int(21.7)` is `21` and `int(-21.7)` is `-21`. Use
`math.Round` when rounding is what you mean.

## 5. Strings are bytes; letters are runes (`05-strings/`)

Plain English: a Go string is an immutable sequence of bytes, stored as UTF-8.
`len(s)` counts *bytes*, not letters, so a multi-byte character such as `é`
reports 2. A rune is a Unicode code point and is just an `int32`. `range` over
a string decodes one rune at a time and hands back its byte offset, and
`utf8.RuneCountInString` counts letters. Indexing `s[i]` yields a single byte,
which may be only part of a character. Single quotes denote a rune (`'é'`),
double quotes denote a string (`"café"`).

Real world: a string is a sealed row of byte-sized boxes. A letter like `é`
occupies two boxes. Counting boxes gives the byte length; counting decoded
letters gives the rune count. `range` is the tool that opens the boxes one
letter at a time, and it tells you where each letter started.

```go
package main

import (
	"fmt"
	"unicode/utf8"
)

func main() {
	tag := "café"
	fmt.Println("bytes:", len(tag), "letters:", utf8.RuneCountInString(tag))

	// WARNING: the index is a byte offset, so it jumps by 2 for "é",
	// while r is a rune (int32) holding the decoded character.
	for i, r := range tag {
		fmt.Printf("byte %d -> %c\n", i, r)
	}
}
```

Output:

```text
bytes: 5 letters: 4
byte 0 -> c
byte 1 -> a
byte 2 -> f
byte 3 -> é
```

Beginner trap: trusting `len` for "how many characters." For `"café"` it reports
5, not 4. The matching trap is indexing into the middle of a multi-byte
character: `tag[3]` is a byte, not the rune `é`, and slicing at a byte boundary
that lands inside a character produces invalid UTF-8. Always slice strings at
rune boundaries, or convert to `[]rune` first when you need per-character
indexing.

## 6. Conversions are always explicit (`06-conversions/`)

Plain English: Go never converts a value for you, not even between `int` and
`int64` or between `int` and `float64`. Write the conversion as `T(value)`. A
few conversions matter constantly: `int(celsius)` truncates a float toward
zero, `[]byte(s)` copies a string's bytes into a mutable slice, and
`string(raw)` copies bytes back into a string. The one that bites is `int` to
string: a cast produces the rune with that code point, not the decimal digits,
so `42` becomes `"*"` (U+002A). Use `strconv.Itoa` for decimal text.

Real world: currency exchange with no automatic rate. You state the conversion
explicitly every time, and the compiler checks that both sides agree on the
target type. A cast is relabeling a box: the bytes do not change, only the type
the compiler attaches to them.

```go
package main

import (
	"fmt"
	"strconv"
)

func main() {
	celsius := 21.7
	fmt.Println("truncated to int:", int(celsius)) // drops .7, no rounding

	bib := 42
	fmt.Println("bib number:", strconv.Itoa(bib)) // int -> decimal string

	raw := []byte("pit")
	fmt.Println("bytes:", raw, "back to string:", string(raw))
}
```

Output:

```text
truncated to int: 21
bib number: 42
bytes: [112 105 116] back to string: pit
```

Beginner trap: `string(42)` and expecting `"42"`. It produces a one-character
string holding the rune U+002A. The fix is `strconv.Itoa(42)` (or
`strconv.FormatInt` for other bases and widths). A second trap: forgetting that
`[]byte(s)` and `string(b)` *copy*. Mutating the byte slice does not change the
original string, because strings are immutable; the slice is a separate array.

## Rules I want to remember

- Every variable starts at its type's zero value; design types so that zero is
  already correct.
- `append` works on a nil slice and allocates on first use; a nil map panics on
  write.
- `var` works everywhere; `:=` works only inside functions and infers the type.
- Reassign an existing variable with `=`, never `:=`.
- Untyped constants adapt to their context; typed constants are locked to one
  type. Unused constants are allowed; unused local variables are not.
- Integer division truncates toward zero. Convert to `float64` before dividing.
- `int(x)` truncates toward zero; it does not round.
- `len(string)` counts bytes; `utf8.RuneCountInString` and `range` count runes.
- Single quotes are runes (`'é'`), double quotes are strings (`"café"`).
- No implicit conversion anywhere: write `T(value)` every time.
- `int` to decimal string needs `strconv.Itoa`, not a cast: `string(42)` is a
  rune, not `"42"`.
- `[]byte(s)` and `string(b)` copy; strings are immutable.

## Exercises

### Exercise 1 — Zero values and a nil slice

Write a small program that declares `var s scoreboard` where

```go
type scoreboard struct {
	home, away int
	overtime   bool
}
```

prints it with `fmt.Printf("%d-%d (ot=%t)\n", ...)`, then declares a nil
`[]string` named `events`, prints `events == nil`, appends `"kickoff"` and
`"goal"` to it, and prints the slice.

Hint: `var` gives every field its zero value, and `append` allocates the backing
array on the first call even when the slice starts out nil.

<details>
<summary>Solution</summary>

```go
package main

import "fmt"

type scoreboard struct {
	home, away int
	overtime   bool
}

func main() {
	var s scoreboard
	fmt.Printf("%d-%d (ot=%t)\n", s.home, s.away, s.overtime)

	var events []string
	fmt.Println("nil:", events == nil)
	events = append(events, "kickoff", "goal")
	fmt.Println(events)
}
```

Expected output:

```text
0-0 (ot=false)
nil: true
[kickoff goal]
```

</details>

### Exercise 2 — Fix the declaration errors

This program does not compile. Fix it so it prints `4 58.4 8080` without
changing what it computes.

```go
package main

import "fmt"

port := 8080

func main() {
	laps := 3
	laps := 4
	var best float64 = 58.4
	fmt.Println(laps, best, port)
}
```

Hint: `:=` is illegal at package level, and redeclaring an existing variable in
the same scope needs `=` instead of `:=`.

<details>
<summary>Solution</summary>

```go
package main

import "fmt"

var port = 8080

func main() {
	laps := 3
	laps = 4
	var best float64 = 58.4
	fmt.Println(laps, best, port)
}
```

Expected output:

```text
4 58.4 8080
```

</details>

### Exercise 3 — Untyped vs typed constants

Write a program with an untyped constant `limit = 90_000` and a typed constant
`trackName string = "Harbor Loop"`. In `main`, store `limit` in an `int` and a
`float64`, print both, then print `trackName`. Predict the output.

Hint: the untyped constant adapts to each target type; the typed one is already
a string.

<details>
<summary>Solution</summary>

```go
package main

import "fmt"

const (
	limit    = 90_000
	trackName string = "Harbor Loop"
)

func main() {
	var cutoff int = limit
	var cutoffF float64 = limit
	fmt.Println(trackName, cutoff, cutoffF)
}
```

Expected output:

```text
Harbor Loop 90000 90000
```

</details>

### Exercise 4 — Average without truncation

Given `finished := 7` and `total := 10`, print the completion percentage with one
decimal place (`70.0%`). Use an underscore in at least one numeric literal.

Hint: convert to `float64` before dividing; `%.1f%%` formats one decimal place.

<details>
<summary>Solution</summary>

```go
package main

import "fmt"

func main() {
	finished, total := 7, 10
	rate := float64(finished) / float64(total)
	fmt.Printf("completion: %.1f%%\n", rate*100)

	const scale = 1_000
	fmt.Println("scaled:", int(rate*float64(scale)))
}
```

Expected output:

```text
completion: 70.0%
scaled: 700
```

</details>

### Exercise 5 — Bytes, runes, and the `int`-to-string trap

For `tag := "café"`, print the byte length, the rune count, and each rune's
byte offset. Then convert `"pit"` to `[]byte`, print it, convert it back to a
string, and print that. Finally show that `string(42)` is not `"42"`.

Hint: use `len`, `utf8.RuneCountInString`, and `range`; use `strconv.Itoa` for
the decimal text.

<details>
<summary>Solution</summary>

```go
package main

import (
	"fmt"
	"strconv"
	"unicode/utf8"
)

func main() {
	tag := "café"
	fmt.Println("bytes:", len(tag), "letters:", utf8.RuneCountInString(tag))
	for i, r := range tag {
		fmt.Printf("byte %d -> %c\n", i, r)
	}

	raw := []byte("pit")
	fmt.Println("bytes:", raw, "back to string:", string(raw))

	fmt.Println("not forty-two:", string(42))
	fmt.Println("forty-two:", strconv.Itoa(42))
}
```

Expected output:

```text
bytes: 5 letters: 4
byte 0 -> c
byte 1 -> a
byte 2 -> f
byte 3 -> é
bytes: [112 105 116] back to string: pit
not forty-two: *
forty-two: 42
```

</details>
