package main

import "fmt"

// In a previous example we saw how for and range provide iteration over basic data structures.
// We can also use this syntax to iterate over values received from a channel.

func main() {

	queue := make(chan string, 2)
	queue <- "one"
	queue <- "two"
	close(queue) // Without this phrase it will deadlock

	// This range iterates over each element as it’s received from queue. Because we closed the channel above, the iteration terminates after receiving the 2 elements.
	for elem := range queue {
		fmt.Println(elem)
	}
}
