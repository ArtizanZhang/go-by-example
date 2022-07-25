package main

import (
	"fmt"
	"time"
)

//  We often want to execute Go code at some point in the future, or repeatedly at some interval.
//	Go’s built-in timer and ticker features make both of these tasks easy. We’ll look first at timers and then at tickers.
func main() {

	timer1 := time.NewTimer(2 * time.Second)
	// The <-timer1.C blocks on the timer’s channel C until it sends a value indicating that the timer fired
	<-timer1.C
	fmt.Println("Timer1 fired") // equal to time.Sleep()

	// If you just wanted to wait, you could have used time.Sleep.
	// One reason a timer may be useful is that you can cancel the timer before it fires. Here’s an example of that.
	time2 := time.NewTimer(time.Second)
	go func() {
		<-time2.C
		fmt.Println("Timer2 fired")
	}()

	stopped := time2.Stop()
	if stopped {
		fmt.Println("Timer2 stopped")
	}

	// Give the timer2 enough time to fire, if it ever was going to, to show it is in fact stopped.
	time.Sleep(2 * time.Second)
}
