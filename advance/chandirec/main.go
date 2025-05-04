package main

// channel diretoins are used in functions and goroutines

import (
	"fmt"
)

func main() {
	ch := make(chan int)

	// accepts send only channels
	// after go routine preforms action
	// }()

	// //receiving
	// for value := range ch {
	// 	fmt.Println("value:", value)
	// }

	// receiuve only channll

	// close it to prevent memory leaks
	// this ch is a bidiretion channl
	// uni direction would be make(<- chan int)

	producer(ch)
	consumer(ch)
}

// send only channel
func producer(ch chan<- int) {
	go func() {
		for i := range 5 {
			ch <- i
		}
		close(ch)
	}()
}

// receive only channel
func consumer(ch <-chan int) {
	for value := range ch {
		fmt.Println("value:", value)
	}
}
