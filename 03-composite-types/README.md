# 03 — Composite Types

The big idea: Go gives you four ways to group data — arrays, slices, maps,
and structs — and they are all **values**. Assigning one copies it; nothing is
shared unless you explicitly take an address or pass a slice header. Slices
are the workhorse: a tiny header (pointer, length, capacity) sitting on top
of a backing array. Think of an array as an egg carton with a fixed number
of slots, a slice as a sticky note saying "look at slots 2 through 5 of that
carton", a map as a drawer label that still opens the same drawer, and a
struct as a photocopy of a form.

## How to run

- `go build ./...` type-checks the overview and all eleven subexamples at once.
- `go run .` is only an orientation; every topic is its own program.
- Or use the chapter Makefile: `make run <topic>` (e.g. `make run 02-slices`); `make list` lists the topics and `make check` runs the gate.
- Suggested order for the first pass:

| Order | Run it | What it teaches you |
| ----- | ------ | ------------------- |
| 1 | `go run ./01-arrays` | Fixed length is part of the type; assignment copies everything |
| 2 | `go run ./02-slices` | `len` is what you may index, `cap` is the room left |
| 3 | `go run ./03-slice-aliasing` | A slice header is a pointer, so every alias shares data |
| 4 | `go run ./04-slicing` | `a[i:j]` versus `a[i:j:k]`, and what bounds panic |
| 5 | `go run ./05-slice-emptying` | `len 0`, `nil`, and `[:0]` are three different things |
| 6 | `go run ./06-maps` | nil maps read but panic on write; order is randomized |
| 7 | `go run ./07-map-sets` | `map[T]struct{}` occupies no memory |
| 8 | `go run ./08-map-comparing` | Maps are not comparable; use `maps.Equal` |
| 9 | `go run ./09-structs` | Structs are values: assignment copies, `==` is structural |
| 10 | `go run ./10-strings-bytes` | Converting to `[]byte`/`[]rune` always copies |
| 11 | `go run ./11-array-conversion` | Converting an array to a slice allocates and copies |

Topics are numbered folders (`01-`, `02-`, …) in reading order, so `ls` alone
tells a newcomer where to begin. The numbers are the curriculum, not a version.

## 1. Arrays: the length is part of the type (`01-arrays/`)

Plain English: an array has a fixed size written into its type — `[3]int` and
`[4]int` are different types and do not convert freely. Declaring one gives
you that many zero values, and assigning one array to another copies every
element. Arrays show up as fixed buffers and in `[N]byte` style APIs; for
anything dynamic you reach for a slice.

Real world: an egg carton with exactly twelve slots. You cannot hand someone
a "carton" and let them decide the size later, and giving someone your carton
means copying all twelve eggs into a new carton — the original stays on your
shelf.

```go
var grid [3][2]int // 3 rows of 2 columns
grid[1][0] = 7
fmt.Println("grid:", grid, "rows:", len(grid), "cols:", len(grid[0]))

scores := [5]int{9, 8, 7, 6, 5}
backup := scores // full element copy, not a reference
backup[0] = 0
fmt.Println("scores:", scores, "backup:", backup, "sum:", total(scores))
```

`total` takes the array by value, so the whole five-element array is copied
onto the stack:

```go
func total(vals [5]int) int {
	var sum int
	for _, v := range vals {
		sum += v
	}
	return sum
}
```

Output:

```text
grid: [[0 0] [7 0] [0 0]] rows: 3 cols: 2
scores: [9 8 7 6 5] backup: [0 8 7 6 5] sum: 35
```

Beginner trap: `backup := scores` looks like aliasing, but it is a full copy —
mutating `backup[0]` leaves `scores[0]` at 9, which is the whole point of
the printout. The second trap is size: `total([4]int{...})` does not
compile, because `[4]int` is a different type from `[5]int`. If the size
needs to vary, the parameter should be `[]int` and the caller slices the
array (`total(scores[:])`).

## 2. Slices: a window, not a container (`02-slices/`)

