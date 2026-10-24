# 08 — Interfaces

The big idea: an interface is a contract about behavior. It lists method
signatures and says nothing about fields, state, or implementation. A type
satisfies the interface by having those methods — no declaration, no
`implements` keyword, no registration. This is duck typing, and it is how Go
gets polymorphism without inheritance.

## 1. Interface basics: a set of methods (`main.go`)

Plain English: an interface names the methods a value must have. `speaker` below
says only that a value can `speak() string`; it does not say what the value is
or how it speaks. Any type with that exact method satisfies the interface, and a
function that takes `speaker` can then work with every one of them.

Real world: a conference stage does not care whether the speaker is a cat, a
dog, or a person. It needs someone who can speak, so the program asks for
`speaker`, not for a `cat`.

```go
type speaker interface {
	speak() string
}

type cat struct{}
type dog struct{}
type person struct{ name string }

func (c cat) speak() string    { return "Meow Meow" }
func (d dog) speak() string    { return "woof woof" }
func (p person) speak() string { return "Hi my name is: " + p.name }

func speaking(s speaker) {
	fmt.Println(s.speak())
}

// or
// using a variadic function
func saySomething(say ...speaker) {
	for _, s := range say {
		fmt.Println(s.speak())
	}
}
```

Output:

```text
Meow Meow
woof woof
Hi my name is: John Doe
Meow Meow
woof woof
Hi my name is: John Doe
Meow Meow
woof woof
Hi my name is: John Doe
```

Points to notice:

- The interface names behavior only: `speak() string`. It cannot hold state, and
  it says nothing about `cat`, `dog`, or `person` beyond that method.
- Satisfaction is automatic. None of the three types mentions `speaker`; each
  simply has the method, and the compiler checks the match where the value is
  used.
- `speaking` handles one value; `saySomething` takes `...speaker`, so it accepts
  zero or more values that satisfy the interface.
- The first three calls use the concrete functions from the earlier part of the
  file, the next three go through `speaking`, and the variadic call does it again
  — that is why each line appears three times.
- `speak` and `speaker` are unexported. A type outside this package could not
  satisfy `speaker`, because an unexported method name belongs to the package
  that declared it.

Beginner trap: an interface is not a base class. `cat` does not inherit from
`speaker`, and an interface value exposes only the interface's methods — `s.name`
does not compile until you assert the concrete type.

## 2. Duck typing: satisfaction is implicit (`examples/duck-typing/`)

Plain English: duck typing means the compiler cares about what a value can do,
not what it is called. If it has `Quack()`, it quacks. Go checks this
structurally, at compile time, so a typo in the method name or signature is a
compile error, not a runtime surprise.

Real world: a real duck, a toy duck, and a robot duck share nothing except the
quack. A function that needs a quack should ask for "something that quacks"
(`Quacker`), not for one specific kind of duck.

```go
type Quacker interface {
	Quack()
}

type RealDuck struct{}
type ToyDuck struct{}
type RobotDuck struct{}

func (d RealDuck) Quack()  { fmt.Println("RealDuck: Quack, Quack, Quack") }
func (t ToyDuck) Quack()   { fmt.Println("ToyDuck: Quack, Quack, Quack") }
func (r RobotDuck) Quack() { fmt.Println("RobotDuck: Quack, Quack, Quack") }

// using the Quacker interface: any value with a Quack method works here
func MakeItQuack(q Quacker) {
	q.Quack()
}

// or an even shorter, variadic form: pass as many Quackers as you like
func makeItQuack(quack ...Quacker) {
	for _, q := range quack {
		q.Quack()
	}
}
```

Output:

```text
RealDuck: Quack, Quack, Quack
ToyDuck: Quack, Quack, Quack
RobotDuck: Quack, Quack, Quack
RealDuck: Quack, Quack, Quack
ToyDuck: Quack, Quack, Quack
RobotDuck: Quack, Quack, Quack
```

Points to notice:

- No type declares `implements Quacker`. Having the method *is* the contract;
  there is nothing else to write.
- `MakeItQuack` takes one value, `makeItQuack` takes any number, including none.
  Variadic interface parameters are a compact way to accept a mixed set of
  implementations.
