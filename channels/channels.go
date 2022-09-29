package main

import "fmt"

// Channels are the pipes that connect concurrent goroutines.
// Send values into channels from one goroutine and receive those values into another goroutine.
func main() {
	// Create a new chain with make(chan val-type).
	// Channels are typed by the values they convey.
	messages := make(chan string)
	// Send a value into a channel useing the channel <- syntacx.
	// Here we send "ping" to the message channel above, from a new goroutine.
	go func() { messages <- "ping" }()

	// The <-channel syntax receives a value from the channel.
	// By default sends and receives block until both the sender and receiver are ready.
	msg := <-messages
	fmt.Println(msg)
}
