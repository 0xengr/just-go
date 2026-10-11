// NOTE: := always declares in the current block, so it can quietly shadow an
// outer variable instead of updating it. Use = when you mean to reuse.
package main

import "fmt"

func main() {
	fuel := 10
	if fuel > 5 {
		fmt.Println("before:", fuel)
		// WARNING: this declares a brand-new fuel that shadows the outer one
		// and dies with the if body, so the outer fuel is never updated.
		fuel := 5
		fmt.Println("inside:", fuel)
	}
	fmt.Println("after:", fuel) // outer fuel never changed

	if fuel > 5 {
		// := reuses only names from the same block, so fuel shadows again
		// even though pit is the only truly new variable here.
		fuel, pit := 5, 20
		fmt.Println("shadowed pair:", fuel, pit)
	}
	fmt.Println("outer still:", fuel)

	// The fix is plain assignment when the variable already exists:
	fuel = 5
	fmt.Println("reassigned:", fuel)

	// WARNING: the same trap applies to predeclared names (true, len, ...) and
	// to imports. A local called fmt or len would hide the package or builtin
	// for the rest of its block, so never reuse those names. go vet stays
	// silent about this; a shadowing linter is the tool that catches it.
}
