---
description: Fix grammar, prose, and comment tags in a file
agent: build
---

Fix grammar, spelling, and clarity in $ARGUMENTS. Preserve the author's voice and all code. Apply technical-writing and unslop rules. If no path is given, ask which file.

Also polish comment tags per the repo convention (AGENTS.md):
- Tag only the first line of a comment block; continuation lines stay untagged.
- Use the word tags only: NOTE:, WARNING:, BUG:, PERF:, TODO:. (The extension accepts ?:, *:, !:, #:, NAME(user): too, but this repo uses the word tags.)
- Tag meanings: NOTE: file header + ordinary explanation; WARNING: beginner traps; BUG: commented-out code that panics or fails to compile; PERF: capacity growth / map ordering / copy and GC costs; TODO: genuinely unfinished.
- Never tag commented-out code itself — it is code, not prose.
- Fix tag misuse, e.g. a lone `#` or missing colon, by converting to the appropriate word tag (usually NOTE:).
- For consistency, tag the first line of each explanatory comment block with NOTE: when it is currently untagged, leaving short inline comments (like `// or`, `// Example:`) untagged.
- Do not change what any program does; only fix prose and tags. Run gofmt/go vet/go build when the file is Go.