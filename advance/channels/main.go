package main

import (
	"fmt"
	"time"
)

// channles provides communitaion between goroutine, helps in concurreny programs
// you cannot use a channel within ANY FUNCTION
// they allow one gorountine to send data to another gorountine in a thread safe manner

func main() {
	// variable := make(chan type)
	greeting := make(chan string)
	greetString := "Hello"

	go func() {
		greeting <- greetString
		greeting <- "world"
		for _, e := range "abcde" {
			greeting <- "Alphabet: " + string(e)
		}
		// for loop so we need a looping receiver

		// greeting is a open channel which receives data
		//the greetString could be an api or something along those lines/ stream

	}()
	//

	// greeting <- greetString // Blocking because it is continuously trying to receive

	// trys to send value to channel without having go routine redder
	// a go routine should be ready to recive this, channels in go are BLOCKINg
	// go routines are not blocking so we use channels to control flow within
	// our program

	//so basically the channel is communicating between the channel of the go func and
	// the main function, because receiver is inside the go routine of the main thread

	// receiveing values from a channel
	// receiver is part of the main thread/ a go rountine
	//so basically
	// receivering from a channel is blocking so if there is not value to be received
	// it waits for a value to be received
	// go func() {
	// 	receiver := <-greeting
	// 	fmt.Println(receiver)
	// 	receiver = <-greeting
	// 	fmt.Println(receiver)
	// }()

	receiver := <-greeting
	fmt.Println(receiver)
	receiver = <-greeting
	fmt.Println(receiver)

	// need a loop to receive values

	for range 5 {
		revr := <-greeting /// greeting goes to the receiver variable
		fmt.Println(revr)
	}

	// fmt.Println(receiver)

	// receiver = <-greeting
	// fmt.Println(receiver)
	// re excutute the statment
	time.Sleep(1 * time.Second)
	// pause so it has time to receive the go routine before the main thread stops

}
