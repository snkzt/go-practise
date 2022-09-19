package main

import "fmt"

func main() {
	i := 1

	// A single condition
	for i <= 3 {
		fmt.Println(i)
		i = i + 1
	}

	// Initial/condition/action after for loop
	for j := 7; j <= 9; j++ {
		fmt.Println(j)
	}

	// Without condition will loop until break out of it or return from the enclosing function.
	for {
		fmt.Println("loop until break")
		break
	}

	// Continue inside a for loop continues to the next iteration
	for n := 0; n <= 5; n++ {
		if n%2 == 0 {
			continue
		}
		fmt.Println(n)							
	}
}