Plain English: a slice is a small header — pointer, length, capacity — that
points at a backing array. `var s []T` is a nil slice, `[]T{}` is an empty
but non-nil slice, and `make([]T, len, cap)` preallocates room. Only `make`
sets the capacity; only `append` grows it. `len` is what you may index,
`cap` is how much room is left before `append` has to allocate.

Real world: a bookmark with a start page and an end page, pointing at a book
on the shelf. The bookmark is not the book. Two bookmarks can point at the
same book and see each other's scribbles; replacing your bookmark with a new
one does not change the book or anyone else's bookmark.

```go
var nilScores []int // nil slice: len 0, cap 0, prints []
fmt.Printf("nil:     len=%d cap=%d %v\n", len(nilScores), cap(nilScores), nilScores)

empty := []int{} // non-nil but still empty
fmt.Printf("empty:   len=%d cap=%d %v\n", len(empty), cap(empty), empty)

literal := []int{3, 1, 2} // len == cap == 3
fmt.Printf("literal: len=%d cap=%d %v\n", len(literal), cap(literal), literal)

made := make([]int, 2, 4) // len 2 zeroed, cap 4
made[1] = 5
fmt.Printf("make:    len=%d cap=%d %v\n", len(made), cap(made), made)
```

`append` writes into spare capacity, or allocates a bigger array and copies:

```go
grown := make([]int, 0, 2)
for i := range 6 {
	grown = append(grown, i)
	fmt.Printf("append %d -> len=%d cap=%d\n", i, len(grown), cap(grown))
}
```

Output:

```text
nil:     len=0 cap=0 []
empty:   len=0 cap=0 []
literal: len=3 cap=3 [3 1 2]
make:    len=2 cap=4 [0 5]
append 0 -> len=1 cap=2
append 1 -> len=2 cap=2
append 2 -> len=3 cap=4
append 3 -> len=4 cap=4
append 4 -> len=5 cap=8
append 5 -> len=6 cap=8
```

Beginner trap: a nil slice and an empty slice both print `[]` and have
length 0, but they are not the same value — `nilScores == nil` is true and
`empty == nil` is false. That distinction matters for JSON encoding and for
"has this been initialized?" checks. The second trap is capacity: growth
doubles (2 → 4 → 8) in this run, but that is an implementation detail —
never write code that depends on the exact cap after an append, and never
assume `append` reused your array.

## 3. Slice aliasing: one backing array, many headers (`03-slice-aliasing/`)

Plain English: copying a slice copies only the header, so every alias points
at the same backing array. Writing through one alias is visible through all
the others, and `append` may write past the end of one alias into space
another alias thinks is free. The three-index form `a[i:j:k]` caps the
capacity, so the next `append` is forced to allocate a fresh array instead
of trampling on data someone else holds.

Real world: several people holding bookmarks into the same book. Scribbling
on page 3 is seen by everyone. The three-index slice is a bookmark with a
"do not extend past page 5" sticker on it, so adding pages means photocopying
the book first.

```go
base := make([]int, 3, 5) // room to spare: cap 5
view := base[:2]
view[0] = 100
fmt.Println("base after writing through view:", base)

grown := append(view, 7) // cap 5 is enough: same backing array
grown[2] = 200
fmt.Println("base after writing through grown:", base)

fixed := base[2:3:5]     // len 1, cap 3 (k - i), not 1
over := append(fixed, 9) // cap 3 fits one more: still the same array
over[0] = -1
fmt.Println("base after a capped append:", base)
fmt.Println("over (separate array):", over)
```

Output:

```text
base after writing through view: [100 0 0]
base after writing through grown: [100 0 200]
base after a capped append: [100 0 -1]
over (separate array): [-1 9]
```

Beginner trap: `grown := append(view, 7)` looks harmless but writes `7`
into `base[2]` — the slot right after `view`'s length — because there was
spare capacity. Any code holding `base` or a longer view now sees `7` where
it expected a zero. This is the aliasing bug that shows up as "my slice
changed by itself". Cap the capacity (`view := base[:2:2]`, or
`base[i:j:k]` with `k == j`) when handing a slice to code that may append,
and the next append is forced to copy. A note on the example above: the
source comment claims `base[2:3:5]` has cap 1 and forces a new array, but
`k - i` is `3`, so the append still fits and `over` shares the backing
array — that is why `over[0] = -1` overwrites `base[2]`. To actually force
an allocation, use `base[2:3:3]`.

