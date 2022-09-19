package main

import "fmt"

func main() {

	// make(map[key type]value type)
	m := make(map[string]int)

	m["k1"] = 7
	m["k2"] = 13

	fmt.Println("map:", m)

	// Get value as arrays and slices with varianle = mapname[key]
	v1 := m["k1"]
	fmt.Println("v1:", v1)
	// The builtin len returns the number of key/value pairs
	fmt.Println("len:", len(m))

	// The builtin delete removes key/value pairs from a map.
	delete(m, "k2")
	fmt.Println("map:", m)

	// This is to retrieve and check if the value of the key exists in the map.
	// The optional second return value when getting a value from a map indicates
	// if the key was present in the map.
	// This can be used to disambiguate between missingkeys and keys with zero values
	// zero values like 0 or "". 
	_, present := m["k2"] // First return value ignored with blank identifier as we don't need it
	fmt.Println("present:", present)

	value, presence := m["k1"]
	fmt.Printf("value for k1: %v\nPresence: %v\n", value, presence)

	n := map[string]int{"foo": 1, "bar": 2}
	fmt.Println("map:", n)
}
