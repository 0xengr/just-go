---
description: Review current changes for bugs, missing tests, and Go idioms
agent: plan
subagent: true
---

Review $ARGUMENTS (or the diff below) for correctness, missing tests, and Go idioms. Report bugs first, then gaps, then nits.

Staged diff:
!`git diff --cached --stat && git diff --cached`

Working tree diff:
!`git diff --stat && git diff`
