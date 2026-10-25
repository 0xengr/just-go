# 09 — Errors

The big idea: Go has no exceptions. A function that can fail returns an `error`
value next to its result, and the caller checks it with an ordinary `if`. That
puts failure in the signature: you can see from the type that a call may fail,
and there is no hidden jump to a `catch` block. The `error` interface is tiny —
anything with `Error() string` is an error — so packages can report failures
without sharing a concrete type. Panic exists too, but it is for bugs and
impossible states, not for expected failures. `defer` is what still runs while a
panic unwinds.

## 1. Errors are values (`main.go`)

Plain English: a function that can fail returns an extra value of type `error`.
Success is `nil`; failure is a value that carries a message. The caller checks
it explicitly. Nothing is thrown and nothing is caught.

Real world: a calculator that divides. Division by zero is not a crash; it is a
case the caller must handle.

```go
func calcQuotientAndRemainder(numerator, denominator int) (int, int, error) {
	if denominator == 0 {
		return 0, 0, errors.New("denominator is 0")
	}
	return numerator / denominator, numerator % denominator, nil
}
```

```go
quotient, remainder, err := calcQuotientAndRemainder(30, 3)
if err != nil {
	fmt.Println(err)
	os.Exit(1)
}
fmt.Printf("quotient: %d\nremainder: %d\n", quotient, remainder)
```

Output:

```text
quotient: 10
remainder: 0
```

Points to notice:

- The error is the last return value by convention, and `nil` means success.
- The zero results returned with an error are placeholders. The caller must not
  use them when `err != nil`.
- `errors.New` is a function, not a method. It returns a value that satisfies
  `error`.
- `if err != nil` goes immediately after the call. That is the Go idiom; there
  is no special construct for detecting an error.
- `fmt.Println(err)` calls `Error()` automatically, so printing an error and
  printing `err.Error()` produce the same text.
- `os.Exit(1)` ends the program right away. Deferred calls do not run after it,
  which matters when cleanup is involved.

Beginner trap: ignoring the error with `_`, or checking it only sometimes.
Handle it at the call site, or return it so the caller can decide.

## 2. Syntax, runtime, and semantic errors (`examples/01-error-kinds/`)

Plain English: not every failure is an error value. A syntax error is rejected
by the compiler before the program runs. A runtime error is a panic while the
program runs. A semantic error compiles and runs but produces the wrong result,
so only tests and careful reading catch it.

Real world: a typo in a function name stops the build; indexing past the end of a
slice crashes the run; an off-by-one boundary quietly gives the wrong answer.

```go
func main() {
	// BUG: uncommenting the next line does not compile:
	// fmt.println("Enter your city:")
	// undefined: fmt.println (the package has Println, not println)
	fmt.Println("Enter your city:")

	nums := []int{2, 4, 6, 8}
	total := 0
	for i := range nums {
		total += nums[i]
	}

	// BUG: uncommenting this loop reads past the end of nums and panics:
	// for i := 0; i <= 10; i++ {
	//     total += nums[i]
	// }
	// panic: runtime error: index out of range [4] with length 4
}
```

```go
func shouldTakeCarBuggy(km int) string {
	if km > 2 {
		return "Take the car"
	}
	return "Going to walk today"
}

func shouldTakeCar(km int) string {
	if km >= 2 {
		return "Take the car"
	}
	return "Going to walk today"
}
```

Output:

```text
Enter your city:
Total: 20
buggy: Going to walk today
fixed: Take the car
```

Points to notice:

- `fmt.println` does not exist: Go is case-sensitive, and the exported function
  is `fmt.Println`. The compiler stops before any code runs.
- The panic message names both the bad index and the slice length:
  `index out of range [4] with length 4`. The valid indexes for a four-element
  slice are `0` to `3`.
- `shouldTakeCarBuggy` uses `> 2`, so exactly 2 km falls through to walking even
  though the intended rule is "2 km or more, take the car".
