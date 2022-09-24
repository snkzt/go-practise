package main

import (
	"fmt"
	"sync"
)

// Package name usually being the repo name the code resides.
// Mutex is to access data across multiple goroutines safely.
// Add mutex to synchronise access to a map of counters which we
// want to update concurrently from multiple goroutines.
// Mutex must not be copied so if this Container struct is passed around,
// it should be done by pointer.
type Container struct {
	mutex    sync.Mutex
	counters map[string]int
}

func (c *Container) inc(name string) {
	// Lock the mutex before accessing counters
	c.mutex.Lock()
	// Unlock it at the end of the function
	defer c.mutex.Unlock()
	c.counters[name]++
}

func main() {
	c := Container{
		// The zero value of a mutex is usable as-is, so no initialisation is required.
		counters: map[string]int{"a": 0, "b": 0},
	}
	var wg sync.WaitGroup
	doIncrement := func(name string, n int) {
		for i := 0; i < n; i++ {
			c.inc(name)
		}
		wg.Done()
	}

	wg.Add(3)
	// Run several goroutines concurrently; they all access the same Container,
	// and two of them access the same counters.
	go doIncrement("a", 10000)
	go doIncrement("a", 10000)
	go doIncrement("b", 10000)

	wg.Wait() // Wait goroutines to finish
	fmt.Println(c.counters)
}
