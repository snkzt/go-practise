package main

import (
	"fmt"
	"time"
)

func main() {

	i := 2
	fmt.Print("Write ", i, " as ")
	switch i {
	case 1:
		fmt.Println("one")
	case 2:
		fmt.Println("two")
	case 3:
		fmt.Println("three")
	}

	// You can use comma to separate multiple expressions in the same case statement.
	// default case is optional.
	switch time.Now().Weekday() {
	case time.Saturday, time.Sunday:
		fmt.Println("It's the Weekend XD")
	default:
		fmt.Println("It's a weekday <3")
	}

	t := time.Now()
	// Switch without an expression is an alternative way to
	// express if/else logic.
	switch {
	case t.Hour() < 12:
		fmt.Println("It's time for coffee before noon")
	default:
		fmt.Println("It's time for tea after noon")
	}

	whatAmI := func(i interface{}) {
		// A type switch compares types instead of values.
		switch t := i.(type) {
		case bool:
			fmt.Println("I'm a boolean")
		case int:
			fmt.Println("I'm an integer")
		default:
			fmt.Printf("Need a DNA test for the type %T\n", t)
		}
	}

	whatAmI(true)
	whatAmI(7)
	whatAmI("Egg")
}

