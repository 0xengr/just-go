# Just Go

[![check](https://github.com/0xengr/just-go/actions/workflows/check.yml/badge.svg)](https://github.com/0xengr/just-go/actions/workflows/check.yml)
![Go](https://img.shields.io/badge/Go-1.27-00ADD8?logo=go)
![License](https://img.shields.io/badge/license-MIT-green)

Learning Go by writing it. Every topic here is a small runnable program, and
each chapter ends with notes in my own words — the bits I want to remember when
I come back to the code.

Chapters follow the order of *Learning Go* by Jon Bodner (2nd edition), one
chapter per folder, so each `NN-slug/` can be read on its own.

> **Status:** this repo is a complete beginner guide covering chapters 1–7 of
> the book (environment, predeclared types, composite types, blocks/control,
> functions, pointers, types & methods). Chapters 8–16 (interfaces, generics,
> errors, modules, tooling, concurrency, stdlib, context, testing,
> reflect/unsafe/cgo) are **not** covered here — read them in the book or add
> them yourself. Fixes and typos are welcome; see "Contributing" below.

## Who this is for

You are new to Go, or new to a statically typed, compiled language, and want to
learn by running small programs instead of reading a textbook cover to cover.
You should already be comfortable with at least one programming language (loops,
functions, basic types). No prior Go knowledge is assumed.

## Prerequisites

- Go 1.27 or newer (`go version`). Install from <https://go.dev/dl/>.
- `make` (any POSIX make; on Debian/Ubuntu: `sudo apt install make`).
- `git` (to clone and to switch chapters).
- A terminal and any text editor. GoLand/VS Code are optional, not required.

## Get started in 5 minutes

```sh
git clone https://github.com/0xengr/just-go.git
cd just-go
make list                      # every chapter and every topic inside it
make check                     # gofmt + vet + build across all chapters
make -C 01-hello-go run 01-greet Ada   # run one topic, with the banner
```

Then open `01-hello-go/README.md`: it opens with the big idea, a run-order
table, and one section per topic with the code and its real output.

## How to read a chapter

Each chapter is a self-contained Go module:

```
NN-slug/
├── go.mod          module path for this chapter
├── main.go         orientation stub (overview only)
├── Makefile        same in every chapter, auto-discovers topics
├── README.md       big idea, run order, notes, traps, cheat sheet, exercises
└── 01-topic/       a runnable package main
    └── main.go
```

- Topics are numbered folders (`01-`, `02-`, …) in reading order, so `ls`
  alone tells a newcomer where to begin. The numbers are the curriculum, not a
  version.
- Every topic is its own `package main`: run it with `go run ./01-topic` or
  `make -C NN-slug run 01-topic`.
- The chapter `README.md` is the notes: read it after running the code.
- Exercises at the end of each README are for you to try before peeking at the
  collapsible solution.

## How to run examples

```sh
make -C 03-composite-types run 02-slices          # with animated banner
make -C 03-composite-types run 02-slices ANIM=0   # skip the banner
make -C 03-composite-types list                   # list topics
make -C 03-composite-types check                  # gofmt + vet + build
make -C 03-composite-types new my-topic           # scaffold a new topic
```

At the repo root, `make list` shows every chapter's topics and `make check`
runs the gate across all of them; `make -C <chapter> <target>` works too.

Reference notes: `fmt-printing.md` — all 12 `fmt` print functions + verbs.

## Layout

43 runnable topics across 7 chapters. `Start here` runs the first topic in
reading order; the chapter `README.md` has the full run-order table.

| Dir                          | Chapter                              | Topics | Start here                                          | You'll learn |
| ---------------------------- | ------------------------------------ | -----: | --------------------------------------------------- | ------------ |
| `01-hello-go/`               | Go environment & toolchain           |    2 | `make -C 01-hello-go run 01-greet Ada`              | modules, go run/build/fmt/vet |
| `02-predeclared-types/`      | Predeclared types & declarations     |    6 | `make -C 02-predeclared-types run 01-zero-values`   | zero values, var/:=, const, numerics, strings, conversions |
| `03-composite-types/`        | Arrays, slices, maps & structs       |   11 | `make -C 03-composite-types run 01-arrays`          | arrays, slices, maps, structs, strings/bytes |
| `04-blocks-shadows-control/` | Blocks, shadows & control structures |    9 | `make -C 04-blocks-shadows-control run 01-blocks`   | scope, shadowing, if/for/switch/goto |
| `05-functions/`              | Functions, closures & `defer`        |    5 | `make -C 05-functions run 01-multiple-returns`      | multiple returns, variadic, closures, defer |
| `06-pointers/`               | Pointers, nil, `new()` & call by value |  6 | `make -C 06-pointers run 01-call-by-value`          | pointers, mutability, map/slice diffs, perf |
| `07-types-and-methods/`      | Types and methods (types + methods)  |  4 | `make -C 07-types-and-methods run 01-user-defined-types` | defined types, methods, receivers, embedding |

Each chapter is its own Go module (`go.mod` naming `github.com/0xengr/just-go/<slug>`)
with its own overview `main.go` and `README.md`; there is no root module and no
`go.work`. Chapter 1 is the exception: its overview `main.go` is a runnable hello
world, the rest are orientation stubs.

Topic folders are numbered in reading order (`01-`, `02-`, …) within each
chapter, so `ls` on any chapter shows a newcomer where to begin. The numbers are
the curriculum, not a version.

## What's NOT covered

This repo stops at chapter 7 of *Learning Go*. The book continues with:

| Book chapter | Topic | Where to learn it |
| ------------ | ----- | ----------------- |
| 8 | Generics | *Learning Go* ch. 8 |
| 9 | Errors (sentinels, wrapping, panic/recover) | *Learning Go* ch. 9 |
| 10 | Modules, packages, imports | *Learning Go* ch. 10 |
| 11 | Go tooling (goimports, linters, build tags, embed) | *Learning Go* ch. 11 |
| 12 | Concurrency (goroutines, channels, select) | *Learning Go* ch. 12 |
| 13 | Standard library (io, time, json, http) | *Learning Go* ch. 13 |
| 14 | Context | *Learning Go* ch. 14 |
| 15 | Writing tests | *Learning Go* ch. 15 |
| 16 | reflect, unsafe, cgo | *Learning Go* ch. 16 |

If you want these in the repo, they are great follow-up PRs — keep the same
chapter layout and Makefile.

## Make

Every chapter has the same `Makefile`; it finds the chapter name and its topics
by itself, so there is nothing to configure.

| Command                   | What it does                                                |
| ------------------------- | ----------------------------------------------------------- |
| `make` / `make help`      | list the targets                                            |
| `make list`               | list this chapter's topics                                  |
| `make run <topic> [args]` | `go run ./<topic> [args]`, with a banner (`ANIM=0` to skip) |
| `make new <topic>`        | scaffold `<topic>/main.go`                                  |
| `make test`               | `go test ./...`                                             |
| `make check`              | `gofmt` + `vet` + `build` for the chapter                   |

Run them from a chapter folder, e.g. `cd 05-functions && make run 01-multiple-returns`.
At the repo root, `make list` shows every chapter's topics and `make check` runs
the gate across all of them; `make -C <chapter> <target>` works too. Add
`ANIM=0` to skip the banner, e.g. `make run 01-greet Ada ANIM=0`.

Reference notes: `fmt-printing.md` — all 12 `fmt` print functions + verbs.

## Why Go

- Small language, few ways to do things: easy to read other people's code.
- Fast static binaries with no runtime to install: great for CLIs and services.
- Built-in concurrency (`goroutines` + `channels`) without extra libraries.
- Strong stdlib (HTTP, JSON, testing, tooling) so projects stay lean.
- Backward-compatibility promise: old code keeps building on new toolchains.

## Go features at a glance

- Static typing with inference (`:=`), garbage collected, compiles to one binary.
- Explicit error returns instead of exceptions; `defer` for cleanup.
- Interfaces are satisfied implicitly (no `implements` keyword).
- `struct` + embedding for composition instead of inheritance.
- Generics (since 1.18) for type-safe containers/algorithms without `any` sprawl.
- Cross-compilation via `GOOS`/`GOARCH` env vars.

## Tooling reference

| Command                        | What it does                                  | Example                            |
| ------------------------------ | --------------------------------------------- | ---------------------------------- |
| `go run .` / `go run ./<dir>`  | Compile to temp dir and run                   | `go run .`, `go run ./01-greet`    |
| `go build ./...`               | Compile all packages, check for errors        | `go build ./...`                   |
| `go vet ./...`                 | Static checks (printf verbs, dead code, etc.) | `go vet ./...`                     |
| `gofmt -l .`                   | List files needing format (empty = clean)     | `gofmt -l .`, `gofmt -w .`         |
| `go mod init <path>`           | Start a new module                            | `go mod init example.com/demo`     |
| `go mod tidy`                  | Add missing / drop unused deps                | `go mod tidy`                      |
| `go get <pkg>@<ver>`           | Add or bump a dependency                      | `go get github.com/foo/bar@v1.2.3` |
| `go test ./...`                | Run tests in all packages                     | `go test ./...`                    |
| `go test -run TestName ./...`  | Run one focused test                          | `go test -run TestGreet .`         |
| `go env GOVERSION GOOS GOARCH` | Show toolchain / target platform              | `go env GOVERSION`                 |
| `go version`                   | Installed toolchain version                   | `go version`                       |

Per-chapter check: `gofmt -l . && go vet ./... && go build ./...`.

## Troubleshooting

- **"go: cannot find main module"** — you ran `go run` from the repo root. `cd`
  into a chapter folder first (`cd 03-composite-types`), then `go run ./02-slices`.
  Or use `make -C 03-composite-types run 02-slices` from anywhere.
- **"unknown topic"** — topic names are the folder names including the number
  (`02-slices`, not `slices`).
- **`gofmt -l .` lists files** — run `gofmt -w .` (or `make fmt`) and re-check.
- **`go vet` complains about printf verbs** — fix the format string; vet is
  right more often than you are.
- **Module path mismatch after cloning/forking** — the module path is
  `github.com/0xengr/just-go/<slug>`. If you fork under your own account, do a
  find-and-replace across `go.mod` files (and update imports) or keep the
  original path and build locally.

## Glossary

See [`GLOSSARY.md`](./GLOSSARY.md) for plain-English definitions of the Go
terms used across the chapters.

## Further reading

- *Learning Go* (Jon Bodner, 2nd ed.) — the book this repo follows.
- [A Tour of Go](https://go.dev/tour/) — interactive intro.
- [Effective Go](https://go.dev/doc/effective_go) — idioms.
- [Go by Example](https://gobyexample.com/) — short runnable snippets.
- [Go standard library docs](https://pkg.go.dev/std).

## Contributing

This repo is considered complete as a beginner guide (chapters 1–7). Contributions
that fix typos, correct facts or output, improve notes, or add the missing
chapters (8–16) following the existing layout are welcome. Please keep each
chapter a self-contained module with the shared `Makefile`, run `make check`
before submitting, and open a PR rather than pushing to `main`.

## License

MIT — see [`LICENSE`](./LICENSE).