package main

import (
	"fmt"
	"time"
)

// Timeouts help to control execution time
// Programms that connects to external resources benefits from this.
func main() {
	// The channel is buffered so the send in the goroutine is nonblocking.
	c1 := make(chan string, 1)
	go func() {
		time.Sleep(2 * time.Second)
		c1 <- "resule 1"
	}()

	// The select implementing a timeout after 1s.
	// Since select proceeds with the first receive that's ready,
	// it will take timeout case if the operation takes more than the allwoed 1s.
	// In this case it takes timeout because sending message to the first channel
	// takes 2 seconds.
	select {
	case res := <-c1:
		fmt.Println(res)
	case <-time.After(1 * time.Second):
		fmt.Println("timeout 1")
	}

	c2 := make(chan string, 1)
	go func() {
		time.Sleep(2 * time.Second)
		c2 <- "result 2"
	}()
	select {
	case res := <-c2:
		fmt.Println(res)
	case <-time.After(3 * time.Second):
		fmt.Println("timeout 2")
	}
}
