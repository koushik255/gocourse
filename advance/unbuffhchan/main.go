package main

// bufferd channel is a channel with storage, a buffer is storage kind of like flush
// bufferend channels allow asyncronous communiation, allows working without block
// unbufferend channels need an immediate receiver
// send data into channel (ch <- value) then
// receiver data from channel (result:= <-ch)

import (
	"fmt"
	"time"
)

func main() {
	ch := make(chan int)
	go func() {
		time.Sleep(3 * time.Second)

		fmt.Println(<-ch)
		fmt.Println("3 second Goroutine finished")
	}()
	ch <- 1
	// receiver := <-ch
	// fmt.Println(receiver)

	// channels are a way to talk between go routines, when using un-bufferd channels
	// the channel needs an IMMEDIATE receiver, so in go routine since its on different,
	// threads and both running at same time it can fine that reciever but in this
	// example the receiver is on the next line and the channel cannot wait that long

	///////
	// ch := make(chan int)
	// // needs a reciever
	// go func() {
	// 	ch <- 1
	// 	time.Sleep(2 * time.Second)
	// 	fmt.Println("2 second Goroutine finished")
	// }()

	// go func() {
	// 	// ch <- 1
	// 	time.Sleep(3 * time.Second)
	// 	fmt.Println("3 second Goroutine finished")
	// }()

	// receiver := <-ch
	// fmt.Println(receiver)
	fmt.Println("End of program")

}
