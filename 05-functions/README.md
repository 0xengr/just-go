# 05 — Functions

The big idea: functions are Go's only unit of behavior, and in this chapter they
wear every hat at once — several return values, named results, a stand-in for
optional parameters, and first-class values that can be passed around and that
remember things (closures). `defer` is the cleanup hook riding along the whole
time. Picture the chapter as the pit wall of a race team: every helper has a job,
hands back a report, and cleans up its station before the next stint.

## How to run

- `go build ./...` type-checks all five subexamples at once.
- `go run .` is only an orientation; every topic is its own program.
- Or use the chapter Makefile: `make run <topic>` (e.g. `make run 01-multiple-returns`); `make list` lists the topics and `make check` runs the gate.
- Suggested order for the first pass:

| Order | Run it | What it teaches you |
| ----- | ------ | ------------------- |
| 1 | `go run ./01-multiple-returns` | Returning `(value, ok)` instead of throwing |
| 2 | `go run ./02-named-return-shadowing` | Named results, blank `return`, and `:=` shadows |
| 3 | `go run ./03-options-variadic` | No named/optional params: struct options + `...T` |
| 4 | `go run ./04-anonymous-closure-func` | Functions as values; closures keep state |
| 5 | `go run ./05-defer-func` | `defer` cleanup, last-in-first-out |

Topics are numbered folders (`01-`, `02-`, …) in reading order, so `ls` alone
tells a newcomer where to begin. The numbers are the curriculum, not a version.

## 1. Multiple returns: an answer plus a receipt (`01-multiple-returns/`)

Plain English: a Go function can hand back more than one value, and the language
leans on that for `(value, ok)` and `(value, error)`. The function returns the
answer *and* a note about whether the answer is real, instead of throwing an
exception.

Real world: the lost-and-found desk. You ask for a driver's badge by ID and get
the badge *plus* a slip saying whether it was on file. `lookupUser("001")` comes
back as Alice's record and `true`; an unknown ID comes back as an empty record
and `false` — no explosion, just a slip that reads "not found".

```go
func lookupUser(id string) (user, bool) {
	user, found := users[id]
	return user, found
}
```

`main` checks the receipt before using the badge:

```go
u, ok := lookupUser("001")
if !ok {
	fmt.Println("Not found")
} else {
	fmt.Printf("%s <%s>\n", u.Name, u.Email)
}
```

Output:

```text
Alice <alice@example.com>
```

Beginner trap: a failed lookup still returns a **zero-valued** `user` (empty
name, empty email). Ignore `ok` and print `u.Name` anyway and you get a blank
that looks like data, with no error to warn you. Always look at the second
value — it is the entire reason the return exists.

## 2. Named returns: labeled crates (`02-named-return-shadowing/`)

Plain English: you may name a function's results in its signature. The names
become zero-valued variables that live in the body, and a bare `return` ships
whatever they currently hold. It buys readability on long signatures and cheap
error setup — at the price of hidden control flow.

Real world: instead of tossing loose parts at the loading dock
(`return num/denom, num%denom, nil`), you keep labeled crates — `result`,
`remainder`, `err` — by the door and just wave "ship it" when ready.

```go
func divide(num, denom int) (result int, err error) {
	if denom == 0 {
		return 0, errors.New("division by zero")
	}
	result = num / denom
	return // ships the current values of result and err
}
```

Output:

```text
divide(10, 2)                -> 5 <nil>
```

Beginner trap: one character flips it. Writing `result := num / denom` inside an
`if` (with `:=`) declares a **new** `result` scoped to that block, so the named
one is never assigned and the bare `return` hands back 0. The program's
`divideShadowed` does exactly that — it prints the right answer on its way out,
then returns nothing:

```text
     shadowed result = 5
divideShadowed(10, 2)        -> 0 <nil>
```

Assign to the named result with `=`, and prefer writing the values out
(`return result, nil`) over a blank `return` that forces the reader to scan
upward.

## 3. No optional params: a form and a bottomless basket (`03-options-variadic/`)

Plain English: Go has no named or optional parameters and no overloads. You get
two substitutes: pass a **struct** whose unfilled fields default to zero values
(that is your "optional"), or declare a **variadic** parameter `...T` to accept
any number of same-typed arguments.

Real world: ordering at a counter. The struct is the paper form — write your
name, skip age and it just reads back 0. The variadic is the bottomless basket:
`addNumbers("addNumbers", 1, 2, 3, ...)` keeps taking numbers until you stop.

