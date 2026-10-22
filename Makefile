# ─────────────────────────────────────────────────────────────────────────────
#  Just GO · repo tasks
#
#  The repo is a shelf of independent chapter modules (each with its own go.mod
#  and its own Makefile). This one fans tasks out across the whole shelf.
#
#    make                     this help
#    make list                chapters and every topic they contain
#    make check               gofmt + vet + build across all chapters
#    make run <chapter> [topic] [args]
#                             run a chapter's overview (main.go) or one of its topics
#    make new-chapter <name>  scaffold a new chapter (alias: make new <name>)
#
#  Chapter work stays inside the chapter's own Makefile:
#    make -C 05-functions            help for that chapter
#    make -C 05-functions run slices
# ─────────────────────────────────────────────────────────────────────────────

SHELL := /bin/sh
.DEFAULT_GOAL := help

# A chapter is any folder that is a Go module.
CHAPTERS := $(patsubst %/go.mod,%,$(wildcard */go.mod))

# `make run 08-interfaces topic arg` / `make new-chapter 08-interfaces` leave
# the extra words behind as stray goals; read them back here.
CHAP  := $(word 2,$(MAKECMDGOALS))
TOPIC := $(word 3,$(MAKECMDGOALS))
ARGS  := $(wordlist 4,$(words $(MAKECMDGOALS)),$(MAKECMDGOALS))

C_RESET := \033[0m
C_BOLD  := \033[1m
C_DIM   := \033[2m
C_GREEN := \033[38;5;42m
C_RED   := \033[38;5;203m
C_CYAN  := \033[38;5;44m

.PHONY: help list check run new new-chapter chapter
ifneq ($(CHAP),)
.PHONY: $(CHAP)
endif

help:
	@printf '\n'
	@printf '  %bJust GO%b  ·  learn Go by running it\n\n' '$(C_BOLD)' '$(C_RESET)'
	@printf '  %bmake list%b               chapters and their topics\n' '$(C_GREEN)' '$(C_RESET)'
	@printf '  %bmake check%b              gofmt + vet + build every chapter\n' '$(C_GREEN)' '$(C_RESET)'
	@printf '  %bmake run%b <chapter> [topic] [args]\n' '$(C_GREEN)' '$(C_RESET)'
	@printf '                         run a chapter overview (main.go) or a topic under examples/, exercises/, or a nested path\n' '$(C_DIM)' '$(C_RESET)'
	@printf '  %bmake new-chapter%b <name>  scaffold a new chapter (alias: make new)\n' '$(C_GREEN)' '$(C_RESET)'
	@printf '  %bmake -C <chapter> ...%b   any chapter task (run, new, list, check)\n\n' '$(C_GREEN)' '$(C_RESET)'
	@printf '  %bexample%b  make -C 05-functions run defer-func\n' '$(C_DIM)' '$(C_RESET)'
	@printf '  %bexample%b  make -C 08-interfaces run exercises/duck-typing-2\n\n' '$(C_DIM)' '$(C_RESET)'

list:
	@for c in $(CHAPTERS); do \
	  printf '\n  %b%s%b\n' '$(C_BOLD)' "$$c" '$(C_RESET)'; \
	  $(MAKE) --no-print-directory -C $$c list; \
	done

check:
	@fail=0; \
	for c in $(CHAPTERS); do \
	  printf '\n  %b== %s ==%b\n' '$(C_BOLD)' "$$c" '$(C_RESET)'; \
	  $(MAKE) --no-print-directory -C $$c check || fail=1; \
	done; \
	if [ $$fail -ne 0 ]; then \
	  printf '\n  %bsome chapters are not clean%b\n' '$(C_RED)' '$(C_RESET)'; \
	  exit 1; \
	fi; \
	printf '\n  %ball chapters clean%b\n' '$(C_GREEN)' '$(C_RESET)'

run:
	@if [ -z "$(CHAP)" ]; then \
	  printf '  %bno chapter given%b — try: make run <chapter> [topic]\n\n' '$(C_RED)' '$(C_RESET)'; \
	  $(MAKE) --no-print-directory list; \
	  exit 1; \
	fi
	@if [ ! -d "$(CHAP)" ]; then \
	  printf '  %bunknown chapter:%b %s\n\n' '$(C_RED)' '$(C_RESET)' '$(CHAP)'; \
	  $(MAKE) --no-print-directory list; \
	  exit 1; \
	fi
	@$(MAKE) --no-print-directory -C $(CHAP) run $(TOPIC) $(ARGS)

new-chapter new chapter:
	@if [ -z "$(CHAP)" ]; then \
	  printf '  %bno name given%b — try: make new-chapter <name>\n' '$(C_RED)' '$(C_RESET)'; \
	  exit 1; \
	fi
	@if [ -e "$(CHAP)" ]; then \
	  printf '  %b%s already exists%b\n' '$(C_RED)' '$(CHAP)' '$(C_RESET)'; \
	  exit 1; \
	fi
	@if [ ! -f "01-hello-go/Makefile" ]; then \
	  printf '  %btemplate chapter 01-hello-go/Makefile not found%b\n' '$(C_RED)' '$(C_RESET)'; \
	  exit 1; \
	fi
	@mkdir -p "$(CHAP)/examples" "$(CHAP)/exercises"
	@cp 01-hello-go/Makefile "$(CHAP)/Makefile"
	@printf 'module github.com/0xengr/just-go/%s\n\ngo 1.27.2\n' '$(CHAP)' > "$(CHAP)/go.mod"
	@printf 'package main\n\nimport "fmt"\n\n// NOTE: chapter overview\nfunc main() {\n\tfmt.Println("Hello from %s")\n}\n' '$(CHAP)' > "$(CHAP)/main.go"
	@gofmt -w "$(CHAP)/main.go"
	@touch "$(CHAP)/README.md" "$(CHAP)/notes.md"
	@printf '  %b✓%b created chapter %s\n' '$(C_GREEN)' '$(C_RESET)' '$(CHAP)'
	@printf '  %bnext:%b cd %s && make run   (or add examples under examples/)\n' '$(C_DIM)' '$(C_RESET)' '$(CHAP)'

# Swallow stray chapter/topic goals passed to `make run ...` / `make new-chapter ...`.
ifneq ($(CHAP),)
$(CHAP):
	@if [ "$(MAKECMDGOALS)" = "$@" ]; then \
	  printf '  %s is a chapter, not a target — try: make run %s\n' '$(C_DIM)' '$@' '$(C_RESET)' '$@'; \
	  exit 1; \
	fi
	@:
endif

%:
	@if [ "$(MAKECMDGOALS)" = "$@" ]; then \
	  printf '  %bunknown target:%b %s — try: make help\n' '$(C_RED)' '$(C_RESET)' '$@'; \
	  exit 1; \
	fi
	@:
