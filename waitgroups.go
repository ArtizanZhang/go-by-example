package main

import (
	"fmt"
	"sync"
	"time"
)

func main() {

	var wg sync.WaitGroup

	for i := 1; i < 5; i++ {
		// This WaitGroup is used to wait for all the goroutines launched here to finish.
		// Note: if a WaitGroup is explicitly passed into functions, it should be done by pointer.
		wg.Add(1)

		// Avoid re-use of the same i value in each goroutine closure. See the FAQ for more details
		i := i // create a new 'i'.   can also use anonymous function . refer: https://go.dev/doc/faq#closures_and_goroutines

		go func() {
			defer wg.Done() // Wrap the worker call in a closure that makes sure to tell the WaitGroup that this worker is done.
			// This way the worker itself does not have to be aware of the concurrency primitives involved in its execution.

			work2(i)
		}()
	}

	wg.Wait() // Block until the WaitGroup counter goes back to 0; all the workers notified they’re done.

}
func work2(id int) {
	fmt.Println("worker", id, "started job")
	time.Sleep(time.Second)
	fmt.Println(" worker", id, "finished job")
}
