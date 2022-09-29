package main

import "fmt"

// Valiadic function can be called with any number of tailing arguments.
// Use this type of function when you want to deal with an arbitrary number of arguments
func sum(nums ...int) {
	fmt.Print(nums, " ")
	total := 0

	// With in the function, the type nums is equivalent to []int which is
	// the argument type used for the variadic function
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
	sum(nums...)
}
