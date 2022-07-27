package main

import (
	"fmt"
	"sync"
	"sync/atomic"
)

// The primary mechanism for managing state in Go is communication over channels. We saw this for example with worker pools. There are a few other options for managing state though.
//
// Here we’ll look at using the sync/atomic package for atomic counters accessed by multiple goroutines.
func main() {

	var ops uint64 // We’ll use an unsigned integer to represent our (always-positive) counter.

	var wg sync.WaitGroup //  WaitGroup will help us wait for all goroutines to finish their work

	// We’ll start 50 goroutines that each increment the counter exactly 1000 times.
	for i := 0; i < 50; i++ {
		wg.Add(1)

		go func() {
			for c := 0; c < 1000; c++ {

				// To atomically increment the counter we use AddUint64, giving it the memory address of our ops counter with the & syntax.
				atomic.AddUint64(&ops, 1)
			}
			wg.Done() // !!! don't forget this !!!
		}()

		//fmt.Println(atomic.LoadUint64(&ops)) // Reading atomics safely while they are being updated is also possible, using functions like atomic.LoadUint64.
	}

	wg.Wait()

	fmt.Println("ops:", ops) //  It’s safe to access ops now because we know no other goroutine is writing to it.
}

/**



We expect to get exactly 50,000 operations.

Had we used the non-atomic[[ ops++ ]]to increment the counter, we’d likely get a different number, changing between runs,

because the goroutines would interfere with each other. Moreover, we’d get data race failures when running with the -race flag.

*/