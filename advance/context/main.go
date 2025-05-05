package main

import (
	"context"
	"fmt"
	"time"
)

// contexts are good for requesting scoped values
// requesting scopes data across api boundaries
// context is a object which carrys request scoped data across go routines and apis
// contexts are variables that store values in key value pairs and carry signals
// context is importnat because it allows functions to stop their work if connections
// is cancelled, request can also be completed aswell
// in a web server you need to handle multiple requessts at once so each requests needs
// its own context, kinda like a phone call

// Plug: Represents the variable you want to pass (e.g., user ID).
// Socket: Represents the context.Context, which provides a way to pass the variable through the call chain.
// Plugging In: Using context.WithValue to associate a key with a value in the context.
// Using the Plug: Functions access the variable through the context, allowing them to use the data without needing to pass it explicitly.
// Unplugging: Once the function is done, it can stop using the variable, and the context manages its lifecycle.

func doWork(ctx context.Context) {
	for {
		select {
		case <-ctx.Done():
			fmt.Println("Work cancelled", ctx.Err())
			return
		default:
			fmt.Println("Working.....")
		}
		time.Sleep(500 * time.Millisecond)
	}
}

func main() {
	ctx := context.Background()
	ctx, cancel := context.WithTimeout(ctx, 2*time.Second)
	defer cancel()

	ctx = context.WithValue(ctx, "requestID", "kjsdfkjsafkljashfkjas")

	go doWork(ctx)

	// time.Sleep(1 * time.Second) -- this one for working stream
	time.Sleep(3 * time.Second)

	// this shows that even after the time limit it still holds that value
	requestID := ctx.Value("requestID")
	if requestID != nil {
		fmt.Println("request ID", requestID)
	} else {
		fmt.Println("No request id Found")
	}

}

// func checkEvenOdd(ctx context.Context, num int) string {
// 	select {
// 	case <-ctx.Done(): // check for anything then it returns that
// 		// ctx.Done is what would receive the cancel signal from the context
// 		return "Operation canceled"
// 	default:
// 		if num%2 == 0 {
// 			return fmt.Sprintf("%d is even ", num)
// 		} else {
// 			return fmt.Sprintf("%d is odd", num)
// 		}
// 	} // We are able to modify the behavior of the function based on the contexts state
// }

// func main() {
// 	ctx := context.TODO()
// 	result := checkEvenOdd(ctx, 5)
// 	fmt.Println("Result with context.TODO:", result)
// 	//The ctx variable serves as a control mechanism for the checkEvenOdd function. It allows the function to check if it should continue processing or if it should stop early due to a cancellation signal.
// 	ctx = context.Background()
// 	ctx, cancel := context.WithTimeout(ctx, 1*time.Second)
// 	// this context will cancel after 1 sec/ send cancel signal and it retains everything
// 	defer cancel() // cancel the context before the end of the main function

// 	result = checkEvenOdd(ctx, 10)
// 	fmt.Println("result from timeout context:", result)
// 	time.Sleep(2 * time.Second)
// 	result = checkEvenOdd(ctx, 15)
// 	fmt.Println("result after timeout context:", result)
// }

// DIFFERENCE BTWEEN THE CONTEXT.TODO and CONTEXT.BKG
// func main() {
// 	todoContext := context.TODO() // context is a instance of a structs/ object
// 	contextBkg := context.Background()

// 	ctx := context.WithValue(todoContext, "name", "john")
// 	fmt.Println(ctx)
// 	fmt.Println(ctx.Value("name")) // extracted value from key value pair

// 	ctx1 := context.WithValue(contextBkg, "city", "Paris")
// 	fmt.Println(ctx1)
// 	fmt.Println(ctx1.Value("city"))

// }
