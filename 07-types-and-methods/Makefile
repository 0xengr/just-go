# ─────────────────────────────────────────────────────────────────────────────
#  Just GO · chapter Makefile
#
#  One identical Makefile per chapter. It auto-discovers every topic (a folder
#  containing a main.go) and the chapter name from the directory, so there is
#  nothing to edit when a chapter or topic is added.
#
#    make                    help (default)
#    make list               list this chapter's topics
#    make run <topic> [args]  go run ./<topic> [args], with the animated banner
#    make new <topic>        scaffold <topic>/main.go
#    make build              go build ./...
#    make vet                go vet ./...
#    make test               go test ./...
#    make fmt                gofmt -w .
#    make check              gofmt + vet + build (the chapter gate)
#    make tidy               go mod tidy
#
#  Add ANIM=0 to skip the animation:  make run slices ANIM=0
# ─────────────────────────────────────────────────────────────────────────────

SHELL := /bin/sh
.DEFAULT_GOAL := help

# Identity is derived, never edited: the folder name is the chapter, and every
# subfolder with a main.go is a runnable topic.
CHAPTER := $(notdir $(CURDIR))
TOPICS  := $(patsubst %/main.go,%,$(wildcard */main.go))

# `make run defer-func` asks make for two goals: `run` and `defer-func`. The
# catch-all rule at the bottom swallows the second one; here we read it back.
# Anything after the topic is forwarded to the program as arguments, so
# `make run greet Ada` runs `go run ./greet Ada`.
TOPIC := $(word 2,$(MAKECMDGOALS))
ARGS  := $(wordlist 3,$(words $(MAKECMDGOALS)),$(MAKECMDGOALS))

