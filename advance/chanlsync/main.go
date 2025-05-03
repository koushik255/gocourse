package main

//ensures that datta is properly exchaneg between
// goroutines
// import "time"
import "fmt"

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

func main() {
	ch := make(chan int)

	go func() {
		ch <- 9
		fmt.Println("Sent value:")

		value := <-ch
		fmt.Println(value)
	}()
}