- Both failing snippets sit inside `main`, after the values they use are
  declared, so uncommenting one reproduces the compiler error or the panic. They
  stay commented out so the example builds and runs.
- `for i := range nums` is the safe walk: it stops at the slice's length by
  construction, so no bound check can fail.

Beginner trap: "it compiles" is not "it is correct", and a runtime panic is not
an `error` value you can check with `if err != nil`.

## 3. The `error` interface (`examples/02-error-interface/`)

Plain English: `error` is a predeclared interface with one method,
`Error() string`. A type satisfies it by having that method; there is no
`implements` keyword and no registration. A function that takes `error` accepts
every such type.

Real world: three different failures — network, file, validation — can travel
through one reporting function because all three can describe themselves as text.

```go
type error interface {
	Error() string
}

type NetworkError struct{}

func (ne *NetworkError) Error() string {
	return "network unavailable"
}

func printError(err error) {
	fmt.Println(err)
}
```

Output:

```text
network unavailable
failed opening file
validation error
```

Points to notice:

- `error` is declared once in the `builtin` package. You never declare it
  yourself; you implement it.
- The method set decides satisfaction. `Error` has a pointer receiver, so
  `*NetworkError` is an error but the plain `NetworkError` value is not. That is
  why `main` passes `&networkError`.
- `var _ error = (*NetworkError)(nil)` is a compile-time assertion. Assigning to
  the blank identifier makes the compiler verify the method signature; a typo
  fails the build instead of waiting for the first call.
- `printError` accepts the interface, not a concrete type, so it works with all
  three failures and cannot reach their fields.
- The interface value stores the pointer and its dynamic type. That is how the
  call to `Error()` finds the right method at run time.
- `fmt.Println(err)` prints the same text as `err.Error()` because `fmt` calls
  `Error()` for any value that implements `error`.

Beginner trap: returning a typed nil pointer as an error. If `err` holds a `nil`
pointer of a concrete type, the interface itself is not `nil`, because it still
has a dynamic type. Return a plain `nil` on success.

## 4. Creating error values (`examples/03-creating-error-values/`)

Plain English: `errors.New` builds a simple error from a string. Call it once at
package level to create a sentinel error — a well-known value that callers can
compare against — or at the point of failure when the message is all that
matters.

Real world: the `net/http` package declares package-level errors such as
`ErrHijacked` so callers can recognize a specific failure without matching text.

```go
var (
	ErrBodyNotAllowed  = errors.New("http: request method or response status code does not allow body")
	ErrHijacked        = errors.New("http: connection has been hijacked")
	ErrContentLength   = errors.New("http: wrote more than the declared Content-Length")
	ErrWriteAfterFlush = errors.New("unused")
)

func main() {
	errBadData := errors.New("some bad data")
	fmt.Printf("errBadData type: %T\n", errBadData) // *errors.errorString
}
```

Output:

```text
errBadData type: *errors.errorString
```

Points to notice:

- `errors.New` returns a pointer to an unexported `errorString` type. `%T`
  shows `*errors.errorString`, which is why you cannot depend on the concrete
  type; depend on the `error` interface instead.
- Each call to `errors.New` returns a distinct value, even when the text is
  identical. Two separately created errors with the same message are not equal.
- A sentinel such as `ErrHourlyRate` is created once and shared. Callers can
  compare against the same value because the identity is stable.
- Exported sentinel names start with `Err` by convention, for example
  `ErrBodyNotAllowed` and `ErrHijacked`.
- Error messages are lowercase and carry no trailing punctuation. They often
  start with the package or operation, as the `http:` prefix does here.
- The `net/http` examples show the real style: package-level values, one per
  condition, named so the caller can tell them apart. One of them,
  `ErrWriteAfterFlush`, is deprecated in `net/http` and no longer returned; it
  still shows the naming convention.

Beginner trap: matching errors by their text. Strings change; sentinel values
stay comparable. When errors get wrapped, plain `==` stops working too, and that
is a later topic.

