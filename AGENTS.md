# AGENTS.md

- **Git: high strictness.** Never write a commit, create a branch, or push any
  code without the author's explicit consent. Prepare the work, show the diff or
  staged changes, and wait for approval before touching git state.
- Root `README.md` is the codebase index + Go tooling reference; each `NN-slug/README.md` holds that chapter's notes.
- Each `NN-slug/` is an independent Go module with its own `go.mod` (`github.com/aman-void/just-go/<slug>`, Go 1.27). No root `go.mod`, no `go.work`.
- Chapter layout: root `main.go` is overview only; each subtopic is a runnable `package main` in its own folder (`go run ./<topic>`). Never add a nested `go.mod`.
- Every chapter ships an identical `Makefile` (copy it verbatim to new chapters — it derives the chapter name from its folder and the topics from `*/main.go`, so never hand-edit it). Targets: `help`, `list`, `run <topic> [args]`, `new <topic>`, `build`, `vet`, `test`, `fmt`, `check`, `tidy`; `make run <topic> ANIM=0` skips the banner. The root `Makefile` fans `list`/`check` across chapters.
- New chapters: `NN-kebab-slug/` + `go.mod` + `main.go` overview + topic subfolders + `README.md` notes + the chapter `Makefile`. Keep every example original.
- **Branch layout.** Two kinds of work, two homes:
  - Root-level files (`README.md`, `Makefile`, `fmt-printing.md`) are committed on
    `main`.
  - Every `NN-slug/` chapter lives on its own branch, named after the chapter
    (`01-hello-go`, `03-composite-types`, …).
- **Chapter lifecycle.** Create the chapter branch, write the code, commit it there,
  push it, then merge it into `main`. **Never delete the branch** — not locally, not
  on the remote. The branches are a permanent per-chapter reading path, so anyone
  can `git checkout NN-slug` (or browse it on the remote) and read that chapter in
  isolation without paging through the others.
- The reference PDF in the repo root is local reading material only. It is never
  committed, never staged, and stays ignored (`*.pdf` in `.gitignore`). Ask about it
  as context, not as tracked content.
- Chapter check: `make check` (same as `gofmt -l . && go vet ./... && go build ./...`). Root has no Go code to build.
- Session CWD defaults to `01-hello-go/`; repo root is its parent `just-go/`. `cd` to the target chapter dir before running Go commands.
- Chapter `README.md` style (human-friendly notes): big-idea intro first, then a
  run-order table (`Order | Run it | What it teaches you`), then one section per
  subtopic with plain-English idea + real-world analogy + short code/output quoted
  from the actual program + one beginner trap, ending with a rules cheat sheet.
- Comment tags (Zed `zed-comment` extension, already installed): tag the **first
  line only** of a comment block; continuation lines stay untagged. Tag meanings:
  `NOTE:` file header + ordinary explanation, `WARNING:` beginner traps, `BUG:`
  commented-out code that panics or fails to compile, `PERF:` capacity growth /
  map ordering / copy and GC costs, `TODO:` genuinely unfinished. Extension also
  accepts `?:`, `*:`, `!:`, `#:` and `NAME(user):`, but this repo uses the word
  tags above. Never tag commented-out code itself — it is code, not prose.
- **Chapter completion gate.** After writing topic code, always ask: *"Is this
  chapter complete, or are you going to write more examples?"* Never assume a
  chapter is done. When the answer is yes:
  1. Write or update that chapter's `README.md` (style above), covering every
     subtopic, with output quoted from the actual programs — re-run them rather
     than trusting earlier output.
  2. Add or update the chapter's row in the root `README.md` Layout table.
  3. Only then treat the chapter as finished.

  Root `README.md` is the index, so a missing Layout row means the chapter is
  effectively invisible to a reader.
- Chapters are numbered in learning order: `01-hello-go`, `02-predeclared-types`, `03-composite-types`, …
