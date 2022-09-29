package main

import "fmt"

// Closing a channel indicates that no more values will be
// sent on it. This is useful to communicate completion to
// the channel's receivers.

func main() {
	jobs := make(chan int, 5)
	done := make(chan bool)
	// Worker goroutine which gets job repeatedly receives
	// jobs from jobs channel. More value will be false if jobs has been
	// closed after all jobs done and all values in the channel have already been received.
	go func() {
		for {
			j, more := <-jobs
			if more {
				fmt.Println("received job", j)
			} else {
				fmt.Println("received all jobs")
				done <- true
				return
			}
		}
	}()

	// Send 3 jobs to the worker over the jobs channel, then close it.
	for j := 1; j <= 3; j++ {
		jobs <- j
		fmt.Println("sent job", j)
	}
	close(jobs)
	fmt.Println("sent all jobs")
	<-done
}
