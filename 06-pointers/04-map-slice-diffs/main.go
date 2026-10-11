package main

import (
	"fmt"
	"strings"
)

// NOTE: Maps and Slice differences while working with Pointers.
// Go passes every argument by value. What differs is what that value holds: an
// int is the number, a struct copies its fields, but a map and a slice are
// small handles onto data that lives somewhere else. Copying the handle is
// cheap, and both copies still reach the same underlying data. That single fact
// explains most of what feels surprising about passing maps and slices around.

type User struct {
	Name  string
	Email string
}

// Labels is a map field, and it matters: copying the User copies the map handle
// too, not the entries it points at.
type UserWithLabels struct {
	Name   string
	Labels map[string]string
}

// NOTE: a map value reaches the caller's data without any pointer. `m` is a copy
// of the handle, but the handle addresses the same hash table `user` does, so
// writing an entry is visible on both sides.

func setAge(user map[string]int) {
	user["age"] = 25
}

func runMapMutation() {
	fmt.Println("=== a map entry written inside a function is visible outside ===")

	user := map[string]int{"age": 20}

	setAge(user)

	fmt.Println("user after setAge:", user)
	fmt.Println("so no pointer was needed to change age:", user["age"])
}

// NOTE: reassigning the parameter is a different story. The new handle is local,
// exactly like reassigning any other parameter, so the caller keeps the original
// map and the replacement is invisible.

func replaceMap(user map[string]int) {
	user = map[string]int{"age": 99}
}

func runMapReassign() {
	fmt.Println("\n=== assigning a new map to the parameter stays local ===")

	user := map[string]int{"age": 20}

	replaceMap(user)

	fmt.Println("user after replaceMap:", user)
}

// NOTE: the trap. `data["name"]` on a missing key is not an error, it is the
// zero value of the value type, so a typo or a missing field compiles fine and
// fails quietly at runtime.

func createUserFromMap(data map[string]string) {
	fmt.Println("  name:", data["name"])
	fmt.Println("  email:", data["email"])
}

// NOTE: the same function against a struct. The compiler now knows the field
// names, so `data.Nmae` is a build error rather than an empty string.

type CreateUserInput struct {
	Name  string
	Email string
}

func createUserFromStruct(input CreateUserInput) {
	fmt.Println("  name:", input.Name)
	fmt.Println("  email:", input.Email)
}

func runStructVersusMap() {
	fmt.Println("\n=== a struct states the contract; a map only suggests it ===")

	// BUG: "username" is not "name", so this compiles and prints an empty
	// string. The compiler cannot see the difference between the two keys:
	// data := map[string]string{"username": "mohammad"}
	data := map[string]string{"username": "mohammad", "email": "user@example.com"}
	createUserFromMap(data)

	input := CreateUserInput{Name: "mohammad", Email: "user@example.com"}
	createUserFromStruct(input)
}

// NOTE: writing an existing element reaches the shared backing array, so the
// caller sees it.

func scaleFirst(nums []int) {
	nums[0] = nums[0] * 10
}

// NOTE: appending without returning the result does not. `nums` inside is a copy
// of the three-word descriptor, and append hands back a new descriptor with a
// longer length. Assigning it to the local copy cannot rewrite the caller's.

func appendDiscarding(nums []int) {
	nums = append(nums, 40)
}

// NOTE: this is the fix, and the shape of nearly every slice-building function in
// the stdlib. Return the slice, and let the caller assign the new descriptor.

func appendReturning(nums []int) []int {
	return append(nums, 40)
}

func runSliceVisibility() {
	fmt.Println("\n=== element writes are shared, length changes are not ===")

	nums := []int{10, 20, 30}

	scaleFirst(nums)
	fmt.Println("after scaleFirst (element write):", nums)

	appendDiscarding(nums)
	fmt.Println("after appendDiscarding, len unchanged:", nums, "len =", len(nums))

	nums = appendReturning(nums)
	fmt.Println("after appendReturning and assigning:", nums, "len =", len(nums))
}

// NOTE: a struct passed by value is not automatically deep-immutable. The Name
// string is copied and safe, but the Labels map is a handle, so both copies keep
// reaching the same entries. Owning the outer struct buys less than it looks.

func addLabel(user UserWithLabels, label string) {
	user.Labels[label] = "true"
}

func runShallowCopy() {
	fmt.Println("\n=== a struct by value is still only a shallow copy ===")

	user := UserWithLabels{Name: "mohammad", Labels: map[string]string{"beta": "yes"}}

	addLabel(user, "admin")
	fmt.Println("caller sees the new label:", user.Labels)

	// NOTE: the fields that need real isolation are the plain ones. Reassigning
	// user.Name inside the function would have stayed local.
}

// NOTE: maps earn their place when the keys are genuinely not knowable up front.
// Counting words is the honest case: no struct could name every possible key,
// because the keys come from the input itself.

func countWords(text string) map[string]int {
	counts := map[string]int{}

	for _, word := range strings.Fields(strings.ToLower(text)) {
		counts[word]++
	}

	return counts
}

func runGenuineMap() {
	fmt.Println("\n=== a map is right when the keys come from the data ===")

	counts := countWords("Go is go and Go is worth it")

	for _, word := range []string{"go", "is", "worth"} {
		fmt.Printf("  %-6s %d\n", word, counts[word])
	}

	// PERF: iterating a map has randomized order on purpose, so the loop above
	// walks a fixed slice of keys instead of ranging the map. Printing counts
	// straight from the map would change order every run and be untestable.
}

func main() {
	runMapMutation()
	runMapReassign()
	runStructVersusMap()
	runSliceVisibility()
	runShallowCopy()
	runGenuineMap()
}