## 4. Slicing: bounds, capacity, and `copy` (`04-slicing/`)

Plain English: `a[i:j]` gives length `j-i` and keeps capacity running to the
end of the backing array; `a[i:j:k]` stops capacity at `k`. `copy(dst, src)`
moves `min(len(dst), len(src))` elements and never grows a slice. `clear(s)`
(Go 1.21+) zeroes every element in place.

Real world: cutting a piece of tape. `a[i:j]` is a strip with the backing
tape still attached behind it; `a[i:j:k]` is the same strip with the backing
tape trimmed to exactly the strip's length. `copy` is moving items between
two boxes: you can only move as many as the smaller box holds, and the boxes
do not resize.

```go
all := []int{0, 1, 2, 3, 4, 5, 6, 7}

middle := all[2:5] // len 3, cap 6: everything after index 2
capped := all[2:5:5]
fmt.Println("middle:", middle, "len:", len(middle), "cap:", cap(middle))
fmt.Println("capped:", capped, "len:", len(capped), "cap:", cap(capped))

fmt.Println("tail:", all[5:], "head:", all[:2])

dst := make([]int, 3)
n := copy(dst, all[6:])
fmt.Println("copy moved:", n, "dst:", dst)

clear(dst) // Go 1.21+: zero out in place
fmt.Println("cleared:", dst)

fmt.Println("truncated copy:", copy(dst, all)) // destination is smaller
fmt.Println("after copying all:", dst)
```

Output:

```text
middle: [2 3 4] len: 3 cap: 6
capped: [2 3 4] len: 3 cap: 3
tail: [5 6 7] head: [0 1]
copy moved: 2 dst: [6 7 0]
cleared: [0 0 0]
truncated copy: 3
after copying all: [0 1 2]
```

Beginner trap: an out-of-range index panics — `all[2:99]` is a runtime
panic, "slice bounds out of range" — and a high bound past `cap` panics too,
even when it is below `len` of the original. The second trap is `copy`'s
return value: it tells you how many elements moved, and if the destination is
smaller the tail of the source is silently dropped (`truncated copy: 3`,
then `dst` holds `[0 1 2]`). Always size the destination first, or read the
count. Third trap: bytes past the count are stale, not zero — `dst` still
held `[6 7 0]` before the `clear`, and `copy` does not zero anything.

## 5. Emptying a slice: three different zeros (`05-slice-emptying/`)

Plain English: `len 0`, `nil`, and `[:0]` are three different states. A
nil slice has no backing array; an empty non-nil slice has one; `s[:0]`
keeps the backing array and just resets the length to zero, which is the
cheapest reset but leaves the old values reachable. Drop the slice entirely
(`s = nil`) when it holds pointers or large values you want collected.

Real world: a whiteboard. `[:0]` is wiping the board but keeping it on the
wall — fast, and yesterday's notes bleed through if you are not careful.
`make([]T, 0, cap(s))` is buying a fresh board of the same size. `s = nil`
is throwing the board away so the janitor can take it.

```go
items := []string{"a", "b", "c"}

reused := items[:0] // len 0, same backing array: cheapest
reused = append(reused, "x")
fmt.Println("reused:", reused, "cap:", cap(reused), "still aliases items:", len(items) == cap(items))

fresh := make([]string, 0, len(items)) // new array, keeps the capacity
fresh = append(fresh, "y")
fmt.Println("fresh:", fresh)

items = nil // drop the reference so the memory can be collected
fmt.Printf("nil slice: %v len=%d isNil=%t\n", items, len(items), items == nil)
```

Output:

```text
reused: [x] cap: 3 still aliases items: true
fresh: [y]
nil slice: [] len=0 isNil=true
```

