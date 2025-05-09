package main

// each window has a counter that tracks the number of requests
// when the counter exceeds the limit the window is reset

import (
	"fmt"
	"sync"
	"time"
)

type RateLimiter struct {
	mu        sync.Mutex
	count     int
	limit     int
	window    time.Duration
	resetTime time.Time // tracks the amount of request and the count and the limit and the duration
	// protects data using Mutexes
}

// returns a instance of the RateLimiter struct
func NewRateLimiter(limit int, window time.Duration) *RateLimiter {
	return &RateLimiter{
		limit:  limit,
		window: window,
	}
}

// no we need to implement the allow method
func (rl *RateLimiter) Allow() bool {
	rl.mu.Lock()
	defer rl.mu.Unlock() // defer so everythinn inbetween in in the critical sescitons of the code

	now := time.Now()
	if now.After(rl.resetTime) {
		rl.resetTime = now.Add(rl.window) // chanign the value to now  adds the time durations
		// from the present time
		// adds that time to the time after every reset of the window
		// modifiyng the time to have the time form now + the duratoin
		rl.count = 0 // resets the count to 0
	}
	if rl.count < rl.limit {
		rl.count++
		return true
	}
	return false

}

func main() {
	rateLimiter := NewRateLimiter(5, 2*time.Second)
	// takes the limit of 5 and the duration of 1 second
	// creating 10 request
	for range 10 {
		if rateLimiter.Allow() { // reutrns true
			fmt.Println("REQUEST ALLOWED")
		} else {
			fmt.Println("REQUEST BLOCKED")
		}
		time.Sleep(200 * time.Millisecond) // sleeps for 100 milliseconds
	}

}
