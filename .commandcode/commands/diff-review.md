Review $ARGUMENTS (or the diff below) for correctness, missing tests, and Go
idioms.

First gather the diff with your shell tool:

- `git diff --cached --stat && git diff --cached` (staged)
- `git diff --stat && git diff` (working tree)

Report bugs first, then gaps, then nits. Review read-only: do not edit files.
