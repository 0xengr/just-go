# 07 — Types and Methods

The big idea: types are executable documentation. A defined type gives a concept
its own name and identity, and methods let values of that type carry their own
behavior. Go has no inheritance; it favors composition with embedding.

## How to run

- `go build ./...` type-checks all four subexamples at once.
- Every topic is its own program; this chapter has no overview `main.go`.
- Or `make list` lists the topics and `make check` runs the chapter gate.
- Suggested order for the first pass:

| Order | Run it | What it teaches you |
| ----- | ------ | ------------------- |
| 1 | `go run ./examples/01-user-defined-types` | Defined types, aliases, underlying types, and conversions |
| 2 | `go run ./examples/02-methods-and-receivers` | Methods, value vs pointer receivers, and update semantics |
| 3 | `go run ./examples/03-nil-receivers` | Writing methods that handle nil receivers |
| 4 | `go run ./examples/04-embedding-and-promotion` | Composition through embedding and promotion |

Topics are numbered folders (`01-`, `02-`, …) in reading order, so `ls` alone
tells a newcomer where to begin. The numbers are the curriculum, not a version.

## 1. Defined types: giving a concept its own name (`01-user-defined-types/`)

Plain English: `type Score int` creates a brand-new type. Its underlying
representation is `int`, but its identity is distinct, so the compiler refuses
to mix it with another `int`-based type without an explicit conversion. A type
alias (`type ID = int`) is the opposite: it is just a second spelling of the
same type, with no extra safety at all.

Real world: a `UserID` and a `Score` are both stored as integers, but handing a
`Score` where a `UserID` is expected is a bug. The type system is the contract:
`func banUser(id UserID)` will not accept a `Score`, even though both are `int`
underneath. An alias like `ID = int` is a nickname, not a contract — it buys
readability and nothing else.

```go
type UserID int
type Score int
type Email string
type HighScore Score

type ID = int

type MailCategory int

const (
	Uncategorized MailCategory = iota
	Personal
	Spam
	Social
)

func main() {
	var id UserID = 101
	var score Score = 95
	var email Email = "mohammad@example.com"

	fmt.Println(id, score, email) // 101 95 mohammad@example.com

	// var x Score = id // COMPILE ERROR: cannot use id (UserID) as Score

	var n int = 10
	var alias ID = n           // OK: ID is just int
	var converted UserID = UserID(n)
	fmt.Println("alias:", alias, "converted:", converted)

	var hs HighScore = 200
	hs = HighScore(score) // underlying type is int, not Score
	fmt.Println("high score:", hs)

	var literal UserID = 10
	fmt.Println("literal:", literal, "category:", Personal)
}
```

Output:

```text
101 95 mohammad@example.com
score + 10: 105
alias: 10 converted: 10
high score: 95
literal: 10 category: 1
```

Points to notice:

- `type HighScore Score` makes `HighScore`'s underlying type `int`, not `Score`.
  So `HighScore(score)` converts, but methods defined on `Score` do not carry
  over to `HighScore`.
- Two defined types with the same underlying type are not interchangeable:
  `var x Score = id` (UserID → Score) does not compile.
- Untyped constants initialize compatible defined types directly (`var literal
  UserID = 10`, `Personal` is a `MailCategory`); typed values need explicit
  conversion (`UserID(n)`, `HighScore(score)`).
- `score + 10` works because the untyped constant `10` adapts to `Score`;
  `score + id` would not, because both are defined types.
- `iota` starts at 0, so `Uncategorized` is 0 and `Personal` is 1 — that is why
  `category: 1` prints.

Beginner trap: thinking `type X Y` inherits Y's methods or identity. It does
not: X is a new type that merely shares Y's underlying representation. Also,
`type ID = int` gives zero type safety — use defined types for the concepts you
want the compiler to guard. (Methods can be attached to any defined type, not just
structs: `func (s Score) IsPassing() bool { return s >= 50 }`.)

## 2. Methods and receivers: behavior attached to a type (`02-methods-and-receivers/`)

Plain English: a method is a function with a receiver declared right after
`func`. The receiver is the value the method operates on, and its type decides
which type the method belongs to. Choose a value receiver (`func (p Person) …`)
when the method only reads or works on a copy; choose a pointer receiver
(`func (c *Counter) …`) when it must mutate the receiver or when the type is
large.

Real world: a Person's business card (`String`) can be printed from a photocopy;
the card itself does not change. A Counter is a tally clicker on a desk —
incrementing it changes the one on the desk, so you need its address, not a
copy.

```go
type Person struct {
	FirstName string
	LastName  string
	Age       int
}

func (p Person) String() string {
	return fmt.Sprintf("%s %s, age %d", p.FirstName, p.LastName, p.Age)
}

type Counter struct {
	total       int
	lastUpdated time.Time
}

func (c *Counter) Increment() {
	c.total++
	c.lastUpdated = time.Now()
}

func (c Counter) String() string {
	return fmt.Sprintf("total: %d, last updated: %v", c.total, c.lastUpdated)
}
```

