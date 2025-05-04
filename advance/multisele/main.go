package main

//handling multiple channels at once
// allows go routintes to wait for multiple channel completations
//select is just like switch

import (
	"fmt"
	// "time"
)

func main() {

	ch := make(chan int)

	go func() {
		ch <- 1
		close(ch)
	}()

	for {
		select {
		case msg, ok := <-ch:
			if !ok {
				fmt.Println("Channel closed!")
				//clean up activies
				return
			}
			fmt.Println("received:", msg)
		}
	}
	// when a channel is closed the select received the value 0 of the type

}

// using time.after to implement timeouts
// func main() {
// 	ch := make(chan int)

// 	go func() {
// 		time.Sleep(2 * time.Second)
// 		ch <- 1
// 		close(ch) // close unless its an acitve flow of data
// 	}()

// 	select {
// 	case msg := <-ch:
// 		fmt.Println("Received ", msg)
// 	case <-time.After(1 * time.Second):
// 		fmt.Println("timeout.") // returns a channel of type time.time
// 	}
// }

// func main() {
// 	ch1 := make(chan int)
// 	ch2 := make(chan int)

// 	go func() {

// 		time.Sleep(time.Second)
// 		ch1 <- 1

// 	}()

// 	go func() {
// 		time.Sleep(time.Second)
// 		ch2 <- 2

// 	}()

// 	time.Sleep(2 * time.Second)

// 	for range 2 { // looping over it assures we get all possibilitys
// 		select {
// 		case msg := <-ch1:
// 			fmt.Println("received from ch1 :", msg)

// 		case msg := <-ch2:
// 			fmt.Println("received from ch2:", msg)
// 			// default:

// 			// 	fmt.Println("No channels ready..") // no waitnig for go routine to finsih
// 			//so default just goes of
// 		}
// 	}
// 	fmt.Println("End of program!")

// }
