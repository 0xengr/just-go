---
description: Summarize current changes for a human reader
agent: plan
---

Summarize the current changes for a human reader. Keep it short and concrete. Use the diffs below.

Staged diff:
!`git diff --cached --stat && git diff --cached`

Working tree diff:
!`git diff --stat && git diff`
