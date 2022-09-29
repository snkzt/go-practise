package main

import (
	"fmt"
	"time"
)

func main() {
	c1 := make(chan string)
	c2 := make(chan string)

	go func() {
		time.Sleep(1 * time.Second)
		c1 <- "one"
	}()

	go func() {
		time.Sleep(2 * time.Second)
		c2 <- "two"
	}()

	// Select waits on multiple channel operations.
	// Each channel will recevie a value after some amout of time,
	// to simulate like the case of operations executing in concurrent goroutines.
	// In this case, select will await both of these values simultaneously,
	// printing each one as it arrives.
	for i := 0; i < 2; i++ {
		select {
		case msg1 := <-c1:
			fmt.Println("Received", msg1)
		case msg2 := <-c2:
			fmt.Println("Recevied", msg2)
		}
	}
}