Beginner trap: `reused := items[:0]` is the cheapest reset, but the old
values stay reachable through the backing array until something overwrites
them — which is exactly the stale-bytes problem from `04-slicing` and the
mirror of buffer reuse. If the slice holds pointers, structs with pointer
fields, or anything the garbage collector should reclaim, zero the entries
explicitly (`clear(items)`) or drop the slice (`items = nil`) instead of
just reslicing.

## 6. Maps: hash tables with a small handle (`06-maps/`)

Plain English: a map is a hash table reached through a small handle. A nil
map reads fine (returns zero values) but panics the moment you write to it,
so `make(map[K]V)` or a literal is required before storing. A missing key
hands back the zero value, so the comma-ok form is the only way to tell
"absent" from "present and zero". `delete` on a missing key is a no-op,
`clear(m)` empties without making the map nil, and iteration order is
randomized — sort the keys for deterministic output.

Real world: a filing cabinet with labeled drawers. A nil map is a cabinet
that was never delivered: you can open a drawer and find nothing (read), but
trying to file a paper panics. `v, ok := m[k]` is asking the clerk "do you
have this file?" instead of guessing from an empty folder.

```go
var nilMap map[string]int
fmt.Printf("nil map: len=%d isNil=%t read=%d\n", len(nilMap), nilMap == nil, nilMap["ana"])
// nilMap["ana"] = 31 would panic: assignment to entry in nil map

ages := map[string]int{"ana": 31, "bo": 27}
ages["cy"] = 24

age, ok := ages["zz"]
fmt.Println("missing key ->", age, ok)
age, ok = ages["bo"]
fmt.Println("present key ->", age, ok)

delete(ages, "bo") // deleting a missing key is a no-op
fmt.Println("after delete:", len(ages), ages)

clear(ages) // Go 1.21+: empty it without rebuilding
fmt.Println("after clear:", len(ages), "isNil:", ages == nil)
```

For deterministic output, collect and sort the keys:

```go
scores := map[string]int{"c": 3, "a": 1, "b": 2}
keys := make([]string, 0, len(scores))
for k := range scores {
	keys = append(keys, k)
}
sort.Strings(keys)
for _, k := range keys {
	fmt.Printf("%s=%d ", k, scores[k])
}
fmt.Println()
```

Output:

```text
nil map: len=0 isNil=true read=0
missing key -> 0 false
present key -> 27 true
after delete: 2 map[ana:31 cy:24]
after clear: 0 isNil: false
a=1 b=2 c=3 
```

Beginner trap: writing to a nil map panics at runtime, not at compile time,
so `var m map[string]int; m["x"] = 1` builds fine and crashes when run.
Second trap: `m["missing"]` returns the zero value with no error, so a map
with `"username"` instead of `"name"` compiles and prints an empty string —
always use comma-ok when absence matters. Third trap: iteration order is
randomized per run (a hash seed plus jitter, anti hash-DoS armor), so
printing a map directly gives a different order every time; `fmt.Println`
sorts keys for you when it prints a map directly, but a manual `for k, v
:= range m` does not.

## 7. Maps as sets: `map[T]struct{}` (`07-map-sets/`)

Plain English: the set idiom is `map[T]struct{}`. The empty struct
`struct{}` occupies no memory, so a `bool` or `int` value would only waste
space per key. Membership is tested with comma-ok, and deleting entries
while ranging over a map is explicitly allowed.

Real world: a guest list where the only question is "is this name on it?" —
the value column is blank because you never need one. Writing `map[string]bool`
is like printing "yes" next to every name: correct, but it costs a byte per
guest for no reason.

```go
seen := map[string]struct{}{}
for _, w := range []string{"go", "go", "rust", "go", "zig"} {
	seen[w] = struct{}{}
}
_, goSeen := seen["go"]
_, cobolSeen := seen["cobol"]
fmt.Println("unique:", len(seen), "go seen:", goSeen, "cobol seen:", cobolSeen)

if _, dup := seen["go"]; dup {
	fmt.Println("go was already in the set")
}

nums := map[int]struct{}{1: {}, 2: {}, 3: {}}
_, has := nums[4]
fmt.Println("distinct numbers:", len(nums), "contains 4:", has)

// iterating a set yields only keys
for n := range nums {
	delete(nums, n) // deleting while ranging is allowed
}
fmt.Println("emptied by delete:", len(nums))
```

