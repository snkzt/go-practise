package main

import "fmt"

// Go supports embedding of structs and interfaces to express a more
// seamless composition of types.
type base struct {
	num int
}

func (b base) describe() string {
	return fmt.Sprintf("base with num=%v", b.num)
}

// An embedding looks like a field without a name.
type container struct {
	base // embedding
	str  string
}

func main() {
	// When creating structs with literals, we have to initialize the
	// embedding explicitly; here the embedded type serves as the field name.
	co := container{
		base: base{
			num: 1, // base's fields
		},
		str: "some name",
	}
	fmt.Printf("co={num: %v, str: %v}\n", co.num, co.str)
	// struct name.embedding type.field
	fmt.Println("also num:", co.base.num)
	fmt.Println("also num:", co.base.num)
	// Describe() only has base as method but as base is embedded in container
	// The field num in the base can be accessed from container directly.
	fmt.Println("describe:", co.describe())

	type describer interface {
		describe() string
	}
	var d describer = co
	fmt.Println("describer:", d.describe())
}