# --- ANSI colours -----------------------------------------------------------
C_RESET   := \033[0m
C_DIM     := \033[2m
C_BOLD    := \033[1m
C_GREEN   := \033[38;5;42m
C_CYAN    := \033[38;5;44m
C_BLUE    := \033[38;5;39m
C_VIOLET  := \033[38;5;135m
C_MAGENTA := \033[38;5;205m
C_YELLOW  := \033[38;5;220m
C_RED     := \033[38;5;203m

ANIM ?= 1

.PHONY: help list run new build vet test fmt fmt-check check tidy $(TOPICS)

# --- help (default) ---------------------------------------------------------
help:
	@printf '\n'
	@printf '  %bJ%b%bu%b%bs%b%bt%b  %bG%b%bO%b  %blearn Go by running it%b\n\n' \
	  '$(C_CYAN)' '$(C_RESET)' '$(C_BLUE)' '$(C_RESET)' \
	  '$(C_VIOLET)' '$(C_RESET)' '$(C_MAGENTA)' '$(C_RESET)' \
	  '$(C_YELLOW)' '$(C_RESET)' '$(C_GREEN)' '$(C_RESET)' \
	  '$(C_DIM)' '$(C_RESET)'
	@printf '  %bchapter%b  %s\n\n' '$(C_BOLD)' '$(C_RESET)' '$(CHAPTER)'
	@printf '  %busage%b    make <target> [topic]\n\n' '$(C_BOLD)' '$(C_RESET)'
	@printf '    %brun%b <topic> [args]   go run ./<topic> [args] with the banner\n' '$(C_GREEN)' '$(C_RESET)'
	@printf '    %bnew%b <topic>   scaffold <topic>/main.go\n' '$(C_GREEN)' '$(C_RESET)'
	@printf '    %blist%b          list topics in this chapter\n' '$(C_GREEN)' '$(C_RESET)'
	@printf '    %bbuild%b         go build ./...\n' '$(C_GREEN)' '$(C_RESET)'
	@printf '    %bvet%b           go vet ./...\n' '$(C_GREEN)' '$(C_RESET)'
	@printf '    %btest%b          go test ./...\n' '$(C_GREEN)' '$(C_RESET)'
	@printf '    %bfmt%b           gofmt -w .\n' '$(C_GREEN)' '$(C_RESET)'
	@printf '    %bcheck%b         gofmt + vet + build (the chapter gate)\n' '$(C_GREEN)' '$(C_RESET)'
	@printf '    %btidy%b          go mod tidy\n' '$(C_GREEN)' '$(C_RESET)'
	@printf '\n'

# --- list topics ------------------------------------------------------------
list:
	@printf '  %btopics in %s%b\n' '$(C_BOLD)' '$(CHAPTER)' '$(C_RESET)'
	@for t in $(TOPICS); do \
	  printf '    %b·%b %s\n' '$(C_GREEN)' '$(C_RESET)' "$$t"; \
	done
	@printf '\n'

# --- run one topic ----------------------------------------------------------
run:
	@if [ -z "$(TOPIC)" ]; then \
	  printf '  %bno topic given%b — try: make run <topic>\n\n' '$(C_RED)' '$(C_RESET)'; \
	  $(MAKE) --no-print-directory list; \
	  exit 1; \
	fi
	@if [ ! -f "$(TOPIC)/main.go" ]; then \
	  printf '  %bunknown topic:%b %s\n\n' '$(C_RED)' '$(C_RESET)' '$(TOPIC)'; \
	  $(MAKE) --no-print-directory list; \
	  exit 1; \
	fi
	@printf '\n'
	@if [ "$(ANIM)" = "1" ]; then \
	  printf '  '; \
	  printf '%bJ%b' '$(C_CYAN)' '$(C_RESET)'; sleep 0.04; \
	  printf '%bu%b' '$(C_BLUE)' '$(C_RESET)'; sleep 0.04; \
	  printf '%bs%b' '$(C_VIOLET)' '$(C_RESET)'; sleep 0.04; \
	  printf '%bt%b' '$(C_MAGENTA)' '$(C_RESET)'; sleep 0.04; \
	  printf '  '; \
	  printf '%bG%b' '$(C_YELLOW)' '$(C_RESET)'; sleep 0.04; \
	  printf '%bO%b' '$(C_GREEN)' '$(C_RESET)'; sleep 0.04; \
	  printf '\n'; \
	else \
	  printf '  %bJ%b%bu%b%bs%b%bt%b  %bG%b%bO%b\n' \
	    '$(C_CYAN)' '$(C_RESET)' '$(C_BLUE)' '$(C_RESET)' \
	    '$(C_VIOLET)' '$(C_RESET)' '$(C_MAGENTA)' '$(C_RESET)' \
	    '$(C_YELLOW)' '$(C_RESET)' '$(C_GREEN)' '$(C_RESET)'; \
	fi
	@printf '  %bIn chapter %s · running%b %b./%s%b\n\n' \
	  '$(C_DIM)' '$(CHAPTER)' '$(C_RESET)' '$(C_CYAN)' '$(TOPIC)' '$(C_RESET)'
	@if [ "$(ANIM)" = "1" ]; then \
	  for f in '⠋' '⠙' '⠹' '⠸' '⠼' '⠴' '⠦' '⠧'; do \
	    printf '\r  %s warming up ./%s …' "$$f" '$(TOPIC)'; \
	    sleep 0.05; \
	  done; \
	  printf '\r\033[K'; \
	fi
	@go run ./$(TOPIC) $(ARGS) && printf '\n  %b✓ done%b\n' '$(C_GREEN)' '$(C_RESET)' \
	  || { printf '\n  %b✗ go run ./%s failed%b\n' '$(C_RED)' '$(TOPIC)' '$(C_RESET)'; exit 1; }

# --- scaffold a new topic ---------------------------------------------------
new:
	@if [ -z "$(TOPIC)" ]; then \
	  printf '  %bno name given%b — try: make new <topic>\n' '$(C_RED)' '$(C_RESET)'; \
	  exit 1; \
	fi
	@if [ -e "$(TOPIC)" ]; then \
	  printf '  %b%s already exists%b\n' '$(C_RED)' '$(TOPIC)' '$(C_RESET)'; \
	  exit 1; \
	fi
	@mkdir -p "$(TOPIC)"
	@printf 'package main\n\nimport "fmt"\n\n// NOTE: what this example shows\nfunc main() {\n\tfmt.Println("%s")\n}\n' '$(TOPIC)' > "$(TOPIC)/main.go"
	@gofmt -w "$(TOPIC)/main.go"
	@printf '  %b✓%b created %s/main.go — run it with: make run %s\n' \
	  '$(C_GREEN)' '$(C_RESET)' '$(TOPIC)' '$(TOPIC)'

# --- checks and chores ------------------------------------------------------
build:
	@go build ./... && printf '  %b✓%b build\n' '$(C_GREEN)' '$(C_RESET)'

vet:
	@go vet ./... && printf '  %b✓%b vet\n' '$(C_GREEN)' '$(C_RESET)'

test:
	@go test ./... && printf '  %b✓%b test\n' '$(C_GREEN)' '$(C_RESET)'

fmt:
	@gofmt -w .
	@printf '  %b✓%b gofmt\n' '$(C_GREEN)' '$(C_RESET)'

fmt-check:
	@out=$$(gofmt -l .); \
	if [ -n "$$out" ]; then \
	  printf '  %b✗ gofmt needed:%b\n%s\n' '$(C_RED)' '$(C_RESET)' "$$out"; \
	  exit 1; \
	fi; \
	printf '  %b✓%b gofmt\n' '$(C_GREEN)' '$(C_RESET)'

check: fmt-check vet build
	@printf '  %b%s is clean%b\n' '$(C_GREEN)' '$(CHAPTER)' '$(C_RESET)'

tidy:
	@go mod tidy
	@printf '  %b✓%b go mod tidy\n' '$(C_GREEN)' '$(C_RESET)'

# `make run defer-func` leaves `defer-func` behind as a second goal. Known
# topics get an explicit recipe so make stays silent; asking for one on its own
# still points you at `make run <topic>`.
ifneq ($(TOPICS),)
$(TOPICS):
	@if [ "$(MAKECMDGOALS)" = "$@" ]; then \
	  printf '  %b%s is a topic, not a target%b — try: make run %s\n' \
	    '$(C_DIM)' '$@' '$(C_RESET)' '$@'; \
	  exit 1; \
	fi
endif

# Swallow `make run defer-func`'s stray second goal, but still complain when an
# unknown target is the only thing asked for.
%:
	@if [ "$(MAKECMDGOALS)" = "$@" ]; then \
	  printf '  %bunknown target:%b %s — try: make help\n' '$(C_RED)' '$(C_RESET)' '$@'; \
	  exit 1; \
	fi
