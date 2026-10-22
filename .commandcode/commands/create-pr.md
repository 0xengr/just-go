Create a PR, review it, then merge or roll back — with approvals at every step.
Follow AGENTS.md git rules: do not push or merge without my explicit approval.

First gather the branch state with your shell tool:

- `git branch --show-current && git status --short`
- `git diff --stat main...HEAD` (fall back to `git diff --stat` if `main` is
  unavailable)
- `git diff main...HEAD` (same fallback)

Steps:

1. If we are on `main` or there are uncommitted changes, stop and ask me what to do.
2. Draft a PR title and body from the diff above, then ask for approval before
   creating the PR.
3. After creation, ask me to review. Offer to run `/review` (the built-in PR
   review) or show the PR diff. Wait for approval.
4. If approved, ask before merging with `gh pr merge`. If the merge succeeds,
   leave the branch in place — branches are permanent per AGENTS.md.
5. If the review fails or I decline, offer a rollback: close the PR first. Revert
   or reset only as a last resort, and only after explicit approval. Never
   force-push or delete branches.
