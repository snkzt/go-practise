package main

import (
	"fmt"
	"time"
)

// Rate Limiting is an important mechanism for controlling
// resource utilization and maintaining quality of service.
func main() {
	// requests is to limit handling of incoming requests.
	// These requests will off a channel of the same name.
	requests := make(chan int, 5)
	for i := 1; i <= 5; i++ {
		requests <- i
	}
	close(requests)
	// Regulator
	// This limiter channel will receive a value every 200 milliseconds.
	limiter := time.Tick(200 * time.Millisecond)
	// By blocking on a receive from the limiter channel before serving
	// each request, we limit ourselves to 1 request every 200 milliseconds.
	for req := range requests {
		<-limiter
		fmt.Println("request", req, time.Now())
	}
	// When want to allow short bursts of requests in our rate limiting scheme
	// while preserving the overall rate limit, it can be done by buffering
	// the limiter channel.
	// This burstyLimiter channel will allow bursts of up to 3 events.
	burstyLimiter := make(chan time.Time, 3)
	for i := 0; i < 3; i++ {
		burstyLimiter <- time.Now()
	}
	// Every 200 milliseconds it tries to add a new value to burstyLimiter, up to its limit of 3.
	go func() {
		for t := range time.Tick(200 * time.Millisecond) {
			burstyLimiter <- t
		}
	}()

	burstyRequests := make(chan int, 5)
	for i := 1; i <= 5; i++ {
		burstyRequests <- i
	}
	close(burstyRequests)
	for req := range burstyRequests {
		<-burstyLimiter
		fmt.Println("request", req, time.Now())
	}
}