```go
func main() {
	p := Person{FirstName: "Fred", LastName: "Fredson", Age: 52}
	fmt.Println(p.String())

	var c Counter
	fmt.Println(c.String())
	c.Increment()
	fmt.Println(c.String())

	pc := &Counter{}
	fmt.Println(pc.String())
	pc.Increment()

	doUpdateWrong(c)
	fmt.Println("after wrong:", c.String())
	doUpdateRight(&c)
	fmt.Println("after right:", c.String())
}
```

Output (timestamps vary run to run):

```text
Fred Fredson, age 52
total: 0, last updated: 0001-01-01 00:00:00 +0000 UTC
total: 1, last updated: 2026-10-19 10:03:35.272575505 +0545 +0545 m=+0.000046320
total: 0, last updated: 0001-01-01 00:00:00 +0000 UTC
wrong: total: 2, last updated: 2026-10-19 10:03:35.27266099 +0545 +0545 m=+0.000131794
after wrong: total: 1, last updated: 2026-10-19 10:03:35.272575505 +0545 +0545 m=+0.000046320
right: total: 2, last updated: 2026-10-19 10:03:35.272665105 +0545 +0545 m=+0.000135909
after right: total: 2, last updated: 2026-10-19 10:03:35.272665105 +0545 +0545 m=+0.000135909
```

Points to notice:

- Go automatically takes addresses and dereferences pointers for method calls on
  local variables: `c.Increment()` works on a `Counter` value, and `pc.String()`
  works on a `*Counter`.
- That auto-addressing only works when the value is addressable. Map elements and
  function return values are not, so a pointer-receiver method cannot be called on
  them directly (`m["a"].Inc()` fails to compile).
- Method sets: a value type's method set contains only value-receiver methods; a
  pointer type's contains both. That is why a type sometimes satisfies an
  interface only through a pointer.
- If any method of a type needs a pointer receiver, make all of them pointer
  receivers, for consistency and predictable interface satisfaction.
- `doUpdateWrong` receives a copy, so `Increment` mutates the copy and the caller
  is unchanged; `doUpdateRight` takes a pointer and mutates through it.
- Methods are declared at package level with a receiver; there is no `this` or
  `self` keyword, and they are not nested inside the type declaration.

Beginner trap: mixing receiver types on the same type can make it satisfy an
interface only via a pointer, surprising callers. Also, reassigning a pointer
parameter (`g = &x`) never reaches the caller — same rule as any other parameter.

## 3. Nil receivers: methods that survive a nil pointer (`03-nil-receivers/`)

Plain English: a pointer receiver may be `nil`, and the method is still callable.
If you write the method to check for `nil`, it can handle the empty case
gracefully. That is what makes recursive data structures like trees and lists
work: the empty structure is just a `nil` pointer, with no constructor needed.

Real world: an empty filing cabinet is not a crash, it is "nothing filed". A
tree's empty branch is a `nil` pointer; `Insert` and `Contains` must treat `nil`
as "no node here" instead of dereferencing and panicking.

```go
type IntTree struct {
	val         int
	left, right *IntTree
}

func (it *IntTree) Insert(val int) *IntTree {
	if it == nil {
		return &IntTree{val: val}
	}
	if val < it.val {
		it.left = it.left.Insert(val)
	} else if val > it.val {
		it.right = it.right.Insert(val)
	}
	return it
}

func (it *IntTree) Contains(val int) bool {
	switch {
	case it == nil:
		return false
	case val < it.val:
		return it.left.Contains(val)
	case val > it.val:
		return it.right.Contains(val)
	default:
		return true
	}
}

func main() {
	var it *IntTree
	it = it.Insert(5)
	it = it.Insert(3)
	it = it.Insert(10)
	it = it.Insert(2)

	fmt.Println(it.Contains(2))  // true
	fmt.Println(it.Contains(12)) // false
}
```

Output:

```text
true
false
```

Points to notice:

- Only pointer receivers can be `nil`. Calling a value-receiver method on a
  `nil` pointer panics, because Go dereferences it to copy the value.
- A `nil` receiver cannot be mutated in place, so `Insert` returns a new node
  (`*IntTree`) and the caller reassigns (`it = it.Insert(5)`). `Contains` only
  reads, so it returns `bool`.
- Guard every dereference: check `it == nil` before touching `it.val`.
- This pattern makes the zero value useful: `var it *IntTree` is an empty tree.

Beginner trap: writing `it.Insert(5)` without the reassignment leaves the tree
empty, because the new node is returned, not stored. Also, mixing in value
receivers would panic on `nil`.

## 4. Embedding and promotion: composition, not inheritance (`04-embedding-and-promotion/`)

Plain English: embedding an unnamed field of one type inside another promotes
its fields and methods to the outer type, so you can write `m.ID` and
`m.Description()` instead of `m.Employee.ID`. It is composition: the outer type
contains the inner one and reuses its behavior, but the outer type is not a
subtype of the inner one.

