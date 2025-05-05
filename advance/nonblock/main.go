package main

//prevents deadlocks
import (
	"fmt"
	"time"
)

func main() {
	//  === NON BLOCKING RECEIVE OPERATION
	// ch := make(chan int)

	//  === NON BLOCKING RECEIVE OPERATION
	// select {
	// case msg := <-ch:
	// 	fmt.Println("Received:", msg)
	// default:
	// 	fmt.Println("No messages available")
	// }

	// === NON BLOCKING SEND OPERATIOn
	// select {
	// case ch <- 1:
	// 	fmt.Println("sent message")
	// default:
	// 	fmt.Println("CHannel is ready to receive")
	// 	// no receiver so doesnt do nothing
	// }

	// NON BLOCKING OPERATION IN REAL TIME SYSTEMS

	data := make(chan int)
	quit := make(chan bool)
	kool := make(chan string)

	go func() {
		for {
			select {
			case d := <-data:
				fmt.Println("Data received", d)
			case k := <-kool:
				fmt.Println("Koolshik received:", k)
			case <-quit:
				fmt.Println("Stopping......")
				return
			default:
				fmt.Println("Waiting for data...")
				time.Sleep(500 * time.Millisecond)
			}
		}
	}()

	// for
	// selecet{
	// 	case msg <-ch1:

	// 	case ctx <-ctx.Done()
	// 	fmt.error
	// 	}

	for i := range 5 {
		data <- i
		time.Sleep(time.Second)
	}
	names := []string{"hello", "Koushik", "sup", "dude"}

	for _, s := range names {
		kool <- s
		time.Sleep(time.Second)
	}

	quit <- true

}