- All three types use value receivers, so both `RealDuck` and `*RealDuck`
  satisfy `Quacker` (more on method sets in the next section).
- The interface is small — one method — and defined next to the code that uses
  it. Go style favors narrow interfaces defined by the consumer rather than a
  wide interface defined by the producer.
- Because satisfaction is structural, renaming `Quack` to `Quacks` breaks the
  contract at the next assignment or call, and the compiler points at the value
  that no longer fits.

Beginner trap: a method must match the name and signature exactly. Also watch
the receiver: if `Quack` were declared with `func (d *RealDuck) Quack()`, then
only `&RealDuck{}` would satisfy `Quacker`, and `RealDuck{}` would not.

## 3. Polymorphism: one function, many types (`examples/polymorphism/`)

Plain English: polymorphism is the ability to appear in different forms. A
`Shape` can be a circle, a rectangle, or a square, and code written against
`Shape` works with all of them. The interface supplies the shared behavior; each
concrete type supplies its own implementation.

Real world: a shipping desk does not care whether a parcel is a box or an
envelope; it needs the area. One `Area()` method per shape is enough for a single
`areaPrint` to handle every current and future shape.

Without interfaces, every concrete type needs its own function:

```go
func areaOfCircle(c Circle)       { ... } // math.Pi * c.Radius * c.Radius
func areaOfRectangle(r Rectangle) { ... } // r.Length * r.Width
func areaOfSquare(s Square)       { ... } // s.Length * s.Length
```

Every new shape adds another function, and the caller has to pick the right one.
With an interface, the types carry the behavior:

```go
type Shape interface {
	Area() float64
}

func (c Circle) Area() float64    { return math.Pi * c.Radius * c.Radius }
func (r Rectangle) Area() float64 { return r.Length * r.Width }
func (s Square) Area() float64    { return s.Length * s.Length }

func areaPrint(s Shape) {
	fmt.Println(s.Area())
}
```

Output:

```text
Area of circle: 467.59
Area of rectangle: 60.00
Area of square: 16.00
467.59465056030473
60
16
```

Points to notice:

- The first three lines come from the no-interface functions, which use `%.2f`.
  The last three come from `areaPrint`, which uses `Println`'s default float
  formatting. The areas are the same; only the presentation differs.
- All three `Area` methods use value receivers, so both `Circle` and `*Circle`
  satisfy `Shape`. If `Area` were declared on `*Circle`, only `*Circle` would
  satisfy it and `areaPrint(circle)` would not compile.
- `areaPrint` knows nothing about circles, rectangles, or squares. It asks for
  `Shape` and calls `Area`, and Go dispatches to the right method through the
  interface value.
- Passing a value into `areaPrint` copies it into the interface. For large
  structs a pointer avoids the copy; measure before changing a receiver just for
  performance.
- Adding a `Triangle` means writing one method. `areaPrint` and its callers stay
  unchanged — that is the open/closed part of the story in practice.

Beginner trap: method sets decide satisfaction. A pointer-receiver method is not
in the value type's method set, so `var s Shape = Circle{}` fails when `Area` is
declared on `*Circle`. Pick value or pointer receivers once per type and keep
them consistent.

## 4. API design: accept interfaces, return concrete types (`examples/interface-api-design/`)

Plain English: a function should ask for the least it needs and give back the
most useful result. `loadPerson` needs bytes to decode, so it accepts
`io.Reader` — the one-method interface that strings, files, network connections,
and embedded files all satisfy. It returns `(Person, error)`: the concrete type,
so callers get every field instead of only the methods an interface exposes.

Real world: a parcel locker accepts any parcel that fits; it does not demand one
brand of box. `loadPerson2` demands a `string` and builds a reader itself, so a
file or an HTTP response has to be converted to a string first.

```go
// narrow: only strings work
func loadPerson2(s string) (Person, error) {
	var p Person
	err := json.NewDecoder(strings.NewReader(s)).Decode(&p)
	return p, err
}

// flexible: anything that can be read works
func loadPerson(r io.Reader) (Person, error) {
	var p Person
	err := json.NewDecoder(r).Decode(&p)
	return p, err
}
```

The same function now serves different sources:

