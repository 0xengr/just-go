// NOTE: map[T]struct{} is the set idiom. The empty struct occupies no memory,
// so a bool or int value would only waste space per key.
package main

import "fmt"

func main() {
	seen := map[string]struct{}{}
	for _, w := range []string{"go", "go", "rust", "go", "zig"} {
		seen[w] = struct{}{}
	}
	_, goSeen := seen["go"]
	_, cobolSeen := seen["cobol"]
	fmt.Println("unique:", len(seen), "go seen:", goSeen, "cobol seen:", cobolSeen)

	if _, dup := seen["go"]; dup {
		fmt.Println("go was already in the set")
	}

	nums := map[int]struct{}{1: {}, 2: {}, 3: {}}
	_, has := nums[4]
	fmt.Println("distinct numbers:", len(nums), "contains 4:", has)

	// iterating a set yields only keys
	for n := range nums {
		delete(nums, n) // deleting while ranging is allowed
	}
	fmt.Println("emptied by delete:", len(nums))
}