Real world: a Manager has an Employee badge; the badge's ID and Description are
available on the Manager without re-typing them, but a Manager is not an
Employee. You cannot hand a Manager to code expecting an Employee — you hand over
the badge (`m.Employee`) when an Employee is required.

```go
type Employee struct {
	Name string
	ID   string
}

func (e Employee) Description() string {
	return fmt.Sprintf("%s (%s)", e.Name, e.ID)
}

type Manager struct {
	Employee // embedded field
	Reports  []Employee
}

func main() {
	m := Manager{
		Employee: Employee{Name: "Bob Bobson", ID: "12345"},
		Reports:  []Employee{},
	}

	fmt.Println(m.ID)
	fmt.Println(m.Description())

	var e Employee = m.Employee // OK
	fmt.Println(e.Description())
	_ = e
}
```

Output:

```text
12345
Bob Bobson (12345)
Bob Bobson (12345)
```

Points to notice:

- The embedded field's name is its type name (`Employee`), so `m.Employee`
  reaches the inner value. `var e Employee = m.Employee` compiles; `var e
  Employee = m` does not.
- Promotion counts toward the outer type's method set, so embedding can help a
  type satisfy an interface through promoted methods.
- Name collisions: if two embedded types (or the outer type and an embedded
  one) share a field or method name, the unqualified access is ambiguous and
  will not compile — qualify with the embedded field name.
- The outer type can override a promoted method by defining its own with the
  same signature; the outer one wins for the outer type.
- An embedded pointer (`*Employee`) promotes too, and a `nil` embedded pointer
  makes promoted field access panic — initialize it.
- Embedding is not inheritance: no "is-a" relationship, no virtual dispatch, no
  automatic conversion.

Beginner trap: assuming `Manager` can be used where `Employee` is expected. It
cannot — pass `m.Employee`. Also forgetting that an embedded pointer field
defaults to `nil`.

## Rules I want to remember

- A defined type (`type X Y`) is a new type sharing Y's underlying type; it does
  not inherit methods or identity. An alias (`type X = Y`) is just a new name.
- Two defined types with the same underlying type are not interchangeable;
  convert explicitly.
- Untyped constants initialize compatible defined types directly; typed values
  need conversion.
- The underlying type of `type HighScore Score` is `int`, not `Score`.
- Methods are package-level functions with a receiver; the receiver (value vs
  pointer) decides mutability, nil-safety, and method sets.
- Value receivers work on copies and cannot handle `nil`; pointer receivers can
  mutate and can handle `nil` if written to.
- Go auto-addresses and dereferences only for addressable values (locals,
  variables); map elements and function results are not addressable.
- Pointer-receiver methods on a `nil` receiver run only if the method checks for
  `nil`; value receivers on `nil` pointers panic.
- Embedding promotes fields and methods, but is composition: the outer type is
  not the embedded type and cannot be assigned to it.
- Embedding can help satisfy interfaces via promoted methods; name collisions
  and `nil` embedded pointers are the usual traps.
## Exercises

1. **Defined type identity.** Define `type Meters float64` and `type Feet float64`.
   Show they cannot be mixed without conversion.
   *Hint:* `var m Meters = 10; var f Feet = Feet(m)`.
   <details><summary>Solution</summary>

   ```go
   type Meters float64
   type Feet float64
   var m Meters = 10
   // var f Feet = m // compile error
   var f Feet = Feet(m)
   fmt.Println(m, f)
   ```
   </details>

2. **Value vs pointer receiver.** Add a `Reset()` method to a counter using a
   pointer receiver and a `Value() int` using a value receiver; show mutation.
   *Hint:* `func (c *Counter) Reset() { c.n = 0 }`.
   <details><summary>Solution</summary>

   ```go
   type Counter struct{ n int }
   func (c *Counter) Inc() { c.n++ }
   func (c *Counter) Reset() { c.n = 0 }
   func (c Counter) Value() int { return c.n }
   ```
   </details>

3. **Nil receiver.** Write a `Tree` with a nil-safe `Insert` and `Contains`,
   build a tree from nil, and query a present and absent value.
   *Hint:* start with `var t *Tree`.
   <details><summary>Solution</summary>

   ```go
   var t *Tree
   t = t.Insert(5)
   fmt.Println(t.Contains(5), t.Contains(9))
   ```
   </details>

4. **Embedding.** Define `Address` and `Person` embedding `Address`; print a
   promoted field and extract the embedded value explicitly.
   *Hint:* `p.City` and `p.Address`.
   <details><summary>Solution</summary>

   ```go
   type Address struct{ City string }
   type Person struct {
       Name string
       Address
   }
   p := Person{Name: "Ada", Address: Address{City: "London"}}
   fmt.Println(p.City)
   a := p.Address
   _ = a
   ```
   </details>