```go
type greetOptions struct {
	Name    string
	Age     int
	Address string
}

func greetUser(opts greetOptions) {
	fmt.Printf("Hello %s, I know your age is %d\n", opts.Name, opts.Age)
}
```

```go
func addNumbers(label string, nums ...int) (string, int) {
	total := 0
	for _, num := range nums {
		total += num
	}
	return label, total
}
```

Output:

```text
Hello Mohammad, I know your age is 12
addNumbers 39
addNumbers2:  108
```

Beginner trap: inside the function a variadic parameter is just a `[]int` — a
nil slice when you pass none, which `range` handles fine — it *must* be the last
parameter, and at the call site you expand an existing slice with `nums...`
rather than passing it bare. That same `...` rule shows up as
`append(dst, src...)`.

## 4. Functions as values: reusable machines with a memory (`04-anonymous-closure-func/`)

Plain English: a function is a value like an `int`. You can store it in a
variable, pass it to another function, or return it from one. A **closure** is a
function value that remembers the variables around the place it was written —
including the ones that let it keep state between calls.

Real world: `makeAdder(10)` hands back a machine with its dial set to "+10", so
you never pass 10 again. `newCounter()` hands back a tally clicker that remembers
how many times you pressed it.

```go
func makeAdder(alwaysAdd int) func(int) int {
	return func(n int) int { return n + alwaysAdd }
}

func newCounter() func() int {
	count := 0
	return func() int { count++; return count }
}
```

Output:

```text
double(5)             -> 10
addTen(5)             -> 15
addTen(100)           -> 110
count() three times   -> 1 2 3
other()               -> 1
```

Beginner trap: a closure captures the **variable**, not a snapshot of its value.
Before Go 1.22, a closure built in a `for` loop saw the single reused loop
variable, so a slice of adders all added the *last* number. Each iteration now
gets its own copy (driven by the `go` line in `go.mod`), which is why the
program prints `2 + 100 -> 102`, `3 + 100 -> 103`, `4 + 100 -> 104`. Second trap:
function values compare only against `nil` — `f == g` will not compile — so use
`== nil` as the "is this actually set?" guard before calling.

## 5. `defer`: the cleanup that always runs (`05-defer-func/`)

Plain English: `defer` schedules a call to run when the surrounding function
returns — no matter how it returns, normal or early or panicking. Stack several
and they run last-in-first-out, like trays at a dish station.

Real world: opening a file, you book the cleanup the moment the resource exists;
if any later step fails, the close still happens. The classic shape:

```go
func fileLen(fileName string) (int, error) {
	f, err := os.Open(fileName)
	if err != nil {
		return 0, err
	}
	defer f.Close() // runs when fileLen returns

	info, err := f.Stat()
	if err != nil {
		return 0, err
	}
	return int(info.Size()), nil
}
```

Run against the fixture in this folder, `./05-defer-func/hello.txt`, it reports the size:

```text
size of ./05-defer-func/hello.txt is: 19 bytes
```

Stacking five defers shows the order — last deferred, first run:

```go
func runDefer() {
	defer fmt.Println("defer func 1")
	defer fmt.Println("defer func 2")
	// ... f3, f4, f5
}
```

Output:

```text
defer func 5
defer func 4
defer func 3
defer func 2
defer func 1
```

Beginner trap: deferred calls run last-in-first-out, and their **arguments are
evaluated at the `defer` line**, not when they finally run — `defer
fmt.Println(i)` in a loop prints the value `i` held at that moment. Also remember
`defer` is function-scoped, not block-scoped: a `defer` inside a loop keeps every
file open until the whole function returns, so hoist the loop body into its own
function when that matters.

## Notes

- A function signature is part of its type: `func(int) int` and `func(int, int) int`
  are different types, and so are `func(int) int` and `func(int) (int, error)`.
- Named return values are variables, not labels: you can read and assign them, and a
  bare `return` copies their current values into the result slots.
- Variadic parameters arrive as a slice, so `len`, `range`, `append`, and slicing
  all work on them unchanged.
- A closure outlives the function that created it: the captured variables move to
  the heap and stay alive as long as the closure does.
- `defer` arguments are evaluated immediately, but the call itself runs at return
  time; this is why `defer f.Close()` captures the file handle correctly.
- `defer` runs even when the function panics, which makes it the right place for
  cleanup, never for business logic.

## Exercises

Try these before peeking at the collapsible solutions.

