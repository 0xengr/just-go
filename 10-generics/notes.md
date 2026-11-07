# 10 — Generics

The big idea: Go is statically typed, so every parameter and field has a known
type at compile time. Generics let a function or a type leave that type
*unspecified until it is used*, while the compiler still checks every use. A
type parameter is a placeholder for a type, and a constraint is the rule that
says which types may replace it. That removes duplicated code without giving up
the compile-time safety that makes Go pleasant to refactor.

## 1. The duplication problem (`main.go`)

Plain English: before generics, the same algorithm had to be copied once per
type. Two functions that differ only in `int` versus `float64` are the same
logic written twice, and the copies drift apart. A type parameter writes the
body once and lets the caller choose the type.

Real world: numeric helpers, slice utilities, caches, queues, and trees all
repeat the same shape. Copying them for every element type is how bugs and
inconsistent fixes get in.

```go
func findMaxInt(nums []int) int {
	if len(nums) == 0 {
		return -1
	}
	max := nums[0]

	for _, num := range nums {
		if num > max {
			max = num
		}
	}

	return max
}

func findMaxFloat(nums []float64) float64 {
	if len(nums) == 0 {
		return -1
	}
	max := nums[0]

	for _, num := range nums {
		if num > max {
			max = num
		}
	}

	return max
}
```

The generic version keeps one body and puts the type in brackets:

```go
func findMaxGeneric[Num int | float64](nums []Num) Num {
	if len(nums) == 0 {
		return -1
	}
	max := nums[0]

	for _, num := range nums {
		if num > max {
			max = num
		}
	}

	return max
}
```

Output:

```text
Max Int value: 8
Max Float value: 8.60
8
8.6
```

Points to notice:

- `[Num int | float64]` declares one type parameter named `Num`. The part after
  the name is its constraint: the union of `int` and `float64`.
- `nums []Num` and the return type `Num` mean "a slice of whatever `Num` turns
  out to be". At the call, `Num` becomes `int` for the first slice and
  `float64` for the second.
- `fmt.Println` prints the generic results as `8` and `8.6`; `Printf` with
  `%.2f` is what adds the trailing zero in the earlier line.
- The compiler checks each instantiation separately. You cannot call
  `findMaxGeneric` with a `[]string`, because `string` is not in the union.
- Returning `-1` for an empty slice is a demo shortcut, not a good API: `-1` is
  a valid number. Production code returns `(Num, bool)` or an error so the
  caller can tell "empty" from "the maximum is -1".

Beginner trap: treating a type parameter as a runtime value. It does not exist
at run time; it is resolved when the compiler instantiates the function.

## 2. Type constraints (`examples/01-type-constraints/`)

Plain English: a constraint is an interface that describes a set of types.
`any` allows every type. A union such as `int | int64 | float64` allows exactly
those types. A named interface gives that set a reusable name, and the compiler
accepts only calls whose type argument belongs to the set.

Real world: a numeric helper that should reject strings, or a cache that needs
keys that can be compared. The constraint belongs in the signature, where the
compiler can enforce it.

```go
func print(value int) {
	fmt.Println(value)
}

// A variadic any parameter accepts several arguments, but it gives up
// compile-time type checking: the values do not have to share a type.
func printValue(value ...any) {
	fmt.Println(value...)
}

// Type constraints are usually written as interfaces.
type Number interface {
	int | int64 | float64
}

func findGenericMin[Num Number](nums []Num) Num {
	if len(nums) == 0 {
		return -1
	}

	min := nums[0]
	for _, num := range nums {
		if num < min {
			min = num
		}
	}
	return min
}
```

Output:

```text
123
123
123.3
false
Hello
4
4.34
```

Points to notice:

- `print` is fixed to `int`; `printValue` accepts anything because its
  parameter is `...any`. `any` is an alias for `interface{}`, so the values lose
  their static type inside the function.
- The commented generic `printValue[T int | string | bool | float32 | float64]`
  is the middle ground: it accepts one value whose type is in the list, and the
  compiler keeps that type.
- `Number` is an ordinary interface used only in the constraint position. It
  lists types, not methods, so it cannot be used as a variable type.
- `findGenericMin` returns `Num`, not `int`. `Println` prints `4` for the
  integer slice and `4.34` for the float slice because the concrete type was
  chosen at the call.
