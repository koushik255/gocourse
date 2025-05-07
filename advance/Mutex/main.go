package main

import (
	"fmt"
	"sync"
)

// mutex short of mutual synchronization, prevents go routines from accessing the same
// information at the same time. It avoids race conditions and data corruptions.
// Mutual exclusion is used to prevent multiple threads from accessing the same data at
// once.

type counter struct {
	mu    sync.Mutex
	count int // value that go routines will try to use
}

func (c *counter) increment() {
	c.mu.Lock()         // lock it at the start
	defer c.mu.Unlock() // c .count ++ comes inbetween the count and unloick
	// the defer comes to the end of the function
	c.count++
}

func (c *counter) getValue() int {
	c.mu.Lock()
	defer c.mu.Unlock()
	return c.count
}

func main() {
	var wg sync.WaitGroup
	counter := &counter{} // creates a new instance of the counter struct and returns pointer

	numGoroutine := 10

	// wg.Add(numGoroutine)
	for range numGoroutine {
		wg.Add(1)
		go func() {
			defer wg.Done()
			for range 1000 {
				counter.increment()
			}
		}()
	}

	wg.Wait() // wait for all the goroutines to finsih
	fmt.Printf("Final counter value: %d\n", counter.getValue())
}
