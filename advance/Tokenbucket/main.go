package main

import (
	"fmt"
	"time"
)

// rate limiting is a technique used to control the rate of requests to a server
// it is used to prevent too many requests from being made in a short period of time
// also helps manage resources
// / common rate limiting algorithms
// Token Bucket Algorihm
// Leaky Bucket Algorithm
// Fixed window Counter

// using empty stucts are used to signal or counts without needing to store data,
//  there is 0 memory allocation for a empty struct, they dont contain any data
// therefore you are just create a token with has no data
// and then you can send that token to the channel as a signal
// its a Go idiom

type RateLimiter struct {
	tokens     chan struct{} // receives structs
	refillTime time.Duration
}

func NewRateLimiter(rateLimit int, refillTime time.Duration) *RateLimiter {
	rl := &RateLimiter{
		tokens:     make(chan struct{}, rateLimit),
		refillTime: refillTime, // functoins receives this value
	}
	for range rateLimit {
		rl.tokens <- struct{}{} // fill the channel with the rate limit
		// just sending empty tokens
	}
	go rl.startRefill() // starting the refill goroutine
	return rl           // returning the instance of the rl struct
}

func (rl *RateLimiter) startRefill() {
	ticker := time.NewTicker(rl.refillTime) // creating a ticker that will tick every refillTime
	defer ticker.Stop()                     // stopping the ticker when the function ends
	// rl.refilltime is from the NewRateLimiter function and its a
	// param from that function but since this function is called after those
	// the params would be set in the calling of teh NewRateLimiter

	for {
		select {
		case <-ticker.C: // if a ticker is received
			select {
			case rl.tokens <- struct{}{}: // sends values into the channel
			// flls the channels back up to its rate limit
			default:
			}
		}
	}
}

func (rl *RateLimiter) allow() bool {
	select {
	case <-rl.tokens: // if the tokens channel receives a value
		// removes the a token from the token channel?
		return true
	default:
		return false
	}
} // so basically if the channel receives a value then it will return true and if it does not
// then it will drop to the default case and return false

func main() {

	rateLimiter := NewRateLimiter(5, time.Second)
	// allows 5 requests per second

	// accepting 5 but it allows 6
	// this is because the refill time is 1 secound
	// but we are allowing a new request every 0.2 sec 0.2*5 = 1 therefore
	// once we send 5 requests its time for the Ratelimiter to refill

	// So basically how it works is that
	// imagine there i a bowl of oreos
	// there is only 5 oreos in the bowl for prabir
	// I told prabir that he can only have 5 oreos
	// when he takes the 5 oreos the bowl is empty
	// so he cant take anymore oreos untill
	// the bowl is refilled
	// so the bowl / channel starts off as full with 5 oreos
	// but when you use the rate limiter you are taking
	// 1 oreo OUT at a time so the token is just what is
	// allowing you to make the request but it
	// does not play a role in anything else
	// kind of like a middle man, context
	for range 10 {
		if rateLimiter.allow() {
			fmt.Println("Request allowed")
		} else {
			fmt.Println("Request denied")
		}
		time.Sleep(200 * time.Millisecond)
		// it would range over to quicly and not wait for the GO routine to finish
		// so we sleep to slow it down
	}

}
