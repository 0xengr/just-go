Prepare a commit following AGENTS.md strictly. Do not stage, commit, or push
without my explicit approval.

First gather the current state with your shell tool and show it to me:

- `git status --short`
- `git diff --cached --stat && git diff --cached` (staged changes)
- `git diff --stat && git diff` (unstaged changes)

Then:

1. Ask whether this commit targets `main` or a chapter branch. If the changes are
   chapter-specific, suggest a branch name from the changed paths.
2. Propose a commit message, then wait for my approval before running `git add` or
   `git commit`.
3. After a successful commit, ask whether to push: reply 1 to confirm (push) or 0
   to decline (leave unpushed).
