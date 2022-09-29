package main

import "fmt"

// Pass by value
func zeroValue(ival int) {
	ival = 0
}

// * takes an int pointer.
func zeroPointer(iptr *int) {
	*iptr = 0 // Dereference the pointer from its memory address to the current value at the address.
	// Assigning a value to a dereferenced pointer changes the value at the referenced adress.
}

func main() {
	i := 1
	fmt.Println("initial:", i) // Just printing i

	zeroValue(i)
	fmt.Println("zeroValue:", i) // 1 as copied value

	zeroPointer(&i)
	fmt.Println("zeroPointer:", i) // 0 as copied address

	fmt.Println("pointer:", &i) // Address of i
}
