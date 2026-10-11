# 06 — Pointers

The big idea: Go passes everything by value, so a function can never reach back
and change the variable you handed it. Pointers are how you hand over an address
instead of a copy — and, because Go has a garbage collector, that address is
just a note in a notebook, not a promise to free anything by hand. Think of it
as writing an address on a sticky note: the note is not the house, but you can
send someone to the house with it.

## How to run

- `go build ./...` type-checks the overview and the subexample at once.
- `go run .` is only an orientation; every topic is its own program.
- Or `make list` lists the topics and `make check` runs the chapter gate.
- Suggested order for the first pass:

| Order | Run it | What it teaches you |
| ----- | ------ | ------------------- |
| 1 | `go run .` | `&` takes an address, `*` reads the value back out |
| 2 | `go run ./01-call-by-value` | Every argument is a copy — and why slices and maps still surprise you |
| 3 | `go run ./02-pointers-mutable-params` | Passing a pointer so the caller sees the change — and why you still cannot swap the caller's pointer out |
| 4 | `go run ./03-pointers-last-resort` | nil, aliasing, and GC work: the reasons to reach for a pointer only on purpose |
| 5 | `go run ./04-map-slice-diffs` | Why a map or a slice reaches the caller's data with no pointer, and when a struct is the better contract |
| 6 | `go run ./05-pointer-passing-perf` | What copying a 1 MB struct actually costs, measured |

Topics are numbered folders (`01-`, `02-`, …) in reading order, so `ls` alone
tells a newcomer where to begin. The numbers are the curriculum, not a version.

## 1. A pointer is a sticky note with an address on it (`main.go`)

Plain English: a pointer is a variable that holds the memory address of another
value. `&x` hands you that address; `*p` reads whatever is stored there. Two
operators, and they are inverses of each other: take an address, read the value.

Real world: `x` is a crate on a shelf. `pointerToX` is not the crate, it is the
label on the crate. Hand someone the label and they can walk over and touch the
crate. Hand them a photo of the crate (`y := *pointerToX + " world"`) and they
only get a copy.

The second half of the program is the other end of the story: a pointer's zero value
is `nil`, which holds no address at all, so there is nothing to read. And
`new(T)` is the opposite — it allocates a zero `T` and hands back its address, so
that pointer is never `nil`.

```go
x := "hello"
pointerToX := &x // & is the address operator

fmt.Printf("address of x: %v\n", pointerToX)
fmt.Printf("value at pointerToX: %v\n", *pointerToX)

y := *pointerToX + " world"

var z *int
fmt.Println("z is nil:", z == nil)

// BUG: dereferencing z would panic, so it stays commented out:
// fmt.Println("dereference nil pointer:", *z)

// The guard that makes the dereference in the else branch safe:
// if z == nil {
// 	fmt.Println("z is nil, so there is nothing to dereference")
// } else {
// 	fmt.Println("dereference z pointer:", *z)
// }

intPtr := new(int)
strPtr := new(string)
boolPtr := new(bool)

fmt.Println("new() is never nil:", intPtr != nil)
fmt.Printf("zero values: int=%d str=%q bool=%t\n", *intPtr, *strPtr, *boolPtr)
```

Output:

```text
Welcome to pointers
address of x: 0x72390bce070
value at pointerToX: hello
value of y: hello world
z is nil: true
new() is never nil: true
zero values: int=0 str="" bool=false
```

Beginner trap: printing a pointer prints the *address*, and that address is
meaningless to a reader and different on every run. `0x72390bce070` tells you
nothing; `hello` tells you everything. If you want to see data, dereference
first. The one-line tell: a pointer printed with `%v` and no `*` in front of it
is a bug waiting to happen — and the `nil` line is the same trap with the
ground pulled out from under it, since `*z` there is a runtime panic rather than
a surprise number.

The matching trap on the other side: `new(T)` is *not* a fancy way to get a `nil`
pointer. It returns a real address, so a nil check after it guards against
nothing. Reach for `new` only when you want a pointer to a zero value and
nothing else — `n := 0; p := &n` is usually clearer.

## 2. Go is always call by value (`01-call-by-value/`)

Plain English: when you pass a variable to a function, Go makes a copy and hands
the function the copy. Reassigning a parameter changes the copy, never the
original. Slices and maps look like exceptions, and that is where most of the
confusion lives: the *slice header* and the *map header* get copied, but both
still point at the same underlying data.

Real world: you hand a colleague a photocopy of a form. They fill in the copy;
your original is untouched. Now hand them the key to the filing cabinet instead
and everything they file is visible to you — the cabinet itself was never copied.
Slices and maps are that key; `int`, `string`, and `struct` are the photocopy.

