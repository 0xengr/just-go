# 01 — Setting Up Your Go Environment

Original notes on the toolchain and module basics.

## What I ran

- `go mod init github.com/aman-void/just-go/01-hello-go` — one module per chapter.
- `go run .` — overview; `go run ./01-greet Ada` and `go run ./02-toolchain` for subexamples.
- `make run 01-greet Ada` / `make run 02-toolchain` — the chapter Makefile runs a topic; `make list` lists them.

| Order | Run it | What it teaches you |
| ----- | ------ | ------------------- |
| 1 | `go run ./01-greet Ada` | A first program with a command-line argument |
| 2 | `go run ./02-toolchain` | What `go env` and `go version` actually report |

Topics are numbered folders (`01-`, `02-`, …) in reading order, so `ls` alone
tells a newcomer where to begin. The numbers are the curriculum, not a version.
- `make check` — gofmt + vet + build; same as `go vet ./...`, `gofmt -l .`, `go build ./...`.

## Notes

- A module is just a directory with a `go.mod` naming it; the toolchain uses
  that path for imports. No build config beyond that for a small program.
- `go run` compiles to a temp dir (nothing saved); `go build` leaves a binary.
- `gofmt` is the style authority — run it instead of debating formatting.
- `go vet` catches real mistakes (bad format verbs, unreachable code) that
  compile fine, so run it before assuming code is correct.
- The compat promise means code written today should still build on newer Go;
  upgrade the toolchain freely, pin the `go` line in `go.mod` per module.
- Editor choice doesn't matter much: any editor + terminal with the commands
  above is enough; the playground is handy for sharing snippets.
