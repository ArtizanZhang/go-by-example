package main

import (
	"fmt"
	"sync"
)

// In the previous example we saw how to manage simple counter state using atomic operations.
// For more complex state we can use a mutex to safely access data across multiple goroutines.

// Container holds a map of counters; since we want to update it concurrently from multiple goroutines, we add a Mutex to synchronize access.
type Container struct {
	mu       sync.Mutex
	counters map[string]int
}

//  Note that mutexes must not be copied, so if this struct is passed around, it should be done by pointer.
func (c *Container) inc(name string) {
	c.mu.Lock()
	defer c.mu.Unlock() // without mutexes,  fatal error: concurrent map writes

	c.counters[name]++
}

func main() {

	c := Container{
		counters: map[string]int{"a": 0, "b": 0},
	}

	var wg sync.WaitGroup

	doInc := func(name string, n int) {
		for i := 0; i < n; i++ {
			c.inc(name)
		}
		wg.Done()
	}

	// Run several goroutines concurrently; note that they all access the same Container, and two of them access the same counter.
	wg.Add(3)
	go doInc("a", 10000)
	go doInc("a", 10000)
	go doInc("b", 10000)

	wg.Wait()

	fmt.Println(c.counters)
}