- The empty-slice `-1` has the same caveat as the overview: it is a valid value
  for both types, so the caller cannot distinguish it from a real minimum.

Beginner trap: writing `any` when a constraint would keep type safety, then
using type assertions to get the value back. That is the pre-generics style and
throws away the compiler's help.

## 3. Type inference (`examples/02-type-inference/`)

Plain English: you usually do not write the type arguments. The compiler infers
them from the argument types, the same way `:=` infers a variable's type. When a
type parameter appears only in the return type, there is nothing to infer from,
so you must state it.

Real world: `slices.Sort(nums)` should not require `slices.Sort[int](nums)`.
Inference keeps generic calls as short as ordinary calls.

```go
func multiply[Num int](a, b Num) Num {
	return a * b
}

func pair[K any, V any](key K, value V) {
	fmt.Println(key, value)
}

func create[T any]() T {
	var value T
	return value
}
```

```go
fmt.Println(multiply(12, 23))         // Num is inferred as int
fmt.Println(multiply[int](30, 45))    // the same call, written out
pair("name", "Mohammad")              // K=string, V=string
pair(1, true)                         // K=int, V=bool
val := create[string]()               // T cannot be inferred from arguments
fmt.Println(val)
```

Output:

```text
276
1350
name Mohammad
1 true

```

(The last line is the empty string that `create[string]()` returns, printed by
`fmt.Println`.)

Points to notice:

- `multiply(12, 23)` and `multiply[int](30, 45)` compile to the same call. The
  explicit form is available when you want the type to be obvious to a reader.
- `multiply(1.5, 2.5)` fails because `Num` is constrained to `int`:
  `cannot use 1.5 (untyped float constant) as int value in argument to multiply`.
- `pair` has two independent type parameters, so `pair(1, true)` is fine: `K`
  and `V` do not have to match.
- `create[T any]()` has no value parameters, so `T` appears only in the return
  type. The compiler has nothing to infer from and requires `create[string]()`;
  omitting the brackets is an error.
- A type argument is the concrete type supplied for a type parameter. Supplying
  it explicitly is called instantiation.
- Inference works left to right through the parameter list; it does not look at
  how the result is later used.

Beginner trap: expecting inference to work from the return type. It does not.
If `T` is only a return type, write it.

## 4. Type parameter lists and maps (`examples/03-type-generic-params-list/`)

Plain English: a type parameter list can mix independent parameters. In
`map[K]V`, the key must be comparable and the value can be anything, so the
natural signature is `func getValue[K comparable, V any](data map[K]V, key K) V`.

Real world: one lookup helper serves a map of product IDs to products and a map
of user IDs to users. Writing it once removes the copy-paste that appears when
two endpoints store different structs.

```go
func getValue[K comparable, V any](data map[K]V, key K) V {
	return data[key]
}

type Product struct {
	Name     string
	Price    int
	Discount int
}

type User struct {
	Name  string
	Email string
}
```

```go
product := map[int]Product{
	1: {Name: "samsung s21", Price: 1200, Discount: 10},
	2: {Name: "iphone 14", Price: 1500, Discount: 5},
	3: {Name: "redmi note 12", Price: 300, Discount: 15},
	4: {Name: "oneplus 11", Price: 900, Discount: 10},
	5: {Name: "google pixel 8", Price: 800, Discount: 20},
}

user := map[string]User{
	"u1": {Name: "Mohammad", Email: "mohammad@example.com"},
	// ...
}

fmt.Println(getValue(product, 1))
fmt.Println(getValue(user, "u2"))
```

Output:

```text
{samsung s21 1200 10}
{Ali ali@example.com}
{John john@example.com}
{google pixel 8 800 20}
```

Points to notice:

- `K comparable` is required because Go only permits comparable types as map
  keys. Without it the compiler rejects `map[K]V`.
- `V any` allows any value type, so the same function returns `Product` for one
  call and `User` for another; the return type is not `any`.
- A missing key returns `V`'s zero value, exactly like a plain map lookup. The
  signature here has no `ok` result, so callers cannot distinguish "missing"
  from "zero value". Add `(V, bool)` when that matters.
