package main

import (
	"fmt"
	"log"
	"net/http"
	"runtime/debug"
)

// NOTE: Recovering inside an HTTP handler
// * A real API cannot afford to die because one handler panicked. net/http
// * already recovers handler panics per request, but its response is a generic
// * one. A recovery middleware puts a boundary around every handler so you can
// * log the panic with a stack trace and return your own 500 response.

// NOTE: recoveryMiddleware wraps a handler in a deferred recover.
// * The defer is registered before the handler runs, so it is the last thing
// * to run when the handler panics. It logs the panic value and stack trace,
// * writes a 500 response, and swallows the panic so the server keeps serving.
func recoveryMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		defer func() {
			if recovered := recover(); recovered != nil {
				log.Printf("PANIC: %v\n%s", recovered, debug.Stack())
				http.Error(w, "internal server error", http.StatusInternalServerError)
			}
		}()

		next.ServeHTTP(w, r)
	})
}

func userHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Println("Processing user request...")

	// Simulate an unexpected programming bug.
	panic("unexpected nil pointer")
}

func healthHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(w, "API is healthy")
}

func main() {
	mux := http.NewServeMux()
	mux.HandleFunc("GET /users", userHandler)
	mux.HandleFunc("GET /health", healthHandler)

	addr := ":8080"
	log.Printf("server listening on %s", addr)
	if err := http.ListenAndServe(addr, recoveryMiddleware(mux)); err != nil {
		log.Fatal(err)
	}
}

// NOTE: Try it from another terminal while the server runs:
// * curl -i localhost:8080/health   → 200 OK, "API is healthy"
// * curl -i localhost:8080/users    → 500 Internal Server Error
// * The panic and its stack trace appear in the server's log, and /health
// * still answers afterwards: one request failed, the process survived.
