package main

import (
	"fmt"
	"os"
)

type point struct {
	x, y int
}

func main() {
	p := point{x: 1, y: 2}
	fmt.Println(p)

	// v part for %v is called printing verbs designed to format general Go values.
	fmt.Printf("struct1: %v\n", p)
	// For struct, %+v include the struct's field names.
	fmt.Printf("struct2: %+v\n", p)
	// The %#v prints a Go syntax representation of the value.
	// the source code snippet that would produce that value.
	fmt.Printf("struct3: %#v\n", p)
	// T to print type of a value
	fmt.Printf("type: %T\n", p)
	// t for bool
	fmt.Printf("bool: %t\n", true)
	// d for standard base-10 integer formatting
	fmt.Printf("int: %d\n", 123)
	// b for binary representation
	fmt.Printf("bin: %b\n", 14)
	// c to print the character correspondin to the given integer
	fmt.Printf("char: %c\n", 33)
	// x to print hex encoding of the given integer
	fmt.Printf("hex: %x\n", 456)
	// f for decimal formatting
	fmt.Printf("float1: %f\n", 78.9)
	// e / E for format the float in scientific notation
	fmt.Printf("float2: %e\n", 123400000.0)
	fmt.Printf("float3: %E\n", 123400000.0)
	// s for basic string printing
	fmt.Printf("str1: %s\n", "\"string\"")
	// To double-quote strings as in Go source
	fmt.Printf("str2: %q\n", "\"string\"")
	// %x renders the string in base-16 with two output
	// characters per byte of input
	fmt.Printf("str3: %x\n", "hex this")
	// %p for a representation of a pointer
	fmt.Printf("pointer: %p\n", &p)
	// To specify the width of an integer, use a number after the % in the verb. By default the result will be right-justified and padded with spaces.
	// To left-justify, use the - flag.
	fmt.Printf("width1: |%6d|%6d|\n", 12, 345)
	fmt.Printf("width2: |%6.2f|%6.2f|\n", 1.2, 3.45)
	fmt.Printf("width3: |%-6.2f|%-6.2f|\n", 1.2, 3.45)
	fmt.Printf("width4: |%6s|%6s|\n", "foo", "b")
	fmt.Printf("width5: |%-6s|%-6s|\n", "foo", "b")
	s := fmt.Sprintf("sprintf: a %s", "string")
	// Printf, which prints the formatted string to os.Stdout
	// Sprintf formats and returnes a string without printing it anywhere.
	fmt.Println(s)
	// Fprintf can format+print to io.Writers other than os.Stdout.
	fmt.Fprintf(os.Stderr, "io: an %s\n", "error")
}