## 5. Panic (`examples/04-panic/`)

Plain English: `panic` stops normal execution and starts unwinding the stack.
The runtime triggers one for things like an index out of range, a nil pointer
dereference, or integer division by zero. Code can trigger one with
`panic(value)`; the value can be any type, including an `error`.

Real world: a server that cannot find its required `PORT` should fail loudly at
startup instead of running half-configured. Bad user input, by contrast, should
come back as an error the caller can handle.

```go
func mustLoadEnv() {
	port := os.Getenv("PORT")
	if port == "" {
		panic("missing PORT check your .env file")
	}
	databaseURL := os.Getenv("DATABASE_URL")
	if databaseURL == "" {
		panic("missing DATABASE_URL check your .env file")
	}
}
```

Running the example as committed prints nothing, because both the runtime panic
loop and `mustLoadEnv()` are commented out. Uncomment the loop (and add `"fmt"`
to the imports) and it prints the valid elements before the crash:

```text
1
2
3
4
5
panic: runtime error: index out of range [5] with length 5

goroutine 1 [running]:
main.main()
	.../examples/04-panic/main.go:38 +0xad
exit status 2
```

Points to notice:

- A panic prints its value, then a goroutine stack trace, and the process exits
  with a non-zero status. The lines before the panic are the output that already
  ran.
- The runtime catches the out-of-range index itself. The developer panics in
  `mustLoadEnv` because a missing environment variable is a programming or
  deployment mistake, not an expected condition.
- `panic` accepts any value. Passing an `error` preserves the type and message;
  passing a string is fine for a simple message.
- An unhandled panic in any goroutine brings down the whole program. There is no
  catch block at the top.
- The panicking calls are commented out so `go run` exits cleanly. Uncomment
  them to observe the crash — the comments show the real output.
- `recover` can stop a panic during unwinding, but it is not covered in this
  chapter; it belongs with advanced error handling.

Beginner trap: using panic for ordinary failures such as a missing file or bad
form input. Callers expect an error they can check, not a crashed process.

## 6. Defer during a panic (`examples/05-defer-with-panic/`)

Plain English: a panic unwinds the stack, and every deferred call registered in
the unwinding frames still runs. They run in LIFO order, just like a normal
return. Only after the last defer do the panic value and stack trace print.
This is why `defer` is the right place for cleanup.

Real world: a function opens a file, registers `defer file.Close()`, then hits a
failure it cannot handle. The panic crashes the program, but the file descriptor
is still released on the way out.

```go
func main() {
	defer fmt.Println("defer first")
	defer fmt.Println("defer second")

	fmt.Println("BEFORE PANIC")
	readFile()
	fmt.Println("AFTER PANIC") // this won't run
}

func readFile() {
	file, err := os.Open("./examples/05-defer-with-panic/hello.txt")
	if err != nil {
		panic(err)
	}
	defer func() {
		fmt.Println("closing file")
		file.Close()
	}()

	panic("unexpected failure")
}
```

Output:

```text
BEFORE PANIC
closing file
defer second
defer first
panic: unexpected failure

goroutine 1 [running]:
main.readFile()
	.../examples/05-defer-with-panic/main.go:51 +0x78
main.main()
	.../examples/05-defer-with-panic/main.go:25 +0xc5
exit status 2
```

Points to notice:

- `readFile`'s deferred cleanup runs first because that function is the deepest
  frame being unwound. Then `main` runs its own defers in LIFO order: second,
  then first.
- The cleanup prints before the panic message. Defers run during unwinding; the
  crash report comes after every frame has finished.
- `AFTER PANIC` never runs. Once the panic starts, normal statements are
  abandoned; only deferred calls get a chance to execute.
- The file is closed even though the function panicked, which is the whole point
  of putting cleanup in `defer`.
