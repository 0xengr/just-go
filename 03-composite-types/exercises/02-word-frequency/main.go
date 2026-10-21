package main

import (
	"fmt"
	"sort"
	"strings"
)

func main() {
	text := "Go is go and Go is worth it"
	counts := map[string]int{}
	for _, word := range strings.Fields(strings.ToLower(text)) {
		counts[word]++
	}
	keys := make([]string, 0, len(counts))
	for key := range counts {
		keys = append(keys, key)
	}
	sort.Strings(keys)
	for _, key := range keys {
		fmt.Printf("%s=%d ", key, counts[key])
	}
	fmt.Println()
}
