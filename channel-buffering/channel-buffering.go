package main

import "fmt"

// Channels are unbuffered: they will only accept sends (chan <-) if there is a corresponding
// receive (<- chan) ready to receive the sent value.
// Buffered channels accept a limited number of values without a corresponding receiver for those values.
// Sender is a function sending the channel, and the receiver is the function receiving it.
func main() {
	messages := make(chan string, 2)
	messages <- "buffered"
	messages <- "channel"

	fmt.Println(<-messages)
	fmt.Println(<-messages)
}
