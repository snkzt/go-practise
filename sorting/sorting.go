package main

import (
	"fmt"
	"sort"
)

func main() {
	strings := []string{"c", "a", "b"}
	// sort.Strings sort things in-places.
	// It changes the given slice and not return a new one.
	sort.Strings(strings)
	fmt.Println("Strings:", strings) // {a b c} forever if not re-sorted again

	ints := []int{7, 2, 4}
	//
	sort.Ints(ints)
	fmt.Println("Ints: ", ints)

	// We can use sort to check if a slice is already sorted.
	s := sort.IntsAreSorted(ints)
	fmt.Println("Sorted:", s)
}
