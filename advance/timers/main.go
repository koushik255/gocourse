package main

// allows to schedule events to intiate events at certain times,
// good for timeouts to limit how long a specific operation should wait
// also is good for delays
// best practices is to always stop timers to avoid resource leaks
// and use defer to ensure proper celan up

import (
	"fmt"
	"time"
)

// Multiple timers
func main() {
	timer1 := time.NewTimer(1 * time.Second)
	timer2 := time.NewTimer(2 * time.Second)

	select {
	case <-timer1.C:
		fmt.Println("1 second timers")
	case <-timer2.C:
		fmt.Println("2 second timer")
	}

}

// schdeuling delyaed operations
// func main() {
// 	timer := time.NewTimer(2 * time.Second) // non blocking timer starts

// 	go func() {
// 		<-timer.C // as soon as it receives it prints
// 		fmt.Println("Delayed operation executed")
// 	}()

// 	fmt.Println("waiting ....")
// 	time.Sleep(3 * time.Second) // need something to block the flow so we can use go rout
// 	fmt.Println("End of the program")
// }

//
// ======================TIMEOUT
// func longRunningOperation() {
// 	for i := range 20 {
// 		fmt.Println(i)
// 		time.Sleep(time.Second)
// 	}
// }

// func main() {
// 	timeout := time.After(3 * time.Second)
// 	done := make(chan bool)

// 	go func() {
// 		longRunningOperation()
// 		done <- true // sending signal that work is done and its a boolean channel
// 	}()

// 	select {
// 	case <-timeout:
// 		fmt.Println("operation timeout")
// 	case <-done:
// 		fmt.Println("Operation completed")
// 	}

// }
// basic timeout use
// func main() {
// 	fmt.Println("Starting application.......")
// 	timer := time.NewTimer(2 * time.Second) // sends the current time on its channel
// 	fmt.Println("Waiting for timer.c")
// 	stopped := timer.Stop() // stopped holds the boolean value
// 	if stopped {
// 		fmt.Println("timer stopped!")
// 	}
// 	timer.Reset(time.Second)
// 	fmt.Println("Timer Reset")
// 	<-timer.C
// 	// wait for the send channel this is blocking in nature
// 	// C is a channel which receives the timer when the timer expires
// 	fmt.Println("Timer expired")
// }
