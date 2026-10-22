---
description: Explain a topic related to this repository
agent: plan
---

Explain $ARGUMENTS in the context of this repository.

If no argument was provided, first ask the user what they would like to understand, then proceed.

Guidelines:
- Ground the explanation in this repo: read the relevant chapter notes, examples, or source files before answering.
- Cite concrete paths and symbols (e.g. `03-composite-types/examples/maps/main.go`).
- Keep it concise and beginner-friendly; offer a deeper dive if useful.
- If the topic is not covered by the repo, say so and point to the closest related material.

Repository overview:
!`ls && echo "---" && find . -maxdepth 2 -name "README.md" | head -20`