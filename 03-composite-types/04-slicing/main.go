// NOTE: a[i:j] leaves the capacity running to the end of the array, a[i:j:k]
// stops it at k. copy moves min(len(dst), len(src)) elements and never grows.
package main

import "fmt"

func main() {
	all := []int{0, 1, 2, 3, 4, 5, 6, 7}

	middle := all[2:5] // len 3, cap 6: everything after index 2
	capped := all[2:5:5]
	fmt.Println("middle:", middle, "len:", len(middle), "cap:", cap(middle))
	fmt.Println("capped:", capped, "len:", len(capped), "cap:", cap(capped))

	fmt.Println("tail:", all[5:], "head:", all[:2])

	// BUG: an out-of-range index panics, and so does a high bound past cap.
	// all[2:99] would panic: slice bounds out of range

	dst := make([]int, 3)
	n := copy(dst, all[6:])
	fmt.Println("copy moved:", n, "dst:", dst)

	clear(dst) // Go 1.21+: zero out in place
	fmt.Println("cleared:", dst)

	fmt.Println("truncated copy:", copy(dst, all)) // destination is smaller
	fmt.Println("after copying all:", dst)
}
