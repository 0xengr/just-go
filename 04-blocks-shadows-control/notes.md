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
| 1 | `go run ./examples/01-blocks` | Blocks nest; inner sees outer, not the reverse |
| 2 | `go run ./examples/02-shadowing` | `:=` can silently create a lookalike variable |
| 3 | `go run ./examples/03-if-scope` | `if x := ...; cond` keeps temporaries local |
| 4 | `go run ./examples/04-for-complete` | The classic `for i := 0; i < n; i++` and its variations |
| 5 | `go run ./examples/05-for-while` | The `while` shape, infinite loops, `break`/`continue` |
| 6 | `go run ./examples/06-for-range` | The everyday loop over slices, maps, and strings |
| 7 | `go run ./examples/07-for-labels` | Steering an *outer* loop from an inner one |
| 8 | `go run ./examples/08-switch-cases` | Matching values, blank switches, no fallthrough |
| 9 | `go run ./examples/09-goto` | The one rare case where jumping forward reads better |

Topics are numbered folders (`01-`, `02-`, …) in reading order, so `ls` alone
tells a newcomer where to begin. The numbers are the curriculum, not a version.
Here the numbering simply encodes the table above.

## 1. Blocks are rooms with one-way glass (`01-blocks/`)

Plain English: every pair of braces `{}` is a room. Code inside a room can
read the whiteboard of every room outside it, but nobody outside can read
what was written inside once the door closes. The package block (the whole
file) and the function block (`main`) are rooms too, just bigger ones.

Real world: the season name (`season = "spring series"`) is pinned on the
building wall — every function sees it. The lap count (`laps := 3`) lives
in the meeting room called `main` — every block inside `main` sees it.
The best-lap time (`best := 58.4`) is scribbled inside a tiny huddle room
(a bare `{}` block) and thrown away when the huddle ends. Even the `for`
loop gets its own room: `i` lives only for that statement.

```go
var season = "spring series" // package block: whole file sees it

func main() {
	laps := 3 // function block: visible everywhere below in main

	if laps > 0 {
		fmt.Println(season, "laps:", laps) // inner reads outer: fine
		{
			best := 58.4 // huddle room: dies at the closing brace
			fmt.Println("best this stint:", best, "over", laps, "laps")
		}
		// BUG: best is gone once its block ends:
		// fmt.Println(best)
	}

	for i := 0; i < 1; i++ {
		fmt.Println("loop i:", i, "still sees laps:", laps)
	}
	// BUG: i is scoped to the for statement:
	// fmt.Println(i)
}
```

Output:

```text
spring series laps: 3
best this stint: 58.4 over 3 laps
loop i: 0 still sees laps: 3
```

Beginner trap: uncommenting either commented-out line fails with
`undefined: best` / `undefined: i`. That error is the compiler telling you
the room has been demolished. When a temporary is only needed for three
lines, a bare block (or a tighter `if` scope) keeps it from leaking into
the rest of the function.

## 2. Shadowing: the evil twin (`02-shadowing/`)

Plain English: if you declare a new variable with the same name as an outer
one, you get two variables wearing the same jersey. Inside the inner block
you only ever see the twin; the outer one is still there, just unreachable.
The twin dies with its block and the original was never touched.

Real world: race control has a fuel board showing `fuel := 10`. A pit crew
member writes `fuel := 5` on their own clipboard inside the pit box. The
main board still says 10 — the crew just can't see it while holding their
clipboard.

