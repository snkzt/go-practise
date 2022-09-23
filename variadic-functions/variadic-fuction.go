package main

import "fmt"

// Valiadic function can be called with any number of tailing arguments.
func sum(nums ...int) {
	fmt.Print(nums, " ")
	total := 0

	for _, num := range nums {
		total += num
	}
	fmt.Println(total)
}

func main() {
	sum(1, 2)
	sum(1, 2, 3)

	nums := []int{1, 2, 3, 4}
	fmt.Println(len(nums))
	// If you already have multiple args in a slice, apply them
	// to a variadic function using func(slice...) like below
	sum(nums...)
}
