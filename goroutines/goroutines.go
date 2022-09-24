package main

import (
	"fmt"
	"time"
)

func f(from string) {
	for i := 0; i < 3; i++ {
		fmt.Println(from, ":", i)
	}
}

func main() {
	// This run first
	f("direct")
	// Goroutines are running concurrently by the Go runtime.
	go f("goroutine")
	go func(msg string) {
		fmt.Println(msg)
	}("going")

	time.Sleep(time.Second) // Run tne print below with a timelag
	fmt.Println("done")
}
