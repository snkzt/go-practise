package main

import "fmt"

// A closure is a persistent scope which holds on to local
// variables even after the code execution has moved out of that block.
// Go  supports anonymous functions to form a closure.
func intSeq() func() int {
	i := 0
	return func() int {
		i++
		return i
	}
}

func main() {
	nextInt := intSeq()
	fmt.Println(nextInt()) // 1
	fmt.Println(nextInt()) // 2
	fmt.Println(nextInt()) // 3

	newInts := intSeq()
	fmt.Println(newInts()) // 1

}
