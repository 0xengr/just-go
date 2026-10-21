---
description: Prepare a commit following AGENTS.md (asks before commit and push)
agent: plan
---

Follow AGENTS.md git rules strictly. Do not commit or push without explicit approval.

Staged status:
!`git status --short`

Staged diff:
!`git diff --cached --stat && git diff --cached`

Unstaged diff:
!`git diff --stat && git diff`

Ask whether this commit targets main or a chapter branch. If the changes are chapter-specific, suggest a branch name from the changed paths. Propose a commit message, then wait for approval before running git add or git commit. After a successful commit, ask whether to push.