- `getValue(product, 1)` infers `K=int, V=Product`; `getValue(user, "u2")`
  infers `K=string, V=User`. No brackets are needed.

Beginner trap: using `any` as the key type. A map key must be comparable, so
`map[any]V` compiles but panics at run time if a slice is used as a key.

## 5. Generic types (`examples/04-generic-types/`)

Plain English: a type parameter can belong to a named type, not just to a
function. `Box[T]` is a struct that stores one `T` value; the caller chooses the
concrete type with `Box[int]` or `Box[string]`. Each method repeats `[T]` on the
receiver so the type stays in scope.

Real world: a box, result wrapper, cache entry, or queue should not be rewritten
for each payload type. One generic type keeps the behavior identical and the
compiler keeps the payload type exact.

```go
type Box[T any] struct {
	value  T
	filled bool
}

func (b *Box[T]) Put(value T) {
	b.value = value
	b.filled = true
}

func (b Box[T]) Get() (T, bool) {
	return b.value, b.filled
}

func (b Box[T]) IsEmpty() bool {
	return !b.filled
}
```

```go
var numberBox Box[int]
numberBox.Put(42)
value, ok := numberBox.Get()
fmt.Println(value, ok)
fmt.Println(numberBox.IsEmpty())

var wordBox Box[string]
fmt.Println(wordBox.IsEmpty())
wordBox.Put("hello")
fmt.Println(wordBox.Get())
```

Output:

```text
42 true
false
true
hello true
```

Points to notice:

- The receiver is written `(b *Box[T])`, not `(b *Box)`. The `[T]` keeps the
  type parameter visible to the method.
- The zero value of `Box[int]` is a `Box` with an empty `value` and `filled`
  set to `false`. It is ready to use; no constructor is required.
- `Put` has a pointer receiver because it changes the box. `Get` and `IsEmpty`
  have value receivers because they only read it.
- The `filled` field is necessary: without it, a box holding `0` would look
  exactly like a box that was never filled.
- `Get` returns `(T, bool)`, the same pattern as a map lookup. The bool says
  whether the box has a value, not what the value is.
- `numberBox.Put("not a number")` does not compile:
  `cannot use "not a number" (untyped string constant) as int value in argument to numberBox.Put`.
  That build error is the reason to use generics instead of `Box[any]`.
- `Box[any]` would also compile, but `Get` would return `any` and the caller
  would need a type assertion.

Beginner trap: forgetting `[T]` on the receiver. `func (b *Box) Put(v T)` does
not know what `T` is.

## 6. Generic functions that take functions (`examples/05-generic-functions/`)

Plain English: a generic function can accept another function as a value.
`All` and `Any` work on a slice of any element type, and the caller supplies the
check. The type parameter is inferred from the slice, and the function parameter
must accept that same element type.

Real world: validating a list, searching for one match, or checking that every
record is complete are the same loops in every service. Only the check changes.

```go
func All[T any](values []T, check func(T) bool) bool {
	for _, value := range values {
		if !check(value) {
			return false
		}
	}
	return true
}

func Any[T any](values []T, check func(T) bool) bool {
	for _, value := range values {
		if check(value) {
			return true
		}
	}
	return false
}
```

```go
numbers := []int{2, 4, 6}

fmt.Println(All(numbers, func(n int) bool {
	return n%2 == 0
}))
fmt.Println(Any(numbers, func(n int) bool {
	return n > 5
}))
fmt.Println(Any(numbers, func(n int) bool {
	return n > 10
}))

words := []string{"go", "generics"}

fmt.Println(All(words, func(word string) bool {
	return len(word) > 1
}))
fmt.Println(Any(words, func(word string) bool {
	return word == "go"
}))

fmt.Println(All([]int{}, func(n int) bool { return n > 0 }))
fmt.Println(Any([]int{}, func(n int) bool { return n > 0 }))
```

Output:

```text
true
true
false
true
true
true
false
```

Points to notice:

- `All` returns `false` as soon as one check fails; `Any` returns `true` as soon
  as one check passes. Both stop early.
- The closure `func(n int) bool` pins `T=int` for the numbers slice. The same
  code works for `[]string` because `T` is inferred again at the second call.
