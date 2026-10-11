# 04 — Blocks, Shadows, and Control Structures

The big idea: Go has only three control keywords — `if`, `for`, `switch`
(plus a rarely-used `goto`). Everything else in this chapter is about
*where* your variables live (blocks) and *how* you accidentally hide them
(shadowing). Think of it as learning the rules of the race control tower:
who can see which screen, and which button actually stops which race.

## How to run

- `go build ./...` type-checks all nine subexamples at once.
- `go run .` is only an orientation; every topic is its own program.
- Or `make run 01-blocks`, `make run 06-for-range`, … — the chapter Makefile also gives `make list` and `make check`.
- Suggested order for the first pass:

| Order | Run it | What it teaches you |
| ----- | ------ | ------------------- |
| 1 | `go run ./01-blocks` | Blocks nest; inner sees outer, not the reverse |
| 2 | `go run ./02-shadowing` | `:=` can silently create a lookalike variable |
| 3 | `go run ./03-if-scope` | `if x := ...; cond` keeps temporaries local |
| 4 | `go run ./04-for-complete` | The classic `for i := 0; i < n; i++` and its variations |
| 5 | `go run ./05-for-while` | The `while` shape, infinite loops, `break`/`continue` |
| 6 | `go run ./06-for-range` | The everyday loop over slices, maps, and strings |
| 7 | `go run ./07-for-labels` | Steering an *outer* loop from an inner one |
| 8 | `go run ./08-switch-cases` | Matching values, blank switches, no fallthrough |
| 9 | `go run ./09-goto` | The one rare case where jumping forward reads better |

Topics are numbered folders (`01-`, `02-`, …) in reading order, so `ls` alone
tells a newcomer where to begin. The numbers are the curriculum, not a version.
Here the numbering simply encodes the table above.

## 1. Blocks are rooms with one-way glass (`01-blocks/`)

Plain English: every pair of braces `{}` is a room. Code inside a room can
read the whiteboard of every room outside it, but nobody outside can read
what was written inside once the door closes.

Real world: the season name (`season = "spring series"`) is pinned on the
building wall — every function sees it. The lap count (`laps := 3`) lives
in the meeting room called `main` — every block inside `main` sees it.
The best-lap time (`best := 58.4`) is scribbled inside a tiny huddle room
(a bare `{}` block) and thrown away when the huddle ends.

```go
var season = "spring series" // building wall: whole file sees it

func main() {
	laps := 3 // meeting room: everything below in main sees it
	if laps > 0 {
		fmt.Println(season, "laps:", laps) // inner reads outer: fine
		{
			best := 58.4 // huddle room: dies at the closing brace
			fmt.Println("best this stint:", best)
		}
		// fmt.Println(best) // won't compile: best is gone
	}
}
```

Beginner trap: uncommenting either commented-out line fails with
`undefined: best` / `undefined: i`. That error is the compiler telling you
the room has been demolished. When a temporary is only needed for three
lines, a bare block (or a tighter `if` scope) keeps it from leaking.

## 2. Shadowing: the evil twin (`02-shadowing/`)

Plain English: if you declare a new variable with the same name as an outer
one, you get two variables wearing the same jersey. Inside the inner block
you only ever see the twin; the outer one is still there, just unreachable.

Real world: race control has a fuel board showing `fuel := 10`. A pit crew
member writes `fuel := 5` on their own clipboard inside the pit box. The
main board still says 10 — the crew just can't see it while holding their
clipboard.

```go
fuel := 10
if fuel > 5 {
	fuel := 5 // a NEW fuel, alive only inside this if body
	fmt.Println("inside:", fuel) // 5
}
fmt.Println("after:", fuel) // 10 — the outer board never changed
```

Output:

```text
before: 10
inside: 5
after: 10
```

The sneaky version uses multi-assign. `fuel, pit := 5, 20` looks like it
updates `fuel`, but `:=` reuses only names from the *same* block — since
`pit` is new, the whole line declares, and `fuel` shadows again. The fix is
boring on purpose: use plain `=` when the variable already exists.

```go
fuel = 5 // assignment, not declaration: the real board changes
```

Beginner trap: never name a local `fmt`, `len`, or `true`. They are not
keywords — they live in the universe/file blocks — so your local silently
hides the package or builtin for the rest of the block, and `fmt.Println`
suddenly means "field Println on my string". `go vet` won't save you here;
a shadowing linter will.

