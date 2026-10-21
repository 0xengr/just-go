---
description: Create a PR, review it, then merge or roll back (with approvals)
agent: plan
---

Follow AGENTS.md git rules. Do not push or merge without explicit approval.

Current branch and status:
!`git branch --show-current && git status --short`

Branch diff:
!`git diff --stat main...HEAD 2>/dev/null || git diff --stat`
!`git diff main...HEAD 2>/dev/null || git diff`

Steps:
1. If on main or there are uncommitted changes, stop and ask what to do.
2. Draft a PR title and body from the diff above, then ask for approval before creating the PR.
3. After creation, ask the user to review. Offer to run /review or show the PR diff. Wait for approval.
4. If approved, ask before merging with gh pr merge. If merge succeeds, ask whether to delete the branch.
5. If review fails or the user declines, offer rollback: close the PR and revert/reset as appropriate. Never force-push or delete branches without explicit approval.