- On an empty slice, `All` is `true` and `Any` is `false`. That is vacuous
  truth: there is no value that fails, and no value that passes. Both are the
  standard mathematical answers, and both are easy to get wrong.
- `func(T) bool` is an ordinary function value. The caller can pass a closure,
  a named function, or a method expression; the generic function does not care.
- The type parameter belongs to the function, so two calls can use different
  element types without any shared state.

Beginner trap: assuming `All` on an empty slice is `false`. It is `true`, which
is why a validation built on `All` reports success for an empty list.

## 7. Type terms and the `~` prefix (`examples/06-type-terms/`)

Plain English: a type term is a concrete type inside a constraint. Type terms
are joined with `|` into a type set, and `~` before a term also matches any named
type whose underlying type is that term. The operators allowed in the function
body are the ones valid for every type in the set.

Real world: a `Dollars` type with underlying `int64` should use the same summing
helper as plain numbers. Without `~`, the helper would accept only the exact
built-in types and every domain type would need a wrapper function.

```go
type Number interface {
	~int | ~int64 | ~float64
}

func Sum[T Number](values []T) T {
	var total T
	for _, value := range values {
		total += value
	}
	return total
}

type Dollars int64
```

```go
fmt.Println(Sum([]int{1, 2, 3}))
fmt.Println(Sum([]int64{10, 20}))
fmt.Println(Sum([]float64{1.5, 2.5}))
fmt.Println(Sum([]Dollars{100, 250}))
```

Output:

```text
6
30
4
350
```

Points to notice:

- `+` is valid for `int`, `int64`, and `float64`, so the body compiles. If the
  set included a type without `+`, the body would fail.
- `%` is valid only for integers, so a constraint used with `%` must not include
  `~float64`. The operator you write decides how wide the set can be.
- `~int` matches `Dollars`, because `Dollars`'s underlying type is `int`.
  Without the tilde the set would accept only the exact type `int`.
- `var total T` starts at the zero value for whatever `T` is: `0` for integers,
  `0` for floats, `0` for `Dollars`.
- `Sum([]string{"a", "b"})` does not compile:
  `string does not satisfy Number (string missing in ~int | ~int64 | ~float64)`.
- Constants are limited by the set. A literal must fit every type in the
  constraint. With `~int8` in the set, `1_000` would stop compiling because
  1,000 does not fit in an 8-bit integer.
- An interface with a type term is valid only in the constraint position. It
  cannot be a variable, field, parameter, or return type.

Beginner trap: adding a float to a constraint used with `%`. The body then
fails to compile because `%` is not defined for floating-point types.

## 8. Interface constraints (`examples/07-interface-constraints/`)

Plain English: any interface can be a constraint, including one with methods.
`Describer` says a type can describe itself, and `Labeled[T Describer]` stores a
value of any type that satisfies it. Because the type parameter is `T`, the
value keeps its concrete type instead of being erased to the interface.

Real world: a report, log line, or UI row needs a label next to an item, and the
item may be a book, a movie, or any other type that knows how to describe itself.

```go
type Describer interface {
	Describe() string
}

type Labeled[T Describer] struct {
	Label string
	Value T
}

func (l Labeled[T]) Show() string {
	return l.Label + " -> " + l.Value.Describe()
}

func (l Labeled[T]) Unwrap() T {
	return l.Value
}
```

```go
book := Labeled[Book]{
	Label: "book",
	Value: Book{Title: "Learning Go", Pages: 450},
}
fmt.Println(book.Show())
fmt.Println(book.Unwrap().Pages)

movie := Labeled[Movie]{
	Label: "movie",
	Value: Movie{Title: "Arrival", Year: 2016},
}
fmt.Println(movie.Show())
```

Output:

```text
book -> Learning Go (450 pages)
450
movie -> Arrival (2016)
```

Points to notice:

- `T Describer` means `T` must have a `Describe() string` method. The compiler
  checks it at the point where `Labeled[Book]` is used.
- `Show` calls `Describe` on `l.Value`; the constraint is what makes that call
  legal.
- `Unwrap` returns `T`, not `Describer`. `book.Unwrap()` is a `Book`, so
  `.Pages` is available. A plain `struct { Value Describer }` would erase the
  concrete type and require a type assertion to get `Pages` back.