1. **Safe division.** Turn `divide(a, b int) int` into
   `divide(a, b int) (int, error)` that returns `errors.New("division by zero")`
   when `b == 0`. Write a `main` that calls it with `(10, 2)` and `(10, 0)`
   and prints both outcomes.
   *Hint:* check `err != nil` before using the quotient.

   <details>
   <summary>Solution</summary>

   ```go
   package main

   import (
       "errors"
       "fmt"
   )

   func divide(a, b int) (int, error) {
       if b == 0 {
           return 0, errors.New("division by zero")
       }
       return a / b, nil
   }

   func main() {
       for _, c := range []struct{ a, b int }{{10, 2}, {10, 0}} {
           q, err := divide(c.a, c.b)
           if err != nil {
               fmt.Println("err:", err)
               continue
           }
           fmt.Println("quotient:", q)
       }
   }
   ```

   Expected output:

   ```text
   quotient: 5
   err: division by zero
   ```

   </details>

2. **fileLen with defer.** Write `fileLen(name string) (int, error)` that opens
   `name`, defers its `Close`, stats it, and returns the size in bytes (or the
   error). Print the size of `./05-defer-func/hello.txt`.
   *Hint:* `os.Open`, `defer f.Close()`, `f.Stat()`.

   <details>
   <summary>Solution</summary>

   ```go
   package main

   import (
       "fmt"
       "log"
       "os"
   )

   func fileLen(name string) (int, error) {
       f, err := os.Open(name)
       if err != nil {
           return 0, err
       }
       defer f.Close()

       info, err := f.Stat()
       if err != nil {
           return 0, err
       }
       return int(info.Size()), nil
   }

   func main() {
       size, err := fileLen("./05-defer-func/hello.txt")
       if err != nil {
           log.Fatal(err)
       }
       fmt.Printf("size is: %d bytes\n", size)
   }
   ```

   Expected output:

   ```text
   size is: 19 bytes
   ```

   </details>

3. **Prefixer.** Write `prefixer(prefix string) func(string) string` that
   returns a function prefixing its argument with `prefix + " "`. Test it with
   `helloPrefix := prefixer("Hello")` and print `helloPrefix("Bob")` and
   `helloPrefix("Maria")`.
   *Hint:* the returned function is a closure over `prefix`.

   <details>
   <summary>Solution</summary>

   ```go
   package main

   import "fmt"

   func prefixer(prefix string) func(string) string {
       return func(s string) string {
           return prefix + " " + s
       }
   }

   func main() {
       helloPrefix := prefixer("Hello")
       fmt.Println(helloPrefix("Bob"))   // Hello Bob
       fmt.Println(helloPrefix("Maria")) // Hello Maria
   }
   ```

   Expected output:

   ```text
   Hello Bob
   Hello Maria
   ```

   </details>

4. **Name, don't shadow.** The function below always returns 0. Fix it so it
   returns the halved value, and explain why the original printed the right
   answer but returned 0.

   ```go
   func half(n int) (result int) {
       if result := n / 2; result > 0 {
           fmt.Println("computed:", result)
       }
       return
   }
   ```

   *Hint:* `:=` declares a new `result` scoped to the `if`.

   <details>
   <summary>Solution</summary>

   ```go
   package main

   import "fmt"

   func half(n int) (result int) {
       if result = n / 2; result > 0 {
           fmt.Println("computed:", result)
       }
       return
   }

   func main() {
       fmt.Println("half(10) ->", half(10))
   }
   ```

   Expected output:

   ```text
   computed: 5
   half(10) -> 5
   ```

   The original used `:=`, which declared a *new* `result` living only inside
   the `if` body. The named return stayed at its zero value, so the bare
   `return` handed back 0. Assigning with `=` writes to the named result.

   </details>

## Rules I want to remember

- Return `(value, ok)` / `(value, error)` instead of throwing; check the second
  value, because zero values masquerade as real data.
- Named results are zero-valued locals; a bare `return` ships their current
  values. Assign with `=`, never shadow them with `:=`.
- No named/optional parameters: use struct options (unset fields = zero value)
  or a variadic `...T` (last param, becomes a slice). Expand a slice with
  `slice...`.
- Functions are values: assign them, pass them, return them, store them in
  slices.
- Closures capture variables, not snapshots, and keep them alive between calls.
  Since Go 1.22 each loop iteration has its own variable.
- Func values compare only to `nil`.
- `defer` runs at function return (LIFO, on panic too) and its args are evaluated
  at the `defer` line — it is *the* cleanup hook.
