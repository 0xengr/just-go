# Just Go

Learning Go by writing it. Every topic here is a small runnable program, and
each chapter ends with notes in my own words — the bits I want to remember when
I come back to the code.

Chapters follow the order of *Learning Go* by Jon Bodner (2nd edition), one
chapter per folder, so each `NN-slug/` can be read on its own.

## Start here

```sh
make list                      # every chapter and every topic inside it
make check                     # gofmt + vet + build across all chapters
make -C 03-composite-types run 02-slices   # run one topic, with the banner
```

Then open the chapter's own `README.md`: it opens with the big idea, a
run-order table, and one section per topic with the code and its real output.

## Layout

39 runnable topics across 6 chapters. `Start here` runs the first topic in
reading order; the chapter `README.md` has the full run-order table.

| Dir                          | Chapter                              | Topics | Start here                                          |
| ---------------------------- | ------------------------------------ | -----: | --------------------------------------------------- |
| `01-hello-go/`               | Go environment & toolchain           |    2 | `make -C 01-hello-go run 01-greet Ada`              |
| `02-predeclared-types/`      | Predeclared types & declarations     |    6 | `make -C 02-predeclared-types run 01-zero-values`   |
| `03-composite-types/`        | Arrays, slices, maps & structs       |   11 | `make -C 03-composite-types run 01-arrays`          |
| `04-blocks-shadows-control/` | Blocks, shadows & control structures |    9 | `make -C 04-blocks-shadows-control run 01-blocks`   |
| `05-functions/`              | Functions, closures & `defer`        |    5 | `make -C 05-functions run 01-multiple-returns`      |
| `06-pointers/`               | Pointers, nil, `new()` & call by value |  6 | `make -C 06-pointers run 01-call-by-value`          |

Each chapter is its own Go module (`go.mod` naming `github.com/aman-void/just-go/<slug>`)
with its own overview `main.go` and `README.md`; there is no root module and no
`go.work`. Chapter 1 is the exception: its overview `main.go` is a runnable hello
world, the rest are orientation stubs.

Topic folders are numbered in reading order (`01-`, `02-`, …) within each
chapter, so `ls` on any chapter shows a newcomer where to begin. The numbers are
the curriculum, not a version.

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
| `make check`              | `gofmt` + `go vet` + `go build` for the chapter             |

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