```go
func modifyBasic(i int, s string, p person) {
	i = 1 * 2
	s = "Goodbye"
	p.name = "Bob"
}

func modSlice(s []int) {
	for i := range s {
		s[i] = s[1] * 2
	}
	s = append(s, 10)
	fmt.Println("inside modSlice:", s)
}
```

Output:

```text
========== 1. BASIC VALUES ==========
outside: 2 Hello {20 Alice}

========== 2. SLICE ==========
before: [1 2 3]
inside modSlice: [4 4 8 10]
after: [4 4 8]

========== 3. SLICE REASSIGNMENT ==========
inside reassignSlice: [100 200 300]
outside: [1 2 3]

========== 4. MAP ==========
before: map[1:Alice 2:John]
inside modMap: map[2:hello 3:goodbye]
after: map[2:hello 3:goodbye]

========== 5. MAP REASSIGNMENT ==========
inside reassignMap: map[new:100]
outside: map[Alice:20]
outside: map[2:hello 3:goodbye]
```

The two halves to read carefully: `i`, `s`, and `p` come back exactly as they
started, because those parameters were copies. But `after: [4 4 8]` shows the
caller's slice *was* modified, even though only a copy of the slice value was
passed — writing to `s[i]` writes through to the shared backing array.

Sections 3 and 5 are the same test for the other half of the story, and both fail
the way a copy should. `reassignSlice` builds a brand-new slice and `outside` is
still `[1 2 3]`; `reassignMap` builds a brand-new map and the caller still holds
the old one. So a reference-typed value is not "shared" in the sense that matters
here — reassigning it is always local, exactly like an `int`.

Beginner trap: `s = append(s, 10)` inside a function changes only the local
parameter, so the caller's slice keeps its old length and no `append` result ever
escapes — that is section 2's `after: [4 4 8]` against `inside modSlice:
[4 4 8 10]`. Write to the existing elements (`s[i] = v`) and the caller sees it;
reassign the slice itself and the caller does not. Same map story: `m[k] = v` is
visible to the caller in section 4, while `m = map[...]{}` in section 5 is not.

## 3. Pointers indicate mutable params (`02-pointers-mutable-params/`)

Plain English: because every argument is a copy, writing to a `*int` parameter
writes to the *caller's* variable. That is the whole trick — and it stops at the
variable. A function can change what a pointer points at, but it cannot change
which variable the caller holds.

Real world: you hand someone a key to a locker. They can put something new inside
the locker, and you will see it when you open the door. Handing them a *new* key
does not change which key hangs on your ring — that is a copy of the key, not the
keyring.

```go
func updateOne(num int) { num = 100 }
func updateTwo(num *int) { *num = 200 }

func failedUpdate(g *int) {
	x := 10
	g = &x
}
```

Output:

```text
n1:  10
n2:  200
f:  <nil>
```

`n1` stays 10 because the copy was changed. `n2` becomes 200 because the function
wrote through the address. `f` stays nil: `failedUpdate` reassigned its own copy of
the pointer, so the caller's `f` never learned about `x`.

Beginner trap: `g = &x` looks like it should connect the caller to the new value.
It does not — reassigning a parameter is *always* local, pointer or not. If you
want the caller to end up holding a different pointer, return it: `func swap() *int`.

## 4. Pointers are a last resort (`03-pointers-last-resort/`)

Plain English: pointers are not the faster or more modern choice. They buy three
specific things — optionality, shared mutation, and cheap passing of large values —
and they charge for it in `nil` panics, aliasing that makes code harder to reason
about, and extra GC bookkeeping.

Real world: a shared spreadsheet is great when the whole team needs to edit the
same sheet. It is a nightmare when you only needed to read one number. Passing the
file around (by value) is calmer; sharing it (by pointer) is faster and needs
rules.

```go
func updateAge(age int)  { age = 30 }
func updateAgeTwo(age *int) { *age = 30 }
```

Output:

```text
25
30
```

The value version cannot reach the caller; the pointer version can. Neither is
"the Go way" in general — pick per call site.

Beginner trap: a pointer is not automatically cheaper. Copying a small struct is
both simpler and faster than adding a level of indirection; the perf win only
starts once the struct is large enough that the copy dominates.

## 5. Maps and slices are not pointers (`04-map-slice-diffs/`)

Plain English: this chapter's rule is "everything is a copy", and maps and slices
look like the exceptions. They are not exceptions to copying — they *are* copied,
like everything else. What gets copied is a small handle (a map's hash table
address, a slice's pointer/length/capacity triple), and that handle still reaches
the same underlying data. So `user["age"] = 25` inside a function is visible
outside, with no pointer anywhere.

Real world: a struct is a photocopy of a form — editing your copy leaves the
original alone. A map is a catalogue card: writing the same new phone number on
your card updates the listing everyone reads. The card was copied; the listing
was not. A slice is a bookmark with a page range: scribbling on page 3 is seen by
everyone holding the bookmark, but extending your bookmark's page range means
handing the new range back to the caller.

```go
func setAge(user map[string]int) { user["age"] = 25 }
func replaceMap(user map[string]int) { user = map[string]int{"age": 99} }

