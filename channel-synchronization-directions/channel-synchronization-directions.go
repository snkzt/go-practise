package main

import (
	"fmt"
	"time"
)

// Synch
func worker(done chan bool) {
	fmt.Print("working...")
	time.Sleep(time.Second)
	fmt.Println("done")
	done <- true
}

// Directions
// When using channels as function param, we can specify if a
// channel is meant to only send or reveive values.
// ping function only accespts a channel for sending values.
// It would be a compile-time error to try to receive on this channel.
func ping(pings chan<- string, msg string) {
	pings <- msg
}

// The pong function accepts pings for receives and
// pongs for sends.
func pong(pings <-chan string, pongs chan<- string) {
	msg := <-pings
	pongs <- msg
}

func main() {
	done := make(chan bool, 1)
	// Start a worker goroutine, giving it the channel to notify on.
	go worker(done)
	// Block until it receives a notification from the worker on the channel.
	// If <-done is removed, the program would exit before the worker started.
	<-done

	pings := make(chan string, 1)
	pongs := make(chan string, 1)
	ping(pings, "passed message")
	pong(pings, pongs)
	fmt.Println(<-pongs)
}