## 3. `if` with its own private clipboard (`03-if-scope/`)

Plain English: `if x := ...; condition` creates `x` just for that one
`if/else` chain. It is visible in every branch and gone afterwards.

Real world: measuring a driver's callsign to decide which board it fits on.
The length `n` matters for exactly one decision — after that, keeping it
around is clutter.

```go
if n := len(name); n == 0 {
	fmt.Println("no callsign entered")
} else if n > 5 {
	fmt.Println(name, "is a long callsign:", n)
} else {
	fmt.Println(name, "fits the board:", n) // "bo fits the board: 2"
}
// fmt.Println(n) // won't compile: n belonged to the decision, not to main
```

Beginner trap: the short statement shadows too. `if cutoff := 4; ...`
inside a function that already has `cutoff := 6` creates a lookalike for
one branch. Use the slot only for fresh temporaries, never to smuggle in
side-effect calls or assignments to existing variables.

## 4. The complete `for`: start, check, step (`04-for-complete/`)

Plain English: `for i := 0; i < 5; i++` is three instructions — set up the
counter, check before every lap, step after every lap. The setup must use
`:=`; the check must be a boolean.

Real world: sending cars out for laps 0 through 4, then reusing the pattern
for a slice window (positions 1 to 3 of `[9 7 8 6 10]`):

```go
for i := 0; i < 5; i++ {
	fmt.Print(i, " ") // 0 1 2 3 4
}

scores := []int{9, 7, 8, 6, 10}
for k := 1; k < len(scores)-1; k++ {
	fmt.Print(scores[k], " ") // 7 8 6
}
```

You may drop any clause. Hoist the setup when it comes from earlier work
(`i := 0; for ; i < 3; i++`), or drop the step when the stride needs logic
inside the loop. Rule of thumb: reaching for `break`/`continue` just to
walk a sub-window means a plain `for` with honest bounds would read better.

## 5. `for` as `while`, `for` as forever (`05-for-while/`)

Plain English: Go has no `while` keyword — you just leave pieces off.
Drop setup and step (and the semicolons) and `for speed < 8` *is* while.
Drop the condition too and `for {}` runs until something breaks it.

Real world: doubling speed `1 2 4` until it hits the limit; retrying a
connection up to 3 attempts; classifying 12 laps as pit / charge /
pit-charge / green without nesting:

```go
speed := 1
for speed < 8 { // the while shape
	fmt.Print(speed, " ")
	speed *= 2
}

attempts := 0
for { // forever, until the explicit exit
	attempts++
	if attempts >= 3 {
		break
	}
}
```

The `continue` style keeps every case at the same indent instead of a
pyramid of `if/else`:

```go
for lap := 1; lap <= 12; lap++ {
	if lap%2 == 0 && lap%3 == 0 {
		fmt.Println(lap, "pit-charge")
		continue
	}
	if lap%2 == 0 {
		fmt.Println(lap, "pit")
		continue
	}
	// ... charge, then green
}
```

Need do-while ("run at least once")? Write
`for { work; if !cond { break } }` — note the flipped condition versus
Java: Go states how to *exit*, Java states how to *stay*.

## 6. `for-range`: the loop you'll use most (`06-for-range/`)

Plain English: `for i, v := range stints` walks a slice, map, or string and
hands you position + copy-of-value each lap. It is the default choice for
"visit everything".

Real world, three walks from the program:

- Stint times `[58 61 59]` — index plus value, or `_` to blank the index
  when only times matter.
- Start grid `{"ada", "bo", "cy"}` used as a set — `for name := range grid`
  visits keys only.
- The string `"pit_π!"` — range yields the *byte offset* and the *rune*,
  decoding the two-byte `π` (offset jumps 4 → 6). Indexing `s[i]` would
  hand back raw bytes instead.

```go
for i, r := range "pit_π!" {
	fmt.Println(i, r, string(r))
}
// 0 112 p
// 1 105 i
// 2 116 t
// 3 95 _
// 4 960 π
// 6 33 !
```

Three rules that bite beginners:

1. **The value is a copy.** `for _, v := range stints { v *= 2 }` doubles
   the copy and the slice prints `unchanged: [58 61 59]`. Index-assign
   (`stints[i] *= 2`) when you mean to edit.
