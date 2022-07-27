package main

import (
	"fmt"
	"time"
)

// Rate limiting is an important mechanism for controlling resource utilization and maintaining quality of service.
// Go elegantly supports rate limiting with goroutines, channels, and tickers.

func main() {

	// First we’ll look at basic rate limiting. Suppose we want to limit our handling of incoming requests.
	// We’ll serve these requests off a channel of the same name.
	request := make(chan int, 5)
	for i := 1; i <= 5; i++ {
		request <- i
	}
	close(request)
	// This limiter channel will receive a value every 200 milliseconds. This is the regulator in our rate limiting scheme.
	limiter := time.Tick(200 * time.Millisecond)
	for req := range request {
		// By blocking on a receive from the limiter channel before serving each request, we limit ourselves to 1 request every 200 milliseconds.
		<-limiter
		fmt.Println("request", req, time.Now())
	}

	// We may want to allow short bursts of requests in our rate limiting scheme while preserving the overall rate limit.
	// We can accomplish this by buffering our limiter channel. This burstyLimiter channel will allow bursts of up to 3 events.
	burstyLimiter := make(chan time.Time, 3)
	for i := 0; i < 3; i++ {
		burstyLimiter <- time.Now()
	}

	go func() {
		for t := range time.Tick(200 * time.Millisecond) {
			burstyLimiter <- t
		}
	}()

	burstyReqs := make(chan int, 5)
	for i := 0; i < 5; i++ {
		burstyReqs <- i
	}
	close(burstyReqs)

	// For the second batch of requests we serve the first 3 immediately because of the burstable rate limiting, then serve the remaining 2 with ~200ms delays each.
	for burstyReq := range burstyReqs {
		<-burstyLimiter
		fmt.Println("REQUEST ", burstyReq, time.Now())
	}
}
