# ─────────────────────────────────────────────────────────────────────────────
#  Just GO · chapter Makefile
#
#  One identical Makefile per chapter. It auto-discovers every topic (a folder
#  containing a main.go) and the chapter name from the directory, so there is
#  nothing to edit when a chapter or topic is added.
#
#    make                    help (default)
#    make list               list this chapter's topics
#    make run [topic]        run ./main.go (overview) if no topic, else a topic
#                            under examples/, exercises/, or a nested path
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
TOPICS  := $(patsubst examples/%/main.go,%,$(wildcard examples/*/main.go))
# Runnable packages live under examples/ or exercises/, including nested ones.
RUNNABLES := $(sort $(patsubst %/main.go,%,$(shell find examples exercises -type f -name main.go 2>/dev/null)))

# `make run defer-func` asks make for two goals: `run` and `defer-func`. The
# catch-all rule at the bottom swallows the second one; here we read it back.
# Anything after the topic is forwarded to the program as arguments, so
# `make run greet Ada` runs `go run ./greet Ada`.
TOPIC := $(word 2,$(MAKECMDGOALS))
ARGS  := $(wordlist 3,$(words $(MAKECMDGOALS)),$(MAKECMDGOALS))

# What `make run` executes: a topic under examples/, exercises/, or a direct
# relative path; or the chapter's root main.go (the overview) when no topic /
# a reserved name is given.
RUN_TARGET :=
RUN_LABEL :=
ifneq ($(TOPIC),)
  ifneq ($(wildcard $(TOPIC)/main.go),)
    RUN_TARGET := ./$(TOPIC)
    RUN_LABEL := $(TOPIC)
  endif
  ifeq ($(RUN_TARGET),)
    ifneq ($(wildcard examples/$(TOPIC)/main.go),)
      RUN_TARGET := ./examples/$(TOPIC)
      RUN_LABEL := $(TOPIC)
    endif
  endif
  ifeq ($(RUN_TARGET),)
    ifneq ($(wildcard exercises/$(TOPIC)/main.go),)
      RUN_TARGET := ./exercises/$(TOPIC)
      RUN_LABEL := $(TOPIC)
    endif
  endif
  ifeq ($(RUN_TARGET),)
    ifneq ($(filter main root overview .,$(TOPIC)),)
      ifneq ($(wildcard main.go),)
        RUN_TARGET := .
        RUN_LABEL := main.go
      endif
    endif
  endif
else
  ifneq ($(wildcard main.go),)
    RUN_TARGET := .
    RUN_LABEL := main.go
  endif
endif

# New topic scaffolding defaults to examples/ unless an examples/ or exercises/
# prefix is given explicitly.
ifeq ($(filter examples/% exercises/%,$(TOPIC)),)
  NEW_PATH := examples/$(TOPIC)
else
  NEW_PATH := $(TOPIC)
endif

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

.PHONY: help list run new build vet test fmt fmt-check check tidy $(TOPICS) $(RUNNABLES)

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
	@printf '    %brun%b [topic]   run ./main.go (overview) if no topic, else a topic under examples/, exercises/, or a nested path\n' '$(C_GREEN)' '$(C_RESET)'
	@printf '    %bnew%b <topic>   scaffold <topic>/main.go (defaults to examples/)\n' '$(C_GREEN)' '$(C_RESET)'
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
	@printf '  %btopics in %s%b\n\n' '$(C_BOLD)' '$(CHAPTER)' '$(C_RESET)'
	@printf '  %bexamples%b\n' '$(C_BOLD)' '$(C_RESET)'
	@for t in $(RUNNABLES); do \
	  case "$$t" in examples/*) printf '    %s\n' "$$t";; esac; \
	done
	@printf '  %bexercises%b\n' '$(C_BOLD)' '$(C_RESET)'
	@for t in $(RUNNABLES); do \
	  case "$$t" in exercises/*) printf '    %s\n' "$$t";; esac; \
	done
	@printf '\n'

# --- run one topic (or the overview) ---------------------------------------
run:
	@if [ -z "$(RUN_TARGET)" ]; then \
	  if [ -z "$(TOPIC)" ]; then \
	    printf '  %bno topic given%b — try: make run <topic> (overview needs a root main.go)\n\n' '$(C_RED)' '$(C_RESET)'; \
	  else \
	    printf '  %bunknown topic:%b %s\n\n' '$(C_RED)' '$(C_RESET)' '$(TOPIC)'; \
	  fi; \
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
	@printf '  %bIn chapter %s · running%b %b%s%b\n\n' \
	  '$(C_DIM)' '$(CHAPTER)' '$(C_RESET)' '$(C_CYAN)' '$(RUN_TARGET)' '$(C_RESET)'
	@if [ "$(ANIM)" = "1" ]; then \
	  for f in '⠋' '⠙' '⠹' '⠸' '⠼' '⠴' '⠦' '⠧'; do \
	    printf '\r  %s warming up %s …' "$$f" '$(RUN_TARGET)'; \
	    sleep 0.05; \
	  done; \
	  printf '\r\033[K'; \
	fi
	@go run $(RUN_TARGET) $(ARGS) && printf '\n  %b✓ done%b\n' '$(C_GREEN)' '$(C_RESET)' \
	  || { printf '\n  %b✗ go run %s failed%b\n' '$(C_RED)' '$(RUN_TARGET)' '$(C_RESET)'; exit 1; }

# --- scaffold a new topic ---------------------------------------------------
new:
	@if [ -z "$(TOPIC)" ]; then \
	  printf '  %bno name given%b — try: make new <topic>\n' '$(C_RED)' '$(C_RESET)'; \
	  exit 1; \
	fi
	@if [ -e "$(NEW_PATH)" ]; then \
	  printf '  %b%s already exists%b\n' '$(C_RED)' '$(NEW_PATH)' '$(C_RESET)'; \
	  exit 1; \
	fi
	@mkdir -p "$(NEW_PATH)"
	@printf 'package main\n\nimport "fmt"\n\n// NOTE: what this example shows\nfunc main() {\n\tfmt.Println("%s")\n}\n' '$(TOPIC)' > "$(NEW_PATH)/main.go"
	@gofmt -w "$(NEW_PATH)/main.go"
	@printf '  %b✓%b created %s/main.go — run it with: make run %s\n' \
	  '$(C_GREEN)' '$(C_RESET)' '$(NEW_PATH)' '$(NEW_PATH)'

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
	@:
endif

# Same treatment for runnable paths under examples/ and exercises/, including
# nested ones, so make does not mistake them for files or directories.
ifneq ($(RUNNABLES),)
$(RUNNABLES):
	@if [ "$(MAKECMDGOALS)" = "$@" ]; then \
	  printf '  %b%s is a topic, not a target%b — try: make run %s\n' \
	    '$(C_DIM)' '$@' '$(C_RESET)' '$@'; \
	  exit 1; \
	fi
	@:
endif

# Swallow `make run defer-func`'s stray second goal, but still complain when an
# unknown target is the only thing asked for.
%:
	@if [ "$(MAKECMDGOALS)" = "$@" ]; then \
	  printf '  %bunknown target:%b %s — try: make help\n' '$(C_RED)' '$(C_RESET)' '$@'; \
	  exit 1; \
	fi
	@:
