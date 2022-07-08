package main

import (
	"fmt"
	"time"
)

// channel are the pipes that connect concurrent goroutines. you can send values into channels from one goroutine and receive those values into another goroutine

func main() {

	messages := make(chan string) // Create a new channel with make(chan val-type). Channels are typed by the values they convey.

	go func() {
		time.Sleep(time.Second * 10) // the statement "msg := <-messages" below will block
		messages <- "ping"
	}()

	// By default sends and receives block until both the sender and receiver are ready.
	// This property allowed us to wait at the end of our program for the "ping" message without having to use any other synchronization.
	msg := <-messages

	fmt.Println(msg)

}
