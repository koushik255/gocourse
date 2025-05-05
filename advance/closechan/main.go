package main

import "fmt"

// why close channels? for signal completion and to preven resource leaks
// every channel needs to be closed doesnt matter is buffered or unbufferd

func producer(ch chan<- int) {
	for i := range 5 {
		ch <- i
	}
	close(ch)
}

func filter(in <-chan int, out chan<- int) {
	for val := range in {
		if val%2 == 0 {
			out <- val
		}
	}
	close(out)
}

func main() {
	ch1 := make(chan int)
	ch2 := make(chan int)

	go producer(ch1)
	go filter(ch1, ch2)

	for val := range ch2 {
		fmt.Println(val)
	}

}

// RANGE OVER CLOSED CHANNEL
// func main() {
// 	ch := make(chan int)
// 	go func() {
// 		for i := range 5 {
// 			ch <- i
// 		}
// 		close(ch)
// 	}()

// 	for val := range ch {
// 		fmt.Println(val)
// 	}
// }

// // RECEVING FROM A CLOSED CHANNEL
// func main() {

// 	ch := make(chan int)
// 	close(ch)

// 	val, ok := <-ch
// 	if !ok {
// 		fmt.Println("Channel is closed!")
// 		return
// 	}
// 	fmt.Println(val)

// 	// receives 0 since its a

// }

// simple closing channel exmaple
// func main() {

// 	ch := make(chan int)
// 	go func() {
// 		for i := range 5 {
// 			ch <- i
// 		}
// 		close(ch)
// 	}()

// 	for val := range ch {
// 		fmt.Println(val)
// 	}

// }