Output:

```text
unique: 3 go seen: true cobol seen: false
go was already in the set
distinct numbers: 3 contains 4: false
emptied by delete: 0
```

Beginner trap: reaching for `map[string]bool` out of habit works but wastes
memory and reads less clearly — `struct{}` signals "this is a set" to the
next reader. The second trap is assuming you cannot delete while ranging: for
maps you can (the example empties the set that way), but mutating a slice's
length during a `range` is a different story. Third trap: `seen["go"] =
struct{}{}` on a nil set panics — initialize with `map[string]struct{}{}`
or `make`.

## 8. Comparing maps: `==` is not allowed (`08-map-comparing/`)

Plain English: maps cannot be compared with `==` except against `nil`, so
`a == b` does not compile. Use `maps.Equal` (Go 1.21+) for comparable
values, and `reflect.DeepEqual` when the values are slices or maps. Two maps
are equal when they have the same keys and equal values.

Real world: two phone books. You cannot weigh the whole books against each
other on a scale (`==`), but you can page through them side by side
(`maps.Equal`) — and if the entries contain lists of numbers, you need a
deeper comparison (`reflect.DeepEqual`) that opens every list.

```go
a := map[string]int{"x": 1, "y": 2}
b := map[string]int{"x": 1, "y": 2}
// a == b would not compile: invalid operation, maps are not comparable
fmt.Println("a == nil:", a == nil)

fmt.Println("maps.Equal:", maps.Equal(a, b))
b["y"] = 99
fmt.Println("after change:", maps.Equal(a, b))

// with slice values maps.Equal will not compile, so reach for
// reflect.DeepEqual instead.
c := map[string][]int{"x": {1, 2}}
d := map[string][]int{"x": {1, 2}}
fmt.Println("reflect.DeepEqual:", reflect.DeepEqual(c, d))
d["x"][0] = 9
fmt.Println("the two maps are independent:", c["x"], d["x"])
```

Output:

```text
a == nil: false
maps.Equal: true
after change: false
reflect.DeepEqual: true
the two maps are independent: [1 2] [9 2]
```

Beginner trap: `a == b` is a compile error, not a runtime one — "invalid
operation: ... (map can only be compared to nil)". Do not try to work around
it with a loop that forgets key order; use the stdlib. Second trap:
`maps.Equal` requires the value type to be comparable, so
`map[string][]int` will not compile with it — that is the case for
`reflect.DeepEqual`. Third trap: assigning a map variable copies the header,
so both names refer to the same map and entry writes are visible through
both; the example's `c` and `d` are separate maps, which is why mutating
`d`'s slice does not affect `c`.

## 9. Structs: records that are values (`09-structs/`)

Plain English: a struct groups named fields into one type. Structs are values:
assignment copies all fields, and `==` compiles only when every field is
comparable. Embedding promotes the embedded type's fields and methods to the
outer type. Anonymous structs work anywhere a named type does, and two
struct types with identical fields (tags aside) convert to each other freely.

Real world: a photocopy of a form. Handing someone the form gives them their
own copy to scribble on; your original is untouched. Embedding is taping a
smaller form inside a bigger one and letting the reader use the inner fields
as if they were top-level.

```go
type Point struct {
	X, Y int
}

type Labeled struct {
	Point // embedded, so X and Y are promoted to Labeled
	Label string
}

type Weighted struct {
	Point
	Weight float64 // tags are ignored by conversions
}

p := Point{X: 1, Y: 2}
q := p // value copy
q.X = 99
fmt.Println("p:", p, "q:", q)

l := Labeled{Point: p, Label: "start"}
l.Y = 7 // promoted field
fmt.Println("labeled:", l, "X:", l.X, "Y:", l.Y)

// == compares every field
fmt.Println("p == Point{1, 2}:", p == Point{1, 2})
fmt.Println("l == equal literal:", l == Labeled{Point: Point{X: 1, Y: 7}, Label: "start"})
```

