package main

import (
	"fmt"
	"time"
)

func f(from string) {
	for i := 0; i < 3; i++ {
		fmt.Println(from, ":", i)
	}
}

func main() {
	// call func in usual way. running it synchronously.
	f("direct")

	// to invoke func in a goroutine, use go f(s).  this new goroutine will execute concurrently with the calling one .
	go f("goroutine")

	// you can also start a goroutine for an anonymous funciton call
	go func(msg string) {
		fmt.Println("msg")
	}("going")

	// Our two function calls are running asynchronously in separate goroutines now. Wait for them to finish (for a more robust approach, use a WaitGroup).
	time.Sleep(time.Second)
	fmt.Println("done")
}

//  When we run this program, we see the output of the blocking call first,
//  then the output of the two goroutines. The goroutines’ output may be interleaved, because goroutines are being run concurrently by the Go runtime.

/*  output
direct : 0
direct : 1
direct : 2
goroutine : 0
goroutine : 1
goroutine : 2
msg
done
*/
