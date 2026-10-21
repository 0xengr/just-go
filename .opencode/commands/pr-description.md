---
description: Draft a PR description from branch changes
agent: plan
---

Draft a PR description for the changes below. Include summary, motivation, testing, and risks. Keep it readable in under a minute.

Diff stat:
!`git diff --stat main...HEAD 2>/dev/null || git diff --stat`

Diff:
!`git diff main...HEAD 2>/dev/null || git diff`
