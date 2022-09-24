package main

import (
	"fmt"
	"math/rand"
	"sync/atomic"
	"time"
)

// To sychronise access to shared state across multiple goroutines,
// this is another option using built-in sync features of
// goroutine and channels.
// This achieves sharing memory by communicating and gaving each piece
// of data owned by exactly 1 goroutine.

type readOperation struct {
	key      int
	response chan int
}
type writeOperation struct {
	key      int
	value    int
	response chan bool
}

func main() {
	var readOps uint64
	var writeOps uint64
	// Setup 2 channels with each types struct
	reads := make(chan readOperation)
	writes := make(chan writeOperation)

	go func() {
		var state = make(map[int]int)
		for {
			select {
			case read := <-reads:
				read.response <- state[read.key]
			case write := <-writes:
				state[write.key] = write.value
				write.response <- true
			}
		}
	}()

	for r := 0; r < 10; r++ {
		go func() {
			for {
				read := readOperation{
					key:      rand.Intn(5),
					response: make(chan int)}
				reads <- read
				<-read.response
				atomic.AddUint64(&readOps, 1)
				time.Sleep(time.Millisecond)
			}
		}()
	}

	for w := 0; w < 10; w++ {
		go func() {
			for {
				write := writeOperation{
					key:      rand.Intn(5),
					value:    rand.Intn(100),
					response: make(chan bool)}
				writes <- write
				<-write.response
				atomic.AddUint64(&writeOps, 1)
				time.Sleep(time.Millisecond)
			}
		}()
	}

	time.Sleep(time.Second)

	readOpsFinal := atomic.LoadUint64(&readOps)
	fmt.Println("readOps:", readOpsFinal)
	writeOpsFinal := atomic.LoadUint64(&writeOps)
	fmt.Println("writeOps:", writeOpsFinal)
}
