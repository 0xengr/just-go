# AGENTS.md

- **Git: high strictness.** Never write a commit, create a branch, or push any
  code without the author's explicit consent. Prepare the work, show the diff or
  staged changes, and wait for approval before touching git state.
- Root `README.md` is the codebase index + Go tooling reference; each `NN-slug/README.md` is the chapter index and `NN-slug/notes.md` holds the full notes.
- **Knowledge map (system of record).** All context lives in the repo:
  - `README.md` — index and tooling reference
  - `GLOSSARY.md` — terms
  - `fmt-printing.md` — fmt reference
  - `NN-slug/README.md` — chapter index
  - `NN-slug/notes.md` — full notes
  - `NN-slug/examples/` — runnable code
  - `NN-slug/exercises/` — practice tasks
  - `.opencode/commands/` — custom commands
- **Progressive disclosure.** Read in this order: `AGENTS.md` → root `README.md` →
  chapter `README.md` → `notes.md` → `examples/`.
- **Mechanical checks.** `make check` is the gate; rules must be verifiable, not
  just documented.
- Each `NN-slug/` is an independent Go module with its own `go.mod` (`github.com/0xengr/just-go/<slug>`, Go 1.27). No root `go.mod`, no `go.work`.
- Chapter layout: root `main.go` is overview only; each subtopic is a runnable
  `package main` under `examples/<topic>/` (`go run ./examples/<topic>`).
  Exercises live under `exercises/<topic>/`. Never add a nested `go.mod`.
- Every chapter ships an identical `Makefile` (copy it verbatim to new chapters — it derives the chapter name from its folder and the topics from `examples/*/main.go`, so never hand-edit it). Targets: `help`, `list`, `run <topic> [args]`, `new <topic>`, `build`, `vet`, `test`, `fmt`, `check`, `tidy`; `make run <topic> ANIM=0` skips the banner. The root `Makefile` fans `list`/`check` across chapters.
- New chapters: `NN-kebab-slug/` + `go.mod` + `main.go` overview + `examples/<topic>/` + `exercises/<topic>/` + `README.md` index + `notes.md` full notes + the chapter `Makefile`. Keep every example original.
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
- Chapter `README.md` is the index (big idea, run order, cheat sheet, exercises
  pointer); full notes live in `notes.md`. Run-order entries point at
  `examples/<topic>`. Exercises live under `exercises/<topic>/` (code only); the chapter
  README.md lists each exercise with the topics it solves.
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
  1. Write or update that chapter's `README.md` (index) and `notes.md`
     (full notes), covering every subtopic, with output quoted from the
     actual programs — re-run them rather than trusting earlier output.
  2. Add or update the chapter's row in the root `README.md` Layout table.
  3. Only then treat the chapter as finished.

  Root `README.md` is the index, so a missing Layout row means the chapter is
  effectively invisible to a reader.
- **Fix the author's prose and code mistakes.** The author writes quickly and the
  grammar is often broken; the examples are theirs but the errors are not. On
  every pass over a file they wrote, silently correct:
  - Grammar and spelling in comments and identifiers: `it's types` → `its types`,
    `Like most programming language GO also have` → `Like most programming
    languages, Go also has`, subject/verb agreement, `difference` → `different`.
  - `GO` → `Go` (the language is not an acronym).
  - Wrong facts and wrong claims, not just wording. If a comment says something
    does not compile, verify it does not, and quote the real error. If a comment
    is wrong about the language, fix it and say so in the commit or review.
  - Stale output blocks, dead paths, and broken references — re-run the program
    rather than trusting the old text.

  Fix these **without being asked and without asking first**; do not flag grammar
  as a question for the author. Keep the author's voice, their variable names, and
  their example structure — correct the mistakes, do not rewrite their prose into
  a different voice. Never change what a program *does*; only fix things that are
  wrong. If a correction is debatable, prefer the plainer phrasing and mention it
  in the summary rather than leaving a known error in place.
- Chapters are numbered in learning order: `01-hello-go`, `02-predeclared-types`, `03-composite-types`, …
