package main

import (
	"fmt"
	"time"
)

// Timers are for when you want to do something once in the future
// - tickers are for when you want to do something repeatedly at regular intervals.
// Here’s an example of a ticker that ticks periodically until we stop it

func main() {

	ticker := time.NewTicker(500 * time.Millisecond)
	done := make(chan bool)

	// Tickers use a similar mechanism to timers: a channel that is sent values.
	// Here we’ll use the select builtin on the channel to await the values as they arrive every 500ms.
	go func() {
		for {
			select {
			case <-done:
				return
			case t := <-ticker.C:
				fmt.Println("Tick at", t)
			}
		}
	}()

	// Tickers can be stopped like timers. Once a ticker is stopped it won’t receive any more values on its channel. We’ll stop ours after 1600ms.
	time.Sleep(1600 * time.Millisecond)
	ticker.Stop()

	done <- true
	fmt.Println("Ticker Stopped")
}

/**
When we run this program the ticker should tick 3 times before we stop it.

Tick at 2022-07-25 10:59:57.446463 +0800 CST m=+0.500520985
Tick at 2022-07-25 10:59:57.947653 +0800 CST m=+1.001701597
Tick at 2022-07-25 10:59:58.446564 +0800 CST m=+1.500604797
Ticker Stopped

*/