```go
p1, err := loadPerson(strings.NewReader(`{"name":"john", "age":23}`))
if err != nil {
	fmt.Println("Error:", err)
	return
}
fmt.Println(p1)

file, err := exampleFile.Open("example.json") // embedded with //go:embed
if err != nil {
	fmt.Println("Error:", err)
	return
}
defer file.Close()

p2, err := loadPerson(file)
if err != nil {
	fmt.Println("Error:", err)
	return
}
fmt.Println(p2)
```

Output:

```text
{john 23}
{Jane 22}
```

Points to notice:

- `io.Reader` declares one method, `Read(p []byte) (n int, err error)`. That is
  all `json.NewDecoder` needs, so a `strings.Reader`, an `*os.File`, or an
  embedded file works without changing `loadPerson`.
- The return type is `Person`, not an interface. Callers can read its fields,
  compare values, and print them; returning an interface would hide all of that
  for no benefit.
- `loadPerson2` and `loadPerson` decode the same JSON; only the input contract
  differs. The first forces every caller to adapt; the second adapts to them.
- The example embeds `example.json` with `//go:embed`, so it runs from the
  chapter root or its own directory. `embed.FS.Open` returns an `fs.File`,
  which still reaches `loadPerson` as an `io.Reader`.
- `fmt.Println(p1)` prints `{john 23}` because `Person` has no `String` method;
  the default formatting lists the fields in declaration order.

Beginner trap: "accept interfaces" does not mean every parameter should be an
interface. Accept the narrowest behavior you actually call — here, reading
bytes. If a function needs a `Person`, take a `Person`.

## 5. Type switches: recovering the concrete type (`examples/type-assertion-switches/`)

Plain English: a type switch asks which concrete type is stored in an interface
value and runs the matching case. Inside a case, the variable has that concrete
type, so its fields are available directly. `event.(type)` is only legal in a
switch; outside one you use a type assertion instead.

Real world: one event stream carries sign-ups, failed payments, and shipped
orders. The processor must validate each shape differently, but all of them
arrive through the same `any` parameter.

```go
func processEvent(event any) error {
	switch e := event.(type) {
	case UserCreated:
		if e.UserID <= 0 || e.Email == "" {
			return errors.New("invalid user.created event")
		}
		fmt.Println("welcome email to:", e.Email)

	case PaymentFailed:
		if e.UserID <= 0 || e.Reason == "" {
			return errors.New("invalid payment-failed event")
		}
		fmt.Printf("payment failed for user %d: %s\n", e.UserID, e.Reason)

	case OrderShipped:
		if e.OrderID == "" || e.Email == "" {
			return errors.New("invalid order-shipped event")
		}
		fmt.Printf("shipping updates for order %s to %s\n", e.OrderID, e.Email)

	default:
		return fmt.Errorf("unsupported event type: %T", event)
	}
	return nil
}
```

Output:

```text
welcome email to: user@example.com
payment failed for user 42: card declined
shipping updates for order ORD-1001 to user@example.com
```

Points to notice:

- `e` changes type in every case: it is a `UserCreated` in the first, a
  `PaymentFailed` in the second, and an `OrderShipped` in the third. That is why
  `e.UserID`, `e.Reason`, and `e.OrderID` compile without assertions.
- `default` is the safety net. Without it, an unrecognized event would be
  ignored; `%T` puts the dynamic type in the error message so the failure is
  debuggable.
- The `events` slice is `[]any`, so one slice holds different concrete types.
  `processEvent` is the single place that turns them back into specific types.
- A type switch fits a small, known set of types. When the set should stay open,
  put the behavior on an interface instead — `Notifier` in the exercise does
  exactly that.
- Each case validates before using the fields. A type switch recovers the type;
  it does not make the value valid.

Beginner trap: a type switch matches exact dynamic types, not underlying types,
and it does not coerce anything. Forgetting `default` makes unknown types vanish
silently; `case nil` is how you catch a nil interface.

## 6. Interfaces in practice: `error` and `Notifier` (`exercises/notification-system/`)

