Draft a PR description for the current branch changes. Include summary,
motivation, testing, and risks. Keep it readable in under a minute.

First gather the diff with your shell tool:

- `git diff --stat main...HEAD` (fall back to `git diff --stat` if `main` is
  unavailable)
- `git diff main...HEAD` (same fallback)

Draft only; do not create or modify anything.