A struct with a slice field is not comparable:

```go
// BUG: a struct holding a slice has no == at all.
// _ = Bag{} == Bag{} would not compile: non-comparable field
```

Anonymous structs and conversions:

```go
anon := struct {
	Name string
	Age  int
}{Name: "zay", Age: 30}
fmt.Println("anonymous struct:", anon)

same := struct{ X, Y int }{1, 2}
fmt.Println("anonymous -> named:", Point(same) == p)
fmt.Println("named -> weighted:", Weighted{Point: p, Weight: 2.5})
```

Output:

```text
p: {1 2} q: {99 2}
labeled: {{1 7} start} X: 1 Y: 7
p == Point{1, 2}: true
l == equal literal: true
anonymous struct: {zay 30}
anonymous -> named: true
named -> weighted: {{1 2} 2.5}
```

Beginner trap: `q := p` is a copy, so `q.X = 99` does not touch `p` — but
if `Point` contained a slice or map field, that field would still be shared
(shallow copy), which is the pointers lesson recurring here. Second trap:
`==` on a struct only compiles if every field is comparable; a single
`[]int` field makes the whole type non-comparable and `_ = Bag{} == Bag{}`
fails to build. Third trap: struct tags are ignored by conversions, so a
tagged and an untagged struct with identical fields convert freely — useful
for JSON, but easy to miss when you expected tags to gate conversion.

## 10. Strings and bytes: always a copy (`10-strings-bytes/`)

Plain English: a Go string is an immutable sequence of bytes. `len(s)`
counts bytes, not characters; `utf8.RuneCountInString` counts runes.
Converting a string to `[]byte` or `[]rune` always copies, so editing the
result never touches the original. For building a string in a loop, use
`strings.Builder` (with `Grow`) instead of `s += x`, which reallocates
every time.

Real world: a string is a printed page — you cannot erase a letter, you can
only photocopy the page and edit the copy. `[]byte` is a photocopy of the
ink; `[]rune` is a photocopy of the letters, decoded from UTF-8. A builder
is a notebook you keep appending to, with `Grow` reserving pages up front so
you never tear the binding open mid-sentence.

```go
name := "héllo"
fmt.Println("bytes:", len(name), "runes:", utf8.RuneCountInString(name))

bs := []byte(name) // copy of the underlying bytes
bs[0] = 'H'
fmt.Println("from []byte:", string(bs), "| original:", name)

rs := []rune(name) // decoded characters
rs[1] = 'E'
fmt.Println("from []rune:", string(rs), "chars:", len(rs))

var sb strings.Builder
// PERF: Grow reserves the room once, so the loop below never reallocates.
sb.Grow(16)
for i := range 3 {
	fmt.Fprintf(&sb, "part%d-", i)
}
fmt.Println("builder:", strings.TrimSuffix(sb.String(), "-"))
```

Output:

```text
bytes: 6 runes: 5
from []byte: Héllo | original: héllo
from []rune: hEllo chars: 5
builder: part0-part1-part2
```

Beginner trap: `name[0]` is a byte, not a character — for `"héllo"` the
first byte is `0x68` but the second character spans two bytes, so editing
bytes can split a UTF-8 sequence and produce invalid text. Use `[]rune`
when you mean characters. Second trap: `s += x` in a loop allocates a new
string every iteration because strings are immutable; reach for
`strings.Builder` and call `Grow` once to reserve the room. Third trap:
conversions always copy, so there is no zero-copy "view" of a string as a
slice — plan for that allocation in hot paths.

## 11. Array and slice conversion: sharing versus copying (`11-array-conversion/`)

Plain English: `arr[:]` turns an array into a slice that shares the same
memory, so writing through the slice changes the array. `copy` is the only
route to an independent array. Converting a slice to an array (`[N]T(sl)`)
requires the slice to be at least `N` long and panics when it is not.

