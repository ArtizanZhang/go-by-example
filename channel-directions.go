package main

import "fmt"

// When using channels as function parameters, you can specify if a channel is meant to only send or receive values.
// This specificity increases the type-safety of the program.

// This ping function only accepts a channel for sending values. It would be a compile-time error to try to receive on this channel.
func ping(pings chan<- string, msg string) {
	pings <- msg

	//fmt.Println(<-pings) // Invalid operation: <-pings (receive from the send-only type chan<- string)

}

// The pong function accepts one channel for receives (pings) and a second for sends (pongs).
func pong(ping <-chan string, pongs chan<- string) {
	msg := <-ping
	pongs <- msg
}

func main() {
	pings := make(chan string, 1)
	pongs := make(chan string, 1)
	ping(pings, "passed message")
	pong(pings, pongs)
	fmt.Println(<-pongs)
}
