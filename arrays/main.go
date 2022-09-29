package main

import "fmt"

func main() {
	// Array a can hold 5 integeres.
	// The type of elements and the type are part of the array's type.
	// By default an array is zero-valued, which for integer means 0.
	var a [5]int
	fmt.Println("emp:", a)

	// We can assign a value to an index of an array with
	// array[index] = value
	// We also can get a value with
	// value = array[index]
	a[4] = 100
	fmt.Println("set:", a)
	fmt.Println("get:", a[4])

	// The builtin len returns the length of an array.
	fmt.Println("len:", len(a)) // 5

	// Declare and initialise an array at the same time
	b := [5]int{1, 2, 3, 4, 5}
	fmt.Println("declare and initialise:", b)

	// [[0,1,2],[0,1,2]]
	var twoDimensions [2][3]int
	for i := 0; i < 2; i++ {
		for j := 0; j < 3; j++ {
			twoDimensions[i][j] = i + j
		}
	}
	fmt.Println("2D:", twoDimensions)
}
