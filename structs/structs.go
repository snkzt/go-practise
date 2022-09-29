package main

import "fmt"

// Structs are typed collections of fields
type person struct {
	name string
	age  int
}

func newPerson(name string) *person {
	p := person{name: name}
	p.age = 42
	return &p
}

func main() {
	fmt.Println(person{"Bob", 20})              // New person struct
	fmt.Println(person{name: "Alice", age: 30}) // Initialising person struct with value
	fmt.Println(person{name: "Fred"})           // Initialising and omitted field will be zero-valued (0)
	fmt.Println(&person{name: "Ann", age: 40})  // Pointer to the struct, the struct itself will be initialised with these values.
	fmt.Println(newPerson("John"))

	s := person{name: "Sean", age: 50}
	fmt.Println(s.name)
	sp := &s
	fmt.Println(sp.age)
	sp.age = 51
	fmt.Println(sp.age)
}