Real world: `arr[:]` is handing someone a window into your filing cabinet —
they can rearrange the papers and you see it. `copy(dup[:], sl)` is
photocopying the papers into a new cabinet. Converting a slice to a fixed
array is forcing a flexible folder into a rigid binder: it only fits if the
folder is big enough, and forcing it panics.

```go
arr := [4]int{1, 2, 3, 4}
sl := arr[:] // array to slice: shares the same memory
sl[0] = 100
fmt.Println("array after writing through the slice:", arr)

var dup [4]int
copy(dup[:], sl) // the only way to truly copy an array
fmt.Println("independent copy:", dup)

// slice to array: the slice must be at least as long as the array type
back := [3]int(sl[:3])
fmt.Println("slice to array:", back)

defer func() { // recover so the demo does not crash
	if r := recover(); r != nil {
		fmt.Println("recovered from:", r)
	}
}()
tooBig := [8]int(sl) // WARNING: len(sl) is 4, so this panics
fmt.Println("never reached:", tooBig)
```

Output:

```text
array after writing through the slice: [100 2 3 4]
independent copy: [100 2 3 4]
slice to array: [100 2 3]
recovered from: runtime error: cannot convert slice with length 4 to array or pointer to array with length 8
```

Beginner trap: `arr[:]` does not copy — it is the cheapest possible view,
and any write through it lands in the array. The second trap is the
direction: converting a slice to an array copies (it must, because the
array is a value), but it panics when the slice is shorter than the array
type, and the panic message names the lengths. Size the slice first, or
recover only at a real boundary, never as a habit.

## Rules I want to remember

- An array's length is part of its type; assignment copies every element.
  Use arrays for fixed buffers, slices for everything else.
- A slice is a (pointer, len, cap) header. `var s []T` is nil, `[]T{}`
  is empty, `make([]T, len, cap)` preallocates. Only `make` sets cap.
- `append` writes into spare capacity or allocates and copies. Growth is an
  implementation detail — never depend on the exact cap.
- Copies of a slice share the backing array, so writes through one alias are
  visible in all. `a[i:j:k]` caps capacity so the next `append` must copy
  (when `k == j`).
- `a[i:j]` keeps capacity to the end; `a[i:j:k]` stops it at `k`.
- `copy` moves `min(len(dst), len(src))` elements and never grows a slice.
  `clear(s)` zeroes elements; `clear(m)` empties a map.
- Resetting: `s[:0]` reuses the array (cheapest, old values stay
  reachable), `make([]T, 0, cap(s))` gets a fresh array, `s = nil`
  releases it. Zero entries explicitly when they hold pointers.
- `len 0`, `nil`, and `[:0]` are three different states.
- Converting a string to `[]byte` or `[]rune` copies, so the original is
  safe. Use `strings.Builder` (with `Grow`) for repeated concatenation.
  `len(string)` counts bytes; use `utf8.RuneCountInString` for characters.
- Maps: nil maps read fine but panic on write. `v, ok := m[k]` separates a
  missing key from a stored zero value. Iteration order is random — sort
  keys for deterministic output.
- `map[T]struct{}` is the set idiom; the empty struct occupies no memory.
  Deleting entries while ranging over a map is allowed.
- Maps can't be compared with `==` (only against nil). Use `maps.Equal`
  for comparable values, `reflect.DeepEqual` when values contain slices
  or maps.
- Structs are values: assignment copies (shallowly). `==` compiles only
  if every field is comparable. Structs with identical fields convert to
  each other, struct tags aside, and anonymous structs work anywhere a
  named type does.

## Exercises

Try these before peeking at the collapsible solutions.

1. **Append growth.** Given `s := make([]int, 0, 2)` and a loop that
   appends 0..5, what are the `(len, cap)` pairs after each append?

2. **Aliasing.** With `base := make([]int, 3, 5)` and `view := base[:2]`,
   after `view[0] = 100` and `grown := append(view, 7)`, what does
   `base` print? Why do `grown` and `base` share data, and how would you
   prevent it?