```go
fuel := 10
if fuel > 5 {
	fmt.Println("before:", fuel)
	fuel := 5 // a NEW fuel, alive only inside this if body
	fmt.Println("inside:", fuel)
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
`pit` is new, the whole line declares, and `fuel` shadows again:

```go
if fuel > 5 {
	fuel, pit := 5, 20 // fuel shadows again, pit is brand new
	fmt.Println("shadowed pair:", fuel, pit)
}
fmt.Println("outer still:", fuel)
```

Output:

```text
shadowed pair: 5 20
outer still: 10
```

The fix is boring on purpose: use plain `=` when the variable already exists.

```go
fuel = 5 // assignment, not declaration: the real board changes
fmt.Println("reassigned:", fuel)
```

Output:

```text
reassigned: 5
```

Beginner trap: never name a local `fmt`, `len`, or `true`. They are not
keywords — they live in the universe/file blocks — so your local silently
hides the package or builtin for the rest of the block, and `fmt.Println`
suddenly means "field Println on my string". `go vet` won't save you here;
a shadowing linter will. The same trap applies to imports.

## 3. `if` with its own private clipboard (`03-if-scope/`)

Plain English: `if x := ...; condition` creates `x` just for that one
`if/else` chain. It is visible in every branch and gone afterwards. It is
the cleanest way to keep a temporary next to the decision that needs it.

Real world: measuring a driver's callsign to decide which board it fits on.
The length `n` matters for exactly one decision — after that, keeping it
around is clutter.

```go
name := "bo"

if n := len(name); n == 0 {
	fmt.Println("no callsign entered")
} else if n > 5 {
	fmt.Println(name, "is a long callsign:", n)
} else {
	fmt.Println(name, "fits the board:", n) // bo fits the board: 2
}
// BUG: n lived only for that decision:
// fmt.Println(n)
```

Output:

```text
bo fits the board: 2
```

The short statement shadows too. `if cutoff := 4; ...` inside a function
that already has `cutoff := 6` creates a lookalike for one branch, and the
outer value is untouched:

```go
cutoff := 6
if cutoff := 4; cutoff > len(name) {
	fmt.Println("local cutoff", cutoff, "beats outer", 6)
}
fmt.Println("outer cutoff:", cutoff)
```

Output:

```text
local cutoff 4 beats outer 6
outer cutoff: 6
```

Beginner trap: use the slot only for fresh temporaries, never to smuggle in
side-effect calls or assignments to existing variables. `if x := f(); x > 0`
is fine; `if x = f(); ...` is not even legal there, and `if n := len(name);`
followed by mutating `name` in a branch does not reach back out.

## 4. The complete `for`: start, check, step (`04-for-complete/`)

Plain English: `for i := 0; i < 5; i++` is three instructions — set up the
counter, check before every lap, step after every lap. The setup must use
`:=`; the check must be a boolean. Any of the three clauses may be dropped.

Real world: sending cars out for laps 0 through 4, then reusing the pattern
for a slice window (positions 1 to 3 of `[9 7 8 6 10]`):

```go
for i := 0; i < 5; i++ {
	fmt.Print(i, " ") // 0 1 2 3 4
}
fmt.Println()