- `Book` and `Movie` each satisfy the constraint, so one generic type works for
  both without a shared base type.
- A type without `Describe` fails the constraint:
  `Plain does not satisfy Describer (missing method Describe)`.
- A constraint is compile-time only. `Describer` can be a variable type, but
  `Labeled` without a type argument cannot.

Beginner trap: using a generic type when an interface field would do. If the
code only calls `Describe` and never needs the concrete type, a plain
`Describer` field is simpler. The type parameter is worth it when the original
type must survive, as `Unwrap` shows.

## 9. Generic interfaces (`examples/08-generic-interfaces/`)

Plain English: an interface can have its own type parameters. `Comparator[T]`
compares a value with another value of the same type `T`. `Largest[T
Comparator[T]]` then requires both arguments to be the same type and to know how
to compare themselves.

Real world: versions, prices, dates, and scores each have their own ordering.
A generic `Largest` picks the winner without a type switch and without losing
the concrete type.

```go
type Comparator[T any] interface {
	Compare(T) int
}

func Largest[T Comparator[T]](a, b T) T {
	if a.Compare(b) >= 0 {
		return a
	}
	return b
}

type Version struct {
	Major int
	Minor int
}

func (v Version) Compare(other Version) int {
	if v.Major != other.Major {
		return v.Major - other.Major
	}
	return v.Minor - other.Minor
}
```

```go
fmt.Println(Largest(Version{Major: 1, Minor: 4}, Version{Major: 2, Minor: 0}))
fmt.Println(Largest(Version{Major: 2, Minor: 1}, Version{Major: 2, Minor: 0}))
fmt.Println(Largest(Version{Major: 1, Minor: 2}, Version{Major: 1, Minor: 2}))
```

Output:

```text
v2.0
v2.1
v1.2
```

Points to notice:

- `Comparator[T]` uses `T` in its method signature, and `Largest` uses the same
  `T` for both parameters and the return type. That is how the constraint says
  "same type, and comparable with itself".
- `Compare` returns an integer in the style of `cmp.Compare`: negative when the
  receiver is smaller, zero when equal, positive when larger.
- `Largest` returns `T`, so the caller gets a `Version` back, not a
  `Comparator[Version]`.
- `Version.String` is not required by the constraint; it exists so `Println`
  prints `v2.0` instead of the struct's field values.
- A type without `Compare` fails the constraint. A type with `Compare` for a
  different argument type also fails, because the method signature must match.
- Generic interfaces follow the same rule as type terms: they are valid in the
  constraint position, not as a variable's type.

Beginner trap: giving `Compare` the wrong parameter type, such as
`Compare(any) int`. That no longer satisfies `Comparator[Version]`, even though
the method looks similar.

## 10. The `comparable` pitfall (`examples/09-comparable/`)

Plain English: `comparable` means the type supports `==` and `!=`. Slices,
maps, and functions do not. An interface type is itself comparable, though, so
`Same[Item]` compiles even when the value stored inside the interface is a
slice. The comparison then panics at run time.

Real world: a generic equality helper or a cache keyed by `T` looks safe because
the constraint says `comparable`, but a caller can hand it an interface that
holds an uncomparable value. The failure is a panic, not a compile error.

```go
type Item interface {
	Name() string
}

type IntItem int

func (i IntItem) Name() string {
	return fmt.Sprintf("int:%d", i)
}

type ListItem []int

func (l ListItem) Name() string {
	return fmt.Sprintf("list:%v", l)
}

func Same[T comparable](a, b T) bool {
	return a == b
}
```

```go
fmt.Println(Same(7, 7))
fmt.Println(Same("go", "go"))
fmt.Println(Same(IntItem(7), IntItem(7)))

// ListItem itself does not satisfy comparable, so this does not compile:
// fmt.Println(Same(ListItem{1, 2}, ListItem{1, 2}))
// ListItem does not satisfy comparable

var a Item = IntItem(7)
var b Item = IntItem(7)
fmt.Println(Same(a, b))
```

Output:

```text
true
true
true
true
```

The commented block is the trap. Uncomment it and the program compiles, then
panics:

```text
panic: runtime error: comparing uncomparable type main.ListItem

goroutine 1 [running]:
main.Same[...](...)
	.../examples/09-comparable/main.go:...
main.main()
	.../examples/09-comparable/main.go:... +0x...
exit status 2
```