func scaleFirst(nums []int)         { nums[0] = nums[0] * 10 }
func appendDiscarding(nums []int)   { nums = append(nums, 40) }
func appendReturning(nums []int) []int { return append(nums, 40) }
```

Output:

```text
=== a map entry written inside a function is visible outside ===
user after setAge: map[age:25]
so no pointer was needed to change age: 25

=== assigning a new map to the parameter stays local ===
user after replaceMap: map[age:20]

=== a struct states the contract; a map only suggests it ===
  name: 
  email: user@example.com
  name: mohammad
  email: user@example.com

=== element writes are shared, length changes are not ===
after scaleFirst (element write): [100 20 30]
after appendDiscarding, len unchanged: [100 20 30] len = 3
after appendReturning and assigning: [100 20 30 40] len = 4

=== a struct by value is still only a shallow copy ===
caller sees the new label: map[admin:true beta:yes]

=== a map is right when the keys come from the data ===
  go     3
  is     2
  worth  1
```

The two halves to read carefully: `setAge` changes the caller's data while
`replaceMap` does not, even though both take a map. Writing an *entry* goes
through the shared handle; assigning the *handle* is an ordinary local
reassignment. Same asymmetry in slices: writing `nums[0]` mutates the shared
backing array, but `append` returns a *new* descriptor with a longer length, and
only the caller can install it.

Beginner trap: `data["name"]` on a key that does not exist is not an error, it is
the zero value — so a map with `"username"` instead of `"name"` compiles and
prints an empty string. A struct makes that mistake a build error. Related trap:
a struct passed by value is only a *shallow* copy, so a map field inside it is
still shared with the caller.

## 6. What passing cost actually is (`05-pointer-passing-perf/`)

Plain English: passing a 1 MB struct by value copies all 1 MB onto the stack;
passing a pointer copies one 8-byte address. The file also picks up chapter 1's
distinction: a nil `*Data` is not a zero `Data`, it is no `Data` at all — and
`new(Data)` is never nil, but it does allocate the full 1 MB on the heap.

```go
type Data struct{ Buffer [1_000_000]byte }

func processCopy(d Data)        {} // copies 1 MB
func processPointerPassing(d *Data) {} // copies 8 bytes
```

Measured with `go test -bench=. -benchmem ./05-pointer-passing-perf` on go 1.27.1:

```text
BenchmarkByValue-4     10846     126240 ns/op   1007620 B/op   1 allocs/op
BenchmarkByPointer-4   1000000000      0.5125 ns/op         0 B/op   0 allocs/op
```

Read the *gap*, not the absolute numbers: the by-value figure moved between 126k
and 424k ns/op across runs on the same machine. And the pointer row sits near zero
because the compiler inlines `byPointer` and discards the unused result, so it
measures close to an empty loop — "a million bytes copied" against "nothing copied",
not a literal 200,000x speedup.

Beginner trap: `new(Data)` looks like the safe alternative to `var p *Data`, and it
is safe to dereference — but `go build -gcflags=-m` reports `new(Data) escapes to
heap`, so it allocates the entire 1 MB anyway. Taking the address of a value you
already have (`&d`) is what costs 8 bytes and allocates nothing.

## 3. Pointers indicate mutable params (`pointers-mutable-params/`)

Plain English: because every argument is a copy, writing to a `*int` parameter
writes to the *caller's* variable. That is the whole trick — and it stops at the
variable. A function can change what a pointer points at, but it cannot change
which variable the caller holds.

Real world: you hand someone a key to a locker. They can put something new inside
the locker, and you will see it when you open the door. Handing them a *new* key
does not change which key hangs on your ring — that is a copy of the key, not the
keyring.

```go
func updateOne(num int) { num = 100 }
func updateTwo(num *int) { *num = 200 }

