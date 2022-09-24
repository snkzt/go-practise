package main

import (
	"fmt"
)

type rectangle struct {
	width, height int
}

// A method is a function with an extra receiver argument
// The receiver sits between the func keyword and the method name.

// Pointer type method
// Use this type to avoid copying on method calls or
// to allow the method to mutate the receiving struct
func (r *rectangle) area() int {
	return r.width * r.height
}

// Value receiver type method
func (r rectangle) perim() int {
	return 2*r.width + 2*r.height
}

func main() {
	r := rectangle{width: 10, height: 5}
	fmt.Println("area:", r.area())   // 50 mutating the struct
	fmt.Println("perim:", r.perim()) // 30 initialised a new struct

	rp := &r
	fmt.Println("area:", rp.area())   // 50
	fmt.Println("perim:", rp.perim()) // 30

	r = rectangle{height: 5}
	fmt.Println("area:", r.area())                     // 0 mutating the struct
	fmt.Println("Struct is now width:0, height:5", rp) // Because rectangle is dereferenced and mutated above
}
