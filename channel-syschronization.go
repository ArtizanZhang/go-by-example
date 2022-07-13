package main

import (
	"fmt"
	"time"
)

// We can use channels to synchronize execution across goroutines.
// Here’s an example of using a blocking receive to wait for a goroutine to finish.
// When waiting for multiple goroutines to finish, you may prefer to use a WaitGroup.
// https://gobyexample.com/waitgroups

func worker(done chan bool) {
	fmt.Println("working...")
	time.Sleep(time.Second)
	fmt.Println("done")

	done <- true
}

func main() {
	done := make(chan bool, 1) // make returns a pointer

	go worker(done)

	<-done // Block until we receive a notification from the worker on the channel.
}