Points to notice:

- `Same(IntItem(7), IntItem(7))` is fine: the concrete type is an integer,
  which is comparable.
- `Same(ListItem{...}, ListItem{...})` is rejected at compile time because a
  slice type does not satisfy `comparable`.
- `Same(a, b)` where `a` and `b` have type `Item` compiles because an interface
  value is comparable. The compiler only sees the interface type, not the slice
  inside it.
- The panic text names the concrete type: `comparing uncomparable type
  main.ListItem`. That is the runtime finally discovering the problem.
- The same hole exists for `map[any]V` and for `==` on interface values. It is
  not specific to generics; generics just make it easier to trip over.

Beginner trap: assuming `comparable` guarantees a safe comparison for interface
values. It guarantees the static type is comparable; the dynamic value may not
be.

## 11. Generics in the standard library (`examples/10-standard-library/`)

Plain English: Go 1.21 added generic helpers for slices, maps, and ordering.
They are the functions to reach for before writing another loop by hand:
`slices.Sort`, `slices.Contains`, `slices.Max`, `maps.Clone`, `maps.Equal`, and
`cmp.Compare`.

Real world: sorting, searching, cloning a config map, and comparing two maps in
a test are everyday operations. The standard library already wrote and
benchmarked them.

```go
numbers := []int{5, 3, 8, 1}

slices.Sort(numbers)
fmt.Println(numbers)
fmt.Println(slices.Contains(numbers, 8))
fmt.Println(slices.Max(numbers), slices.Min(numbers))

words := []string{"go", "is", "fun"}
slices.SortStableFunc(words, func(a, b string) int {
	return cmp.Compare(len(a), len(b))
})
fmt.Println(words)

first := map[string]int{"a": 1}
second := maps.Clone(first)
second["a"] = 2

fmt.Println(first["a"], second["a"])
fmt.Println(maps.Equal(first, second))
```

Output:

```text
[1 3 5 8]
true
8 1
[go is fun]
1 2
false
```

Points to notice:

- `slices.Sort` accepts any ordered element type, so the same call works for
  `[]int`, `[]string`, and `[]float64`.
- `slices.Contains` returns a `bool`; `slices.Index` would return the position
  or `-1`.
- `slices.Max` and `slices.Min` are generic over ordered element types and
  return the element type, not `any`.
- `slices.SortStableFunc` takes a comparison returning negative, zero, or
  positive. `cmp.Compare(len(a), len(b))` is exactly that shape, and stable
  sorting keeps `go` before `is` when the lengths tie.
- `maps.Clone` makes a shallow copy: changing `second` does not change `first`,
  which is why the two values print `1` and `2`.
- `maps.Equal` compares keys and values, so it prints `false` after the clone is
  changed.
- Other useful generic helpers: `slices.Delete`, `slices.Insert`,
  `slices.Compact`, `slices.Equal`, `maps.Keys`, and `sync.OnceValue`.

Beginner trap: writing your own `contains` or `sort` loop when `slices` already
has a tested, optimized version. Check the package first.

## 12. Exercises

**`exercises/01-last-value/`** returns the last element of any slice. The bool
result separates "the last value is the zero value" from "the slice is empty".

```go
func Last[T any](values []T) (T, bool) {
	if len(values) == 0 {
		var zero T
		return zero, false
	}
	return values[len(values)-1], true
}
```

Output:

```text
3 true
generics true
0 false
true true
```

Points to notice:

- `Last([]int{1, 2, 3})` infers `T=int`; the string call infers `T=string`.
- The empty call returns `0, false`. `0` is the zero value of `int`, not a real
  element, and the bool says so.
- `Last([]bool{false, true})` returns `true, true`; a stored `false` would also
  be distinguishable from an empty slice because of the bool.
- `var zero T` is the only way to produce a value of an unknown type when there
  is nothing to copy.

**`exercises/02-map-keys/`** collects the keys of a map. `K` is the key type and
`V` is the value type; only `K` needs `comparable`.

```go
func Keys[K comparable, V any](m map[K]V) []K {
	keys := make([]K, 0, len(m))
	for key := range m {
		keys = append(keys, key)
	}
	return keys
}
```