- The path is relative to the chapter root, so run the example from there (or
  through `make run 05-defer-with-panic`). `hello.txt` sits next to `main.go`.
- The example is meant to crash. `go run` reports `exit status 2`, and the stack
  trace line numbers depend on the file.

Beginner trap: `os.Exit` is the exception to the rule. It terminates the process
immediately and skips deferred calls, so cleanup registered with `defer` does
not run. Return from the function instead when cleanup matters.

## 7. Exercises

**`exercises/01-calculate-pay/`** turns the same idea into a small validation
exercise. Two sentinel errors name the two failure modes, and `payDay` returns
one of them instead of crashing.

```go
var (
	ErrHourlyRate  = errors.New("invalid hourly rate")
	ErrHoursWorked = errors.New("invalid hours worked per week")
)

func payDay(hoursWorked, hourlyRate int) (int, error) {
	if hourlyRate < 10 || hourlyRate > 75 {
		return 0, ErrHourlyRate
	}
	if hoursWorked < 0 || hoursWorked > 80 {
		return 0, ErrHoursWorked
	}
	// ... pay calculation
}
```

Output:

```text
invalid hours worked per week
invalid hourly rate
6000
```

Points to notice:

- The first call has 81 hours, so it fails the hours check. The second call has
  an hourly rate of 5, so it fails the rate check. The third call is valid and
  returns 6000: 40 hours at 50 plus 40 overtime hours at 2× 50.
- Both failures use package-level values, so a caller could tell them apart
  without reading the message.
- The arithmetic only runs after both checks pass, so an invalid rate or hours
  value never reaches the pay calculation.

**`exercises/02-calculate-pay-panic/`** builds the same function again, but this
time the failures panic instead of returning.

```go
func payDay(hoursWorked, hourlyRate int) int {
	report := func() {
		fmt.Printf("HourlyRate: %d\nHoursWorked: %d\n", hourlyRate, hoursWorked)
	}
	defer report()

	if hourlyRate < 10 || hourlyRate > 75 {
		panic(ErrHourlyRate)
	}
	if hoursWorked < 0 || hoursWorked > 80 {
		panic(ErrHoursWorked)
	}
	// ... pay calculation
}
```

Output with the values in `main` (`payDay(70, 10)`):

```text
HourlyRate: 10
HoursWorked: 70
1000
```

Points to notice:

- The deferred `report` runs on both paths: a normal return and a panic. That is
  why the two lines print before the result.
- Passing `ErrHourlyRate` or `ErrHoursWorked` to `panic` carries an `error`
  value, not just a string.
- Change the arguments to `payDay(81, 50)` or `payDay(80, 5)` to see a panic.
  The report still prints first, then the panic message and stack trace appear.

## Rules I want to remember

- Errors are values. Return them last, use `nil` for success, and check
  `if err != nil` immediately.
- Do not ignore errors with `_` unless the failure truly does not matter, and
  say why in a comment.
- `error` is an interface with one method, `Error() string`. Satisfaction is
  implicit; there is no `implements` keyword.
- Method sets decide satisfaction: a pointer-receiver method means only `*T`
  satisfies the interface, not `T`.
- A compile-time assertion such as `var _ error = (*T)(nil)` catches a wrong
  method signature at build time.
- `errors.New` is for simple messages. Declare package-level `ErrFoo` values
  once when callers need to identify a specific failure.
- Error text is lowercase, without trailing punctuation, and usually names the
  operation or package.
- `fmt.Println(err)` calls `Error()` for you.
- A nil pointer stored in an `error` interface does not make the interface nil.
  Return a plain `nil` on success.
- Use panic for unrecoverable states and programmer mistakes, not for expected
  failures such as bad input or a missing file.
- A panic accepts any value, including an `error`, and an unhandled panic brings
  down the whole program.
- Deferred calls run during a panic, in LIFO order, which makes `defer` reliable
  for cleanup.
- `os.Exit` skips deferred calls. It is the exception to the cleanup rule.
