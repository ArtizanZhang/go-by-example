package main

import "fmt"

// By default channels are unbuffered,
// meaning that they will only accept sends (chan <-) if there is a corresponding receive (<- chan) ready to receive the sent value.
// Buffered channels accept a limited number of values without a corresponding receiver for those values.
func main() {

	// Here we make a channel of strings buffering up to 2 values
	messages := make(chan string, 2)

	// because this channel is buffered, we can send these values into the channel without a corresponding concurrent receive
	messages <- "buffered"
	messages <- "channel"

	// channel's buffering is up to 2 values, the third value will block, if no receiver to handle this , the program fatal error
	//messages <- "exceed" // fatal error: all goroutines are asleep - deadlock!

	fmt.Println(<-messages)
	fmt.Println(<-messages)

}