func failedUpdate(g *int) {
	x := 10
	g = &x
}
```

Output:

```text
n1:  10
n2:  200
f:  <nil>
```

`n1` stays 10 because the copy was changed. `n2` becomes 200 because the function
wrote through the address. `f` stays nil: `failedUpdate` reassigned its own copy of
the pointer, so the caller's `f` never learned about `x`.

Beginner trap: `g = &x` looks like it should connect the caller to the new value.
It does not — reassigning a parameter is *always* local, pointer or not. If you
want the caller to end up holding a different pointer, return it: `func swap() *int`.

## 4. Pointers are a last resort (`pointers-last-resort/`)

Plain English: pointers are not the faster or more modern choice. They buy three
specific things — optionality, shared mutation, and cheap passing of large values —
and they charge for it in `nil` panics, aliasing that makes code harder to reason
about, and extra GC bookkeeping.

Real world: a shared spreadsheet is great when the whole team needs to edit the
same sheet. It is a nightmare when you only needed to read one number. Passing the
file around (by value) is calmer; sharing it (by pointer) is faster and needs
rules.

```go
func updateAge(age int)  { age = 30 }
func updateAgeTwo(age *int) { *age = 30 }
```

Output:

```text
25
30
```

The value version cannot reach the caller; the pointer version can. Neither is
"the Go way" in general — pick per call site.

Beginner trap: a pointer is not automatically cheaper. Copying a small struct is
both simpler and faster than adding a level of indirection; the perf win only
starts once the struct is large enough that the copy dominates.

## 5. What passing cost actually is (`pointer-passing-perf/`)

Plain English: passing a 1 MB struct by value copies all 1 MB onto the stack;
passing a pointer copies one 8-byte address. The file also picks up chapter 1's
distinction: a nil `*Data` is not a zero `Data`, it is no `Data` at all — and
`new(Data)` is never nil, but it does allocate the full 1 MB on the heap.

```go
type Data struct{ Buffer [1_000_000]byte }

func processCopy(d Data)        {} // copies 1 MB
func processPointerPassing(d *Data) {} // copies 8 bytes
```

Measured with `go test -bench=. -benchmem ./pointer-passing-perf` on go 1.27.1:

```text
BenchmarkByValue-4     10846     126240 ns/op   1007620 B/op   1 allocs/op
BenchmarkByPointer-4   1000000000      0.5125 ns/op         0 B/op   0 allocs/op
```

Read the *gap*, not the absolute numbers: the by-value figure moved between 126k
and 424k ns/op across runs on the same machine. And the pointer row sits near zero
because the compiler inlines `byPointer` and discards the unused result, so it
measures close to an empty loop — "a million bytes copied" against "nothing copied",
not a literal 200,000x speedup.

Beginner trap: `new(Data)` looks like the safe alternative to `var p *Data`, and it
is safe to dereference — but `go build -gcflags=-m` reports `new(Data) escapes to
heap`, so it allocates the entire 1 MB anyway. Taking the address of a value you
already have (`&d`) is what costs 8 bytes and allocates nothing.

## Rules I want to remember

- `&x` is the address, `*p` is the value. They are inverses, not interchangeable.
- Print values, never addresses — an address is noise that changes every run.
- A nil pointer has nothing to dereference: `*p` on `p == nil` panics.
- Go has no `free`, no `delete(ptr)`, no dangling pointers. The garbage
  collector owns the memory; a pointer is just a value that gets copied around.
- Every argument is a copy. Reassigning a parameter never reaches the caller.
- `*p = v` reaches the caller; `p = &x` never does. To change which pointer the
  caller holds, return one.
- Pass by value for small structs — it is simpler and usually faster. Pass a
  pointer when the struct is large, when it may be nil, or when the callee must
  mutate it on purpose.
- `var p *T` is nil (nothing to read); `new(T)` is never nil but still allocates
  the whole value on the heap.
- Reference-typed values (slices, maps, channels, pointers) copy the header, not
  the data — so element writes are shared, and reassignment is not. Maps and
  slices are not a break in "everything is a copy"; the handle is what gets
  copied, and it still points at shared data.
- A missing map key is a zero value, not an error. `"username"` instead of
  `"name"` compiles and prints empty; a struct field makes it a build error.
- A struct passed by value is a shallow copy. Map, slice, pointer, or channel
  fields inside it are still shared with the caller.
- `append` hands back a new descriptor. Discard it and the caller's length is
  unchanged; return it and the caller assigns it.
- `unsafe` exists for the cases where the type system's guarantees are in your
  way. It is rare, it is not needed for normal Go, and the compiler will not
  save you from it.