2. **Map order is random on purpose** (per-map hash seed plus jitter, as
   anti-Hash-DoS armor). Sort keys for stable output, or rely on the one
   exception: `fmt` printing a map sorts keys for you.
3. **Fresh variables since Go 1.22.** Each iteration now gets its own `i`
   and `v` (selected by the `go` line in `go.mod`), so goroutines started
   in the loop stop sharing one reused pair.

## 7. Labels: steering the outer loop (`07-for-labels/`)

Plain English: bare `break`/`continue` steer only the innermost loop. Put a
name on the outer loop (`outer:`) and `continue outer` / `break laps`
steers that one from deep inside.

Real world: heats `q1, q2` crossed with lanes `1, 2, 3`. A false start in
`q1` lane 2 abandons the whole heat (`continue outer`), so `q1 complete`
never prints but `q2` runs fully. A red flag on lap 3 stops the entire
session (`break laps`).

```go
outer:
	for _, heat := range heats {
		for _, lane := range lanes {
			if heat == "q1" && lane == 2 {
				continue outer // skip the rest of this heat entirely
			}
		}
		fmt.Println(heat, "complete: all lanes ran")
	}
```

Style note: `gofmt` indents labels to brace level so they stand out. Labels
are rare — reaching for one usually means "skip/stop the whole outer job",
which is exactly when they earn their keep.

## 8. `switch`: matching without falling through (`08-switch-cases/`)

Plain English: unlike C, Go cases never fall through — no `break` needed,
`fallthrough` is opt-in and best avoided. Group equal outcomes with commas;
leave a case body empty to mean "deliberately nothing".

Real world: sorting driver callsigns by length, with mid-length names
deliberately silent:

```go
switch size := len(name); size { // size lives for every branch
case 1, 2, 3:
	fmt.Println(name, "is short") // al, bea, rex
case 4:
	fmt.Println(name, "fits exactly:", size)
case 5, 6, 7:
	// empty on purpose: pitlane (7) prints nothing
default:
	fmt.Println(name, "is long") // chronometer
}
```

The **blank switch** (`switch { case cond: }`) runs boolean tests instead
of equality — same lap-classification as section 5, but the `default`
branch makes "green" explicit and no `continue` chain is needed:

```go
switch {
case lap%2 == 0 && lap%3 == 0:
	fmt.Println(lap, "pit-charge")
case lap%2 == 0:
	fmt.Println(lap, "pit")
// ...
default:
	fmt.Println(lap, "green")
}
```

Two label lessons: `break` inside a case exits the *switch*, so label the
loop (`break laps`) to exit the *race*. And choose by relationship: related
checks on one idea → `switch`; unrelated tests → `if/else`.

## 9. `goto`: the emergency exit (`09-goto/`)

Plain English: `goto done` jumps forward to a labeled line in the same
function. It cannot skip a declaration or dive into a block — the compiler
rejects both — which keeps it honest.

Real world: speed starts at 12 and doubles-plus-one until 100, but any
multiple of 5 (first: 25) jumps to the shared wind-down instead of
finishing normally:

```go
speed := 12
for speed < 100 {
	if speed%5 == 0 {
		goto done
	}
	speed = speed*2 + 1
}
fmt.Println("loop finished on its own")

done:
	fmt.Println("wind down at speed", speed) // wind down at speed 25
```

Reserve it for shared exit/cleanup where a boolean flag litters the logic
and duplicating the cleanup invites drift. Labeled `break`/`continue` and
early returns cover everything else — the real `strconv` package uses
`goto` exactly this way for its overflow/out paths.

## Rules I want to remember

- Inner sees outer; outer never sees inner. `undefined: x` means the room is gone.
- `:=` declares here; `=` updates. Mixed `a, b := ...` still declares if any name is new.
- Never reuse `fmt`, `len`, `true` as locals — the original goes dark for the block.
- Scope temporaries to the decision: `if n := ...; ...` and `switch s := ...; ...`.
- One loop, four shapes: complete, condition-only, infinite, range. Range wins for full walks.
- Range gives copies; map order shuffles; strings yield runes with byte offsets.
- `break`/`continue` hit the inner loop; labels reach the outer. `switch` never falls through.
- `goto` jumps forward only, never over declarations or into blocks — and almost never at all.
