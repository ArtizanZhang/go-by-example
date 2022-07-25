package main

import "fmt"

// Closing a channel indicates that no more values will be sent on it. This can be useful to communicate completion to the channel’s receivers.
func main() {

	jobs := make(chan int, 5)
	done := make(chan bool) // blocking channel

	go func() {
		for {
			//  In this special 2-value form of receive, the more value will be false if jobs has been closed and all values in the channel have already been received.
			j, more := <-jobs
			if more { // We use this to notify on done when we’ve worked all our jobs.
				fmt.Println("received job", j)
			} else {
				fmt.Println("received all jobs")
				done <- true
				return
			}
		}
	}()

	for j := 1; j <= 3; j++ {
		jobs <- j
		fmt.Println("sent job", j)
	}
	close(jobs) // This sends 3 jobs to the worker over the jobs channel, then closes it.
	fmt.Println("sent all jobs")

	<-done // We await the worker using the synchronization approach we saw earlier
}

/*  output

sent job 1
received job 1
sent job 2
received job 2
sent job 3
received job 3
sent all jobs
received all jobs
*/