Plain English: the exercise builds the same feature twice. First with a
`NotifyUser` function and a `switch` over a channel string; then with a
`Notifier` interface and one implementation per channel. The interface version
is open for extension: a new channel is a new type with a `Send` method, and
`NotifyTheUser` never changes. The `error` return works the same way — `error` is
an ordinary interface that any type with `Error() string` satisfies.

Real world: a notification service should not care whether a message leaves by
email, SMS, or push. The `error` interface lets each package report failures
without agreeing on a concrete error type.

```go
type Notifier interface {
	Send(to, message string) error
}

type emailNotifier struct{}
type smsNotifier struct{}
type pushNotifier struct{}

func (smsNotifier) Send(to, message string) error {
	if to == "" {
		return errors.New("phone number is required")
	}
	fmt.Printf("Sending SMS to %s: %s\n", to, message)
	return nil
}

func NotifyTheUser(n Notifier, recipient, message string) error {
	return n.Send(recipient, message)
}
```

Output:

```text
Email: Hello from Email
SMS: Hello from SMS
Push: Hello from Push
Sending EMAIL to 0xengr@gmail.com: Your order has shipped
Sending SMS to +9779800000000: Your verification code is 123456
Sending PUSH notification to device device-42: You have a new message
```

Points to notice:

- Switch dispatch lives in one function and must be edited for every new
  channel. It also silently drops an unknown channel, which is why the
  `"random"` call prints nothing.
- Interface dispatch moves the choice to the type: `NotifyTheUser` calls `Send`,
  and each notifier decides how to send. Adding a channel means adding a type,
  not editing an existing function.
- `error` is an interface: `type error interface { Error() string }`. Returning
  `nil` means success, and `errors.New` produces a value that satisfies it.
- Validation belongs in the implementation that needs it: `smsNotifier` rejects
  an empty phone number, while the email and push notifiers accept theirs
  without checks.
- The exercise discards results with `_ = NotifyTheUser(...)`. Real code should
  check the returned error and decide what to do next.

Beginner trap: ignoring a returned error, and assuming a `switch` without a
`default` reports an unknown channel. It does not — it does nothing at all.

## Rules I want to remember

- An interface lists method signatures; it never declares fields or state.
- Satisfaction is implicit and structural: the method set is the contract, and
  there is no `implements` keyword.
- A value type's method set contains only value-receiver methods; a pointer
  type's method set contains both. That decides whether `T`, `*T`, or both
  satisfy an interface.
- Keep receivers consistent on a type; mixing them makes satisfaction depend on
  whether a caller holds a value or a pointer.
- Small interfaces are idiomatic (`Quacker`, `Notifier`, `error`). Define them
  where they are used, not where the concrete type is defined.
- An interface with unexported methods can only be satisfied by types in the
  same package; exported behavior needs exported method names.
- Passing a value into an interface copies it. Pass a pointer if the copy is
  expensive or the method mutates the receiver.
- An interface value is `nil` only when it holds neither a type nor a value. A
  `nil` pointer stored in an interface makes the interface itself non-nil.
- `any` is an alias for `interface{}`. Use `v, ok := s.(Circle)` to assert the
  dynamic type safely; the single-value form `s.(Circle)` panics on a mismatch,
  and `switch v := s.(type)` branches on it.
- Accept interfaces, return concrete types: ask for the behavior you call and
  give callers the concrete value.
- A type switch recovers a concrete type from an interface; `default` is where
  unknown types belong, and `case nil` catches a nil interface.
- `error` is an interface; return `nil` for success and check what you get back.
- Variadic interface parameters (`...speaker`) accept any number of
  implementations, including none.

## Exercises

- `exercises/notification-system/` — the chapter's exercise. It builds the same
  notifier twice, with a `switch` and with an interface, and shows an `error`
  return. Run it with `go run ./exercises/notification-system`; the `"random"`
  call prints nothing because the `switch` has no `default`.

Things to try:

- Add a `slackNotifier` and call it through `NotifyTheUser`; notice that
  `NotifyUser` would need a new `case` too.
- Make `emailNotifier.Send` reject an empty address, then check the error in
  `main` instead of discarding it with `_ =`.
- Change `Circle.Area` to a pointer receiver and watch `areaPrint(circle)` stop
  compiling. Passing `&circle` fixes it — that is the method-set rule in action.
