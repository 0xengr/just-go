# Just Go

Learning Go by writing it. Every topic here is a small runnable program, and
each chapter ends with notes in my own words — the bits I want to remember when
I come back to the code.

## Layout

| Dir                    | Chapter | Run it                     |
| ---------------------- | ------- | -------------------------- |
| `01-hello-go/`         | Go environment & toolchain | `cd 01-hello-go && go run .` |
| `02-predeclared-types/`| Predeclared types & declarations | `cd 02-predeclared-types && go run .` |
| `03-composite-types/`  | Arrays, slices, maps & structs | `cd 03-composite-types && go run .` |
| `04-blocks-shadows-control/` | Blocks, shadows & control structures | `cd 04-blocks-shadows-control && go run .` |
| `05-functions/`        | Functions, closures & `defer` | `cd 05-functions && go run .` |
| `06-pointers/`         | Pointers, nil, `new()` & call by value | `cd 06-pointers && go run .` |

Each chapter: overview in root `main.go`, runnable subtopic in its own folder
(`go run ./<topic>`), notes in that chapter's `README.md`.

## Make

Every chapter has the same `Makefile`; it finds the chapter name and its topics
by itself, so there is nothing to configure.

| Command | What it does |
| ------- | ------------ |
| `make` / `make help` | list the targets |
| `make list` | list this chapter's topics |
| `make run <topic> [args]` | `go run ./<topic> [args]`, with a banner (`ANIM=0` to skip) |
| `make new <topic>` | scaffold `<topic>/main.go` |
| `make test` | `go test ./...` |
| `make check` | `gofmt` + `go vet` + `go build` for the chapter |

Run them from a chapter folder, e.g. `cd 05-functions && make run multiple-returns`.
At the repo root, `make list` shows every chapter's topics and `make check` runs
the gate across all of them; `make -C <chapter> <target>` works too.

Reference notes: `fmt-printing.md` (all 12 `fmt` print functions + verbs).

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

| Command | What it does | Example |
| ------- | ------------ | ------- |
| `go run .` / `go run ./<dir>` | Compile to temp dir and run | `go run .`, `go run ./greet` |
| `go build ./...` | Compile all packages, check for errors | `go build ./...` |
| `go vet ./...` | Static checks (printf verbs, dead code, etc.) | `go vet ./...` |
| `gofmt -l .` | List files needing format (empty = clean) | `gofmt -l .`, `gofmt -w .` |
| `go mod init <path>` | Start a new module | `go mod init example.com/demo` |
| `go mod tidy` | Add missing / drop unused deps | `go mod tidy` |
| `go get <pkg>@<ver>` | Add or bump a dependency | `go get github.com/foo/bar@v1.2.3` |
| `go test ./...` | Run tests in all packages | `go test ./...` |
| `go test -run TestName ./...` | Run one focused test | `go test -run TestGreet .` |
| `go env GOVERSION GOOS GOARCH` | Show toolchain / target platform | `go env GOVERSION` |
| `go version` | Installed toolchain version | `go version` |

Per-chapter check: `gofmt -l . && go vet ./... && go build ./...`.