scores := []int{9, 7, 8, 6, 10}
for k := 1; k < len(scores)-1; k++ {
	fmt.Print(scores[k], " ") // 7 8 6
}
fmt.Println()
```

Output:

```text
0 1 2 3 4 
7 8 6 
```

Hoist the setup when it comes from earlier work, leaving the init clause
empty:

```go
i := 0
for ; i < 3; i++ {
	fmt.Print(i, " ")
}
fmt.Println() // 0 1 2
```

Drop the post when the stride needs logic inside the loop:

```go
for j := 0; j < 10; {
	fmt.Print(j, " ")
	if j%2 == 0 {
		j++
	} else {
		j += 2
	}
}
fmt.Println() // 0 1 3 5 7 9
```

Output:

```text
0 1 2 
0 1 3 5 7 9 
```

Rule of thumb: reaching for `break`/`continue` just to walk a sub-window
means a plain `for` with honest bounds would read better. The bounds state
the intent; the loop body stays about the work.

Beginner trap: the init clause requires `:=` (or a prior declaration with
an empty init). Writing `for i = 0; ...` without declaring `i` first is a
compile error, and `for i := 0; i < n; i++` declares `i` only for the
statement — it is gone after the loop, same as the `for` block rule in
section 1.

## 5. `for` as `while`, `for` as forever (`05-for-while/`)

Plain English: Go has no `while` keyword — you just leave pieces off. Drop
setup and step (and the semicolons) and `for speed < 8` *is* while. Drop
the condition too and `for {}` runs until something breaks it.

Real world: doubling speed `1 2 4` until it hits the limit; retrying a
connection up to 3 attempts; classifying 12 laps as pit / charge /
pit-charge / green without nesting:

```go
speed := 1
for speed < 8 { // the while shape
	fmt.Print(speed, " ")
	speed *= 2
}
fmt.Println() // 1 2 4
```

Output:

```text
1 2 4 
```

The infinite loop carries its own exit:

```go
attempts := 0
for { // forever, until the explicit exit
	attempts++
	if attempts >= 3 {
		break
	}
}
fmt.Println("attempts:", attempts) // attempts: 3
```

Output:

```text
attempts: 3
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
	if lap%3 == 0 {
		fmt.Println(lap, "charge")
		continue
	}
	fmt.Println(lap, "green")
}
```

Output:

```text
1 green
2 pit
3 charge
4 pit
5 green
6 pit-charge
7 green
8 pit
9 charge
10 pit
11 green
12 pit-charge
```

Need do-while ("run at least once")? Write
`for { work; if !cond { break } }` — note the flipped condition versus
Java: Go states how to *exit*, Java states how to *stay*.

Beginner trap: an infinite `for` needs a `break` or `return` in its body,
or the program never finishes. The compiler will not warn you — it happily
builds a program that spins forever.

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
stints := []int{58, 61, 59}
for i, v := range stints {
	fmt.Println(i, v)
}
```

Output:

```text
0 58
1 61
2 59
```

Blank the index when only values matter; drop the value for key-only loops
(most common when a map is used as a set):

```go
for _, v := range stints {
	fmt.Print(v, " ")
}
fmt.Println()

grid := map[string]bool{"ada": true, "bo": true, "cy": true}
for name := range grid {
	fmt.Print(name, " ")
}
fmt.Println()
```

Output:

```text
58 61 59 
ada bo cy 
```

Range over the string yields byte offset + rune:

```go
for i, r := range "pit_π!" {
	fmt.Println(i, r, string(r))
}
```

Output:

```text
0 112 p
1 105 i
2 116 t
3 95 _
4 960 π
6 33 !
```

Sorted keys for stable output (map order is random on purpose):

```go
keys := make([]string, 0, len(grid))
for k := range grid {
	keys = append(keys, k)
}
sort.Strings(keys)
for _, k := range keys {
	fmt.Print(k, " ")
}
fmt.Println() // ada bo cy
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

Beginner trap: writing `for i, v := range stints { stints[i] = v * 2 }`
is correct but easy to get backwards; the more common bug is mutating `v`
and expecting the slice to change. Also, ranging a string gives *runes*, so
`len("π")` is 2 (bytes) while `range` sees one iteration — mix the two
and your indexing drifts.

## 7. Labels: steering the outer loop (`07-for-labels/`)

Plain English: bare `break`/`continue` steer only the innermost loop. Put a
name on the outer loop (`outer:`) and `continue outer` / `break laps`
steers that one from deep inside. The label must sit on a `for`, `switch`,
or `select`.

Real world: heats `q1, q2` crossed with lanes `1, 2, 3`. A false start in
`q1` lane 2 abandons the whole heat (`continue outer`), so `q1 complete`
never prints but `q2` runs fully. A red flag on lap 3 stops the entire
session (`break laps`).

```go
outer:
	for _, heat := range heats {
		for _, lane := range lanes {
			if heat == "q1" && lane == 2 {
				fmt.Println(heat, "lane", lane, "false start: skipping heat")
				continue outer // skip the rest of this heat entirely
			}
			fmt.Println(heat, "lane", lane, "clean")
		}
		fmt.Println(heat, "complete: all lanes ran")
	}
