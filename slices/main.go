package main

import "fmt"

func main() {

	// Slices are typed ONLY by the elements they contain.
	// To create an empty slice with non-zero value, use the builtin make.
	// Below is to make a slice of strings of length 3 with zero-valued.
	s := make([]string, 3)
	fmt.Println("emp:", s)

	s[0] = "a"
	s[1] = "b"
	s[2] = "c"
	// Set and get are the same as arrays.
	fmt.Println("set:", s)
	fmt.Println("get:", s[2])

	// len returns the length of the slice as it does for arrays
	fmt.Println("len", len(s))

	// Append is a builtin operation which array not support.
	// This returns a slice containing new value which is appened.
	s = append(s, "d")        // s is now len of 4
	s = append(s, "e", "f")   // s is now len of 6
	fmt.Println("append:", s) // s[a b c d e f]

	// slice can be copied with copy.
	// C is an empty slice with the same length as s and
	// copy values of s into c.
	c := make([]string, len(s))
	copy(c, s)
	fmt.Println("copy:", c)

	// slice[low:high] gets the value for the range.
	// low is included but high is not included.
	l := s[2:5]            // s[2,3,4]
	fmt.Println("sl1:", l) // s[c d e]

	// slice[:high] gets everything in the slice other than the last element.
	l = s[:5] // l == [a b c d e]
	fmt.Println("sl2:", l)

	// slice[low:] gets everything equals to and bigger than the low element.
	l = s[2:] // s[c d e f]
	fmt.Println("sl3:", l)

	// Declare and initialise a slice in a single line is the same as arrays.
	t := []string{"g", "h", "i"}
	fmt.Println("declare and initialise:", t)

	//Declare empty two dementional, length of 3 and type int slice
	twoDimensions := make([][]int, 3)
	for i := 0; i < 3; i++ {
		innerLength := i + 1
		twoDimensions[i] = make([]int, innerLength)
		for j := 0; j < innerLength; j++ {
			twoDimensions[i][j] = i + j
		}
	}
	fmt.Println("2d: ", twoDimensions)
}