3. **Set membership.** Write a small program that counts the distinct words
   in `"go go rust go zig"` using `map[string]struct{}` and prints `3`.

4. **Struct conversion.** Define `type Point struct{ X, Y int }` and an
   anonymous struct with the same fields. Convert the anonymous value to
   `Point` and print whether it equals `Point{1, 2}`.

5. **Rune editing.** Change the second character of `"héllo"` to `"E"`
   without corrupting the UTF-8, and print the result plus its byte and
   rune counts.

6. **Independent array copy.** Given `arr := [4]int{1, 2, 3, 4}`, create a
   slice view, change element 0 through the slice, and prove the original
   array did change; then show how `copy` keeps a separate array untouched.

<details>
<summary>Solution 1 — Append growth</summary>

```go
package main

import "fmt"

func main() {
	s := make([]int, 0, 2)
	for i := range 6 {
		s = append(s, i)
		fmt.Printf("append %d -> len=%d cap=%d\n", i, len(s), cap(s))
	}
}
```

Output:

```text
append 0 -> len=1 cap=2
append 1 -> len=2 cap=2
append 2 -> len=3 cap=4
append 3 -> len=4 cap=4
append 4 -> len=5 cap=8
append 5 -> len=6 cap=8
```

</details>

<details>
<summary>Solution 2 — Aliasing</summary>

`base` prints `[100 0 200]`: `view[0] = 100` writes into `base[0]`, and
`append(view, 7)` writes `7` into `base[2]` because there is spare
capacity, then `grown[2] = 200` overwrites it. To prevent it, cap the
capacity when slicing: `view := base[:2:2]`, so the next `append`
allocates a fresh backing array.

```go
package main

import "fmt"

func main() {
	base := make([]int, 3, 5)
	view := base[:2:2] // capped: cap == len == 2
	view[0] = 100
	grown := append(view, 7) // allocates; base is untouched beyond [0]
	grown[1] = 200
	fmt.Println("base:", base)
	fmt.Println("grown:", grown)
}
```

</details>

<details>
<summary>Solution 3 — Set membership</summary>

```go
package main

import "fmt"

func main() {
	seen := map[string]struct{}{}
	for _, w := range []string{"go", "go", "rust", "go", "zig"} {
		seen[w] = struct{}{}
	}
	fmt.Println(len(seen)) // 3
}
```

</details>

<details>
<summary>Solution 4 — Struct conversion</summary>

```go
package main

import "fmt"

type Point struct{ X, Y int }

func main() {
	same := struct{ X, Y int }{1, 2}
	p := Point(same)
	fmt.Println(p == Point{1, 2}) // true
}
```

</details>

<details>
<summary>Solution 5 — Rune editing</summary>

```go
package main

import (
	"fmt"
	"unicode/utf8"
)

func main() {
	name := "héllo"
	rs := []rune(name)
	rs[1] = 'E'
	out := string(rs)
	fmt.Println(out)
	fmt.Println("bytes:", len(out), "runes:", utf8.RuneCountInString(out))
}
```

Output:

```text
hEllo
bytes: 5 runes: 5
```

</details>

<details>
<summary>Solution 6 — Independent array copy</summary>

```go
package main

import "fmt"

func main() {
	arr := [4]int{1, 2, 3, 4}
	view := arr[:] // shares memory with arr
	view[0] = 100
	fmt.Println("array:", arr) // [100 2 3 4] -- view writes through

	var dup [4]int
	copy(dup[:], arr[:]) // independent copy made after the mutation
	fmt.Println("copy: ", dup) // [100 2 3 4]

	// To keep the original untouched, copy first, then mutate the copy:
	var safe [4]int
	src := [4]int{1, 2, 3, 4}
	copy(safe[:], src[:])
	safe[0] = 999
	fmt.Println("original still:", [4]int{1, 2, 3, 4})
	fmt.Println("safe copy:    ", safe)
}
```

The point: `arr[:]` always shares memory with `arr`; `copy` is the only way
to get an independent array.

</details>