Output:

```text
[Ada Bob]
[1 2]
[]
```

Points to notice:

- `Keys(ages)` infers `K=string, V=int`; `Keys(names)` infers `K=int, V=string`.
- The result is `[]K`, so the caller can sort it with `slices.Sort` without a
  cast.
- Map iteration order is random, so `main` sorts before printing. A function
  that returns keys cannot promise an order unless it sorts them itself.
- An empty map returns an empty slice, printed as `[]`.
- `V` is never used in the body, but it must still be declared so the map type
  matches. `any` is the right constraint because the function does not inspect
  values.

**`exercises/03-generic-queue/`** adds values at the back and removes them from
the front. `Queue[T]` is a generic type with a pointer receiver and a `Len`
helper.

```go
type Queue[T any] struct {
	values []T
}

func (q *Queue[T]) Push(value T) {
	q.values = append(q.values, value)
}

func (q *Queue[T]) Pop() (T, bool) {
	if len(q.values) == 0 {
		var zero T
		return zero, false
	}
	value := q.values[0]
	q.values = q.values[1:]
	return value, true
}
```

Output:

```text
1 true
2 true
3 true
empty queue: false
first second
```

Points to notice:

- FIFO order: the numbers go in as `1, 2, 3` and come out in the same order.
  A stack would reverse them.
- `Pop` takes the first element and reslices from index 1. The old backing array
  is released as the slice moves forward.
- The queue prints `empty queue: false` after the last pop, so the bool is the
  only reliable empty check.
- The string queue uses the exact same methods; only the type argument changes.

**`exercises/04-unique-values/`** removes duplicates while keeping the first
occurrence of each value and its original order.

```go
func Unique[T comparable](values []T) []T {
	seen := make(map[T]struct{}, len(values))
	result := make([]T, 0, len(values))

	for _, value := range values {
		if _, ok := seen[value]; ok {
			continue
		}
		seen[value] = struct{}{}
		result = append(result, value)
	}
	return result
}
```

Output:

```text
[1 2 3]
[go is fun]
[]
```

Points to notice:

- `T comparable` is required because `value` is used as a map key.
- `struct{}{}` is the canonical empty value for a set: it costs no memory, and
  membership is all the map needs to store.
- The result keeps the first occurrence because the check happens before the
  append. If the append came first, the last occurrence would win.
- An empty input returns an empty slice, not `nil`, because `make` allocates
  the result with length zero.

## Rules I want to remember

- A type parameter is a compile-time placeholder. It disappears at run time.
- `[T any]` accepts every type; `[T comparable]` accepts types that support
  `==` and `!=`; a union or a named interface narrows the set further.
- The operator you use in the body must be valid for every type in the
  constraint. `%` means integers only; `<` and `>` mean an ordered set.
- `~T` also matches named types whose underlying type is `T`. Use it for
  constraints meant for user-defined types.
- Constants must be representable in every type in the set; a constant that
  overflows the smallest type is rejected.
- Type inference works from value arguments. If a type parameter appears only in
  the return type, state it explicitly.
- Generic type methods cannot declare new type parameters. Use a generic
  function when the algorithm needs a second type.
- An interface with type terms is valid only in the constraint position. It
  cannot be a variable, field, parameter, or return type.
- `any` is an alias for `interface{}`. Prefer `any` in new code.
- `comparable` is a static guarantee. An interface value can still hold a
  slice, and comparing it panics at run time.
- Prefer an ordinary interface parameter when the function only needs methods.
  Use a type parameter when the concrete type must survive, as `Labeled.Unwrap`
  does.
- Do not convert an interface parameter to a type parameter just for speed.
  Measure first; generic dispatch is not automatically faster.
- Start with concrete types and plain functions. Introduce generics when you
  see the same body copied for a second type, not before.
- Use the standard library's generic helpers (`slices`, `maps`, `cmp`, and
  `sync.OnceValue`) instead of rewriting them.
- Go has no operator overloading, specialization, currying, metaprogramming,
  variadic type parameters, or parameterized methods. Design within that.
- A generic type keeps type safety: `Box[int]` rejects a string at build time,
  where `Box[any]` would accept it and fail later.
