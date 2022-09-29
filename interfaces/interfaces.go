package main

import (
	"fmt"
	"math"
)

// Interfaces are named collections of method signatures.
// The signature of a method consists of the name of the method
// and the description of its parameters
//
//	toUpperCase()
//	Println(s string)

type geometry interface {
	area() float64
	perim() float64
}
type rectangle struct {
	width, height float64
}
type circle struct {
	radius float64
}

// To implement an interface in Go, implement all the methods in the interface.
// For rectangle
func (r rectangle) area() float64 {
	return r.width * r.height
}
func (r rectangle) perim() float64 {
	return 2*r.width + 2*r.height
}

// For circle
func (c circle) area() float64 {
	return math.Pi * c.radius * c.radius
}
func (c circle) perim() float64 {
	return 2 * math.Pi * c.radius
}

// If a variable has an interface type, geometry in this case, we can call
// methods that are in the named interface, area and perim in this case.
func measure(g geometry) {
	fmt.Println(g)
	fmt.Println(g.area())
	fmt.Println(g.perim())
}

func main() {
	rectangle := rectangle{width: 3, height: 4}
	circle := circle{radius: 5}
	// The circle and rectangle struct types both implement the
	// geometry interface so we can use instances of those structs
	// as arguments to measure.
	measure(rectangle)
	measure(circle)
}
