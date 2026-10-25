package main

import (
	"os"
)

// NOTE: Panic
// * In other languages exceptions are normal; in Go a panic is not.
// * A panic crashes the program.
// * A panic is triggered by the developer or by the runtime.
// * In Go, an unhandled error does not crash the program, but an unhandled
// * panic does.
// * Panic is used in rare cases, but it is useful in some scenarios.

// NOTE: We use panic based on the condition in our codebase.
// * For example, we may write code to check environment variables and crash
// * the program when one is missing.

// NOTE: This code triggers a panic from the developer.
func mustLoadEnv() {
	port := os.Getenv("PORT")
	if port == "" {
		panic("missing PORT check your .env file")
	}
	databaseURL := os.Getenv("DATABASE_URL")
	if databaseURL == "" {
		panic("missing DATABASE_URL check your .env file")
	}
}

func main() {

	// BUG: this loop reads past the end of nums and panics (add "fmt" to the
	// imports first):
	// nums := []int{1, 2, 3, 4, 5}
	// for i := 0; i <= 10; i++ {
	// 	fmt.Println(nums[i])
	// }
	// output:
	// 1
	// 2
	// 3
	// 4
	// 5
	// panic: runtime error: index out of range [5] with length 5

	// goroutine 1 [running]:
	// main.main()
	// 	/home/aman/Coding/golang/just-go/09-errors/examples/04-panic/main.go:38 +0xad
	// exit status 2

	// NOTE: In the crash above, 1 to 5 print before the panic because those
	// * indexes are valid.

	// BUG: uncommenting the next line panics when PORT or DATABASE_URL is
	// unset:
	// mustLoadEnv()
}
