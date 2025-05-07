package main

import (
	"fmt"
	"sync"
	"sync/atomic"
)

// manage and count resources in a safe way using go routine,
// basically just a much better mutex
// they are also quicker than mutexes because they are not using locks
// atomic operations are operations that are indivisible, they are not interrupted by other operations
// good for sharing resources between goroutines
// uses memory barriers to prevent other operations from being interleaved
// if you dont use a atomic counter then you will have alot of issues with race conditions

// race condition is when two or more threads try to access the same memory location at the same time
// this can cause the value to be incorrect
// they also create unpredictable behavior so that is #notgood blud
// use these for simple operations that are not complex
// for complex operations use mutexes or channels

type AtomicCounter struct {
	count int64
}

func (ac *AtomicCounter) Increment() {
	atomic.AddInt64(&ac.count, 1) // because the type is of int64, we need to p
	//pass the address of the count to acces the count field/ memory address of the struct
}

func (ac AtomicCounter) GetValue() int64 {
	return atomic.LoadInt64(&ac.count)
	// value receiver since it is not modifying the struct
}

func main() {
	var wg sync.WaitGroup
	numGoroutines := 10
	counter := &AtomicCounter{} // pointer to the struct instance
	// value := 0

	for range numGoroutines {
		wg.Add(1)
		go func() {
			defer wg.Done()
			for range 1000 {
				counter.Increment() // increment the count 100000 times
				// value++
			}
		}()
	}
	wg.Wait()
	fmt.Println("Final counter value:", counter.GetValue())
	// fmt.Println("Final counter value:", value)

}
