package main

//ensures that datta is properly exchaneg between
// goroutines
// import "time"
import (
	"fmt"
	"time"
)

// func main() {
// 	done := make(chan bool)

// 	go func() {
// 		fmt.Println("Working...")
// 		time.Sleep(2 * time.Second)
// 		done <- true
// 	}()

// 	<-done
// 	// waits until it receives a value until it moves onto
// 	// the next line of the code
// 	fmt.Println("Finished ")

// }

// func main() {
// 	ch := make(chan int)

// 	go func() {
// 		ch <- 9 // Blocking until the value is received
// 		fmt.Println("Sent value:")
// 	}()
// 	value := <-ch //blocking until a value is sent
// 	fmt.Println(value)
// }

// SYNCHRONIZING MULTIPEL GO ROUTINES
// and ensyring that all goroutines are finished
// func main() {
// 	numGoroutines := 3
// 	done := make(chan int, 3)

// 	for i := range numGoroutines {
// 		// each time the loop runs the go keyword
// 		// launches a new independnt routine
// 		time.Sleep(2 * time.Second)
// 		go func(id int) {
// 			fmt.Printf("Goroutine %d working\n", id)
// 			done <- id // signal to confirm its finished
// 		}(i)
// 	}

// 	for range 2 {
// 		<-done // Waiting for each each gorutine to finsih
// 	}

// 	fmt.Println("ALl goroutines are finished")
// }

// SYNCHRONIZING DATA EXCHANGE
func main() {
	//using channels to pass data around
	data := make(chan string)

	go func() {
		for i := range 5 {
			data <- "Hello :" + string('0'+i)
			time.Sleep(100 * time.Millisecond)
		} // simlutating that we are sneidng data 5 times
		close(data) // close the go routine once there is no more
		//data to send
	}()
	// using channels to send data in a sequencial way

	// since its a continuous flow of data we must use a for loop
	// to receive all the data

	// this channel is still open so that our for loop is still
	// receiving values, it creates a receiver for something which
	// has noo sneder so then it calls deadlock

	for value := range data {
		fmt.Println("Recvied date:", value, time.Now())
	} // alternate way to receive data from a flow of data from
	// channel

}
