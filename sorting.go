package main

import (
	"fmt"
	"sort"
)

//  Go’s sort package implements sorting for builtins and user-defined types. We’ll look at sorting for builtins first.
func main() {
	strs := []string{"c", "a", "b"}

	// Note that sorting is in-place, so it changes the given slice and doesn’t return a new one.
	sort.Strings(strs)
	fmt.Println("Strings", strs)

	ints := []int{7, 1, 1}
	sort.Ints(ints)
	fmt.Println("Ints", ints)

	// We can also use sort to check if a slice is already in sorted order.
	sorted := sort.IntsAreSorted(ints)
	fmt.Println("Sorted ", sorted)
}
