package main

//allowing channels to store values
// so basically they do not need a imediate receiver
// only blocks when we try to feed more information to the buffer withoutoit flushing it
// the capacity is delcared at init
// channels wait for goroutines to finsih then they post the error

import (
	"fmt"
	"time"
)

// func main() {
// 	// ==========bloacking on receiver only if the buffer is empty
// 	ch := make(chan int, 2)

// 	go func() {
// 		time.Sleep(2 * time.Second)
// 		//sending value to channel
// 		ch <- 1
// 		ch <- 2
// 	}()
// 	fmt.Println("Value: ", <-ch)
// 	fmt.Println("Value: ", <-ch)
// 	// these receivers are force to wait for the sender from the GOfun
// 	fmt.Println("End of program")
// }

// ^^^^^^ BLOCKING ON SEN DONLY IF THE BUFFER IS FULL
func main() {
	//make(chan Type, capacity)
	ch := make(chan int, 2) // hold 2 int values
	ch <- 1
	ch <- 2
	go func() {
		time.Sleep(2 * time.Second)
		fmt.Println("Received:", <-ch)
	}()
	// this needs to be a go func because how else would the program
	// know if there are any space open, because remember it needs
	// to be INSTANT so it has to run concurrently

	fmt.Println("Blocking starts-- ")
	ch <- 3 // blocks because the buffer is full
	fmt.Println("Blocking ends--")
	fmt.Println("Received:", <-ch)
	fmt.Println("Received:", <-ch)

	fmt.Println("Buffered channels in Go!")

}