```

Output:

```text
q1 lane 1 clean
q1 lane 2 false start: skipping heat
q2 lane 1 clean
q2 lane 2 clean
q2 lane 3 clean
q2 complete: all lanes ran
```

The other direction — break out of everything at once:

```go
laps:
	for lap := 1; lap <= 5; lap++ {
		for stall := 1; stall <= 5; stall++ {
			if lap == 3 && stall == 2 {
				fmt.Println("red flag on lap", lap)
				break laps
			}
		}
	}
	fmt.Println("session halted")
```

Output:

```text
red flag on lap 3
session halted
```

Style note: `gofmt` indents labels to brace level so they stand out. Labels
are rare — reaching for one usually means "skip/stop the whole outer job",
which is exactly when they earn their keep. If you need one for a single
loop, you almost certainly want a function or an early return instead.

Beginner trap: a label with no matching `break`/`continue`/`goto` is a
compile error, and `break foo` where `foo` labels a `switch` exits the
switch, not any loop around it. Name loops and switches distinctly to
avoid misreading which one a `break` leaves.

## 8. `switch`: matching without falling through (`08-switch-cases/`)

Plain English: unlike C, Go cases never fall through — no `break` needed,
`fallthrough` is opt-in and best avoided. Group equal outcomes with commas;
leave a case body empty to mean "deliberately nothing". A `switch` can
also declare a variable scoped to every branch, just like `if`.

Real world: sorting driver callsigns by length, with mid-length names
deliberately silent:

```go
callsigns := []string{"al", "bea", "rex", "pitlane", "chronometer"}

for _, name := range callsigns {
	switch size := len(name); size {
	case 1, 2, 3:
		fmt.Println(name, "is short") // al, bea, rex
	case 4:
		fmt.Println(name, "fits exactly:", size)
	case 5, 6, 7:
		// empty on purpose: pitlane (7) prints nothing
	default:
		fmt.Println(name, "is long") // chronometer
	}
}
```

Output:

```text
al is short
bea is short
rex is short
chronometer is long
```

The **blank switch** (`switch { case cond: }`) runs boolean tests instead
of equality — same lap-classification as section 5, but the `default`
branch makes "green" explicit and no `continue` chain is needed:

```go
for lap := 1; lap <= 15; lap++ {
	switch {
	case lap%2 == 0 && lap%3 == 0:
		fmt.Println(lap, "pit-charge")
	case lap%2 == 0:
		fmt.Println(lap, "pit")
	case lap%3 == 0:
		fmt.Println(lap, "charge")
	default:
		fmt.Println(lap, "green")
	}
}
```

Output:

```text
1 green
2 pit
3 charge
4 pit
5 green
6 pit-charge
7 green
8 pit
9 charge
10 pit
11 green
12 pit-charge
13 green
14 pit
15 charge
```

Two label lessons: `break` inside a case exits the *switch*, so label the
loop (`break laps`) to exit the *race*. And choose by relationship: related
checks on one idea → `switch`; unrelated tests → `if/else`.

```go
laps:
	for lap := 1; lap <= 10; lap++ {
		switch lap {
		case 2, 4, 6:
			fmt.Println(lap, "even")
		case 7:
			fmt.Println("red flag: leaving the session")
			break laps // leaves the loop, not just the switch
		default:
			fmt.Println(lap, "running")
		}
	}
```

Output:

```text
1 running
2 even
3 running
4 even
5 running
6 even
red flag: leaving the session
```

Beginner trap: expecting C-style fallthrough. `case 1:` followed by
`case 2:` does *not* run both; if you genuinely need fallthrough, the
`fallthrough` keyword is there, but restructuring the cases (group with
commas, or move shared code into a helper) almost always reads better.
Also, an empty case is legal and meaningful — `case 5, 6, 7:` with no
body means "do nothing on purpose"; do not add a `break` "just in case".

## 9. `goto`: the emergency exit (`09-goto/`)

Plain English: `goto done` jumps forward to a labeled line in the same
function. It cannot skip a declaration or dive into a block — the compiler
rejects both — which keeps it honest. It jumps only forward in practice;
the language allows backward jumps but style forbids them.

Real world: speed starts at 12 and doubles-plus-one until 100, but any
multiple of 5 (first: 25) jumps to the shared wind-down instead of
finishing normally:

```go
speed := 12
for speed < 100 {
	if speed%5 == 0 {
		goto done // both exits below need the same wind-down
	}
	speed = speed*2 + 1
}
fmt.Println("loop finished on its own")

done:
	fmt.Println("wind down at speed", speed) // wind down at speed 25
```

Output:

```text
wind down at speed 25
```

Reserve it for shared exit/cleanup where a boolean flag litters the logic
and duplicating the cleanup invites drift. Labeled `break`/`continue` and
early returns cover everything else — the real `strconv` package uses
`goto` exactly this way for its overflow/out paths.

The compiler keeps it honest:

```go
// BUG: both of these jumps are rejected, so they stay commented out.
// A goto may never skip a declaration:
//   speed := 1
//   goto skip
//   extra := 2  // rejects: jumps over declaration of extra
// skip:
//   _ = extra
// Nor may it dive into a block:
//   if speed > 0 {
//       goto inner // rejects: jumps into block
//   }
//   if speed < 100 {
//   inner:
//       fmt.Println("unreachable legally")
//   }
```

Beginner trap: `goto` is the tool of last resort. If you reach for it to
escape a nested loop, a labeled `break` is clearer; to skip cleanup, a
`defer` or early return is clearer. The one good use is a single shared
exit point with cleanup that several paths need — and even then, ask
whether a helper function reads better.

## Exercises

1. **Scope probe.** In `01-blocks`, uncomment `fmt.Println(best)` inside
   the `if` but outside the bare block. Predict the exact compiler error,
   then uncomment `fmt.Println(i)` after the `for`. Run `go vet`/build and
   confirm both messages. Why does the `for`'s `i` vanish even though the
   loop body is in the same function?
2. **Fix the shadow.** Rewrite the first `if` in `02-shadowing` so the
   outer `fuel` becomes 5, using `=` instead of `:=`. Then change
   `fuel, pit := 5, 20` to reuse `fuel` while still declaring `pit`. What
   is the smallest edit that keeps `pit` new and `fuel` updated?
3. **Callsign board.** Using `03-if-scope`'s shape, write an `if n := len(name);`
   that prints "empty", "short" (1–3), or "long" (4+). Make `name := ""`
   and confirm `n` is unavailable after the chain. Bonus: nest a shadowing
   `cutoff := 2` inside one branch and print the outer afterward.
4. **Stride and window.** With `scores := []int{9, 7, 8, 6, 10}`, print
   the elements at even indices only, then print the slice reversed using a
   complete `for` with a descending post. Do not use `break` or `continue`.
5. **Rune walk.** Range over `"Go_π!"` and print only the byte offsets
   where a new rune starts, plus the rune. Then predict: how many
   iterations does `range` perform, and why is `len("Go_π!")` not the
   answer? Confirm by printing `len`.
## Rules I want to remember

- Inner sees outer; outer never sees inner. `undefined: x` means the room is gone.
- `:=` declares here; `=` updates. Mixed `a, b := ...` still declares if any name is new.
- Never reuse `fmt`, `len`, `true` as locals — the original goes dark for the block.
- Scope temporaries to the decision: `if n := ...; ...` and `switch s := ...; ...`.
- One loop, four shapes: complete, condition-only, infinite, range. Range wins for full walks.
- Range gives copies; map order shuffles; strings yield runes with byte offsets.
- `break`/`continue` hit the inner loop; labels reach the outer. `switch` never falls through.
- `goto` jumps forward only, never over declarations or into blocks — and almost never at all.