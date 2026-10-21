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

Topics are numbered folders under `examples/` (`examples/01-`, `examples/02-`, …) in reading order, so `ls examples/` alone
tells a newcomer where to begin. The numbers are the curriculum, not a version.

## Notes

Full chapter notes live in [notes.md](./notes.md).

## Exercises

- `exercises/01-name-badge` — Topics solved: 01-greet, 02-toolchain
- `exercises/02-count-args` — Topics solved: 01-greet
