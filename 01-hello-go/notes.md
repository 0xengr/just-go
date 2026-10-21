# 01 — Setting Up Your Go Environment

The big idea: a Go program is just a package named `main` with a `func main()`,
and the toolchain (`go run`, `go build`, `go vet`, `gofmt`) is your whole
workflow — there is no separate build system to learn first. A module
(`go.mod`) gives your code an identity and an import path, and everything
else in this chapter is about meeting that toolchain and printing your first
"Hello".

## How to run

- `go build ./...` type-checks the overview and all subexamples at once.
- `go run .` is only an orientation (it prints `Hello from Golang!`); every
  topic is its own program.
- `make list` lists the topics; `make check` runs the chapter gate
  (`gofmt` + `vet` + `build`).
- `make run 01-greet Ada` (or `make run 02-toolchain`) runs one topic with a
  banner; add `ANIM=0` to skip the banner.
- Suggested order for the first pass:

| Order | Run it | What it teaches you |
| ----- | ------ | ------------------- |
| 1 | `go run ./examples/01-greet Ada` | A first program with a command-line argument |
| 2 | `go run ./examples/02-toolchain` | What `go env` and `go version` actually report |

Topics are numbered folders (`01-`, `02-`, …) in reading order, so `ls` alone
tells a newcomer where to begin. The numbers are the curriculum, not a version.

## 1. A first program with an argument (`01-greet/`)

Plain English: read the command line, fall back to a default name when nothing
was given, and print a greeting. `os.Args` is the slice of strings the OS
handed your program — index 0 is the program's own name, so the first real
argument lives at index 1.

Real world: a receptionist who greets whoever walks in by name, and says
"Hello, Gopher!" when nobody hands over a badge.

```go
name := "Gopher"
if len(os.Args) > 1 && os.Args[1] != "" {
    name = os.Args[1]
}
fmt.Printf("Hello, %s!\n", name)
```

Output:

```text
$ go run ./examples/01-greet Ada
Hello, Ada!

$ go run ./examples/01-greet
Hello, Gopher!
```

Beginner trap: reaching for `os.Args[0]` thinking it is the first argument —
it is the program name (often `./examples/01-greet`), so your greeting comes out as
"Hello, ./examples/01-greet!". Also note the empty-string guard: `go run ./examples/01-greet ""`
keeps the default instead of printing "Hello, !".

## 2. What the toolchain reports (`02-toolchain/`)

Plain English: ask the runtime what it is — toolchain version, operating
system, architecture, and how many logical CPUs are visible. These values are
baked into the binary at build time, so the same source reports whatever
target you compiled for.

Real world: reading the nameplate bolted to a machine before you trust what
it can do.

```go
fmt.Println("go version:", runtime.Version())
fmt.Println("os/arch:  ", runtime.GOOS+"/"+runtime.GOARCH)
fmt.Println("compilers:", runtime.NumCPU(), "logical CPUs visible")
```

Output:

```text
go version: go1.27.2
os/arch:   linux/amd64
compilers: 4 logical CPUs visible
```

Beginner trap: confusing the `go version` shell command with
`runtime.Version()` — they usually agree, but the latter is fixed when you
*build*, not when you run. Cross-compile and the program still reports the
target: `GOOS=windows GOARCH=amd64 go build ./...`. Also, `NumCPU()` counts
logical (hyperthreaded) CPUs visible to the process, not physical cores.

## Notes

- A module is just a directory with a `go.mod` naming it; the toolchain uses
  that path as the import prefix. No build config beyond that for a small
  program.
- `go run` compiles to a temp dir (nothing saved); `go build` leaves a
  binary you can ship.
- `gofmt` is the style authority — run it instead of debating formatting.
- `go vet` catches real mistakes (bad format verbs, unreachable code) that
  compile fine, so run it before assuming code is correct.
- The compatibility promise means code written today should still build on
  newer Go; upgrade the toolchain freely, pin the `go` line in `go.mod` per
  module.
- Editor choice doesn't matter much: any editor + terminal with the commands
  above is enough; the playground is handy for sharing snippets.

## Exercises

1. **Greet everyone.** Change `01-greet` so it accepts *any* number of names
   on the command line and prints them all, e.g. `go run ./examples/01-greet Ada Bo`
   prints `Hello, Ada and Bo!`. With no arguments, keep the default `Gopher`.
   - Hint: slice `os.Args[1:]` is already the list of names; `strings.Join`
     glues them with `" and "`.
   <details>
   <summary>Solution</summary>

   ```go
   package main

   import (
       "fmt"
       "os"
       "strings"
   )

   func main() {
       names := os.Args[1:]
       if len(names) == 0 {
           names = []string{"Gopher"}
       }
       fmt.Printf("Hello, %s!\n", strings.Join(names, " and "))
   }
   ```

   </details>

2. **Target check.** Write a tiny program (in a new folder or in your head)
   that prints the toolchain version and `GOOS/GOARCH`, then prints
   `"64-bit target"` only when `GOARCH` ends in `"64"`, otherwise
   `"32-bit target"`.
   - Hint: `strings.HasSuffix(runtime.GOARCH, "64")`.
   <details>
   <summary>Solution</summary>

   ```go
   package main

   import (
       "fmt"
       "runtime"
       "strings"
   )

   func main() {
       fmt.Println(runtime.Version())
       fmt.Println(runtime.GOOS + "/" + runtime.GOARCH)
       if strings.HasSuffix(runtime.GOARCH, "64") {
           fmt.Println("64-bit target")
       } else {
           fmt.Println("32-bit target")
       }
   }
   ```

   </details>

3. **Format verbs.** Print a name and its length with one `fmt.Printf` call,
   using the correct verbs, so the output is exactly `name=Gopher len=6`.
   - Hint: `%s` for a string, `%d` for an integer.
   <details>
   <summary>Solution</summary>

   ```go
   package main

   import "fmt"

   func main() {
       name := "Gopher"
       fmt.Printf("name=%s len=%d\n", name, len(name))
   }
   ```

   </details>

## Rules I want to remember

- A Go program starts at `package main`'s `func main()`; everything else is a
  library.
- One module = one directory with a `go.mod`; the module path is the import
  prefix.
- `go run` compiles to a temp dir and runs; `go build` leaves a binary you
  can ship; `go build ./...` checks the whole chapter.
- `gofmt` is the style authority; `go vet` catches mistakes that still compile.
- `os.Args[0]` is the program name; real arguments start at index 1.
- `:=` declares a new variable; `=` updates an existing one.
- `runtime.GOOS`/`GOARCH` are fixed at build time — cross-compile with
  `GOOS=... GOARCH=... go build`.
- Go's compatibility promise: old code keeps building; pin the `go` line in
  `go.mod` per module.
- Editor doesn't matter: any editor + terminal with the commands above is
  enough; the playground is handy for sharing snippets.