package main

// good for resource management, task distrubution, and scalabillity
// tasks, workers, task queue,
// creaate a task channel, create worker goroutines, gistribute taks, Shutdown

import (
	"fmt"
	"time"
)

type ticketRequest struct {
	personID   int
	numTickets int
	cost       int
}

// simulate porcessing of ticket requests
func ticketProcessor(requests <-chan ticketRequest, results chan<- int) {
	for req := range requests {
		// when you delcar a channel with a specific type, Go knowns that nay
		// data sent throuhg that channel must be of that type
		fmt.Printf("Processing %d tickets of person %d and the cost is %d\n", req.numTickets, req.personID, req.cost)
		//Simulating processing time
		time.Sleep(time.Second)
		results <- req.personID
	}
}

func main() {
	numRequests := 5
	price := 5
	ticketRequests := make(chan ticketRequest, numRequests)
	ticketResults := make(chan int)

	// Starting ticket worker
	for range 3 {
		go ticketProcessor(ticketRequests, ticketResults)
	}

	// send tickets
	// creating job queue for these processors
	for i := range numRequests {
		ticketRequests <- ticketRequest{personID: (i + 1), numTickets: (i + 1) * 2, cost: (i + 1) * price}
	}
	close(ticketRequests)

	// ramge pver the results to recieve the values
	for range numRequests {
		fmt.Printf("Ticket for personID	%d process succufly !\n", <-ticketResults)
	}
}

// BASIC WORKER POOL PATTERN
// func worker(id int, tasks <-chan int, results chan<- int) {
// 	for task := range tasks {
// 		fmt.Printf("Worker %d processing tasks %d\n", id, task)
// 		// simulatie some worker
// 		time.Sleep(time.Second)
// 		results <- task * 2
// 	}
// }

// func main() {
// 	numWorker := 3
// 	numJobs := 10
// 	tasks := make(chan int, numJobs)
// 	results := make(chan int, numJobs)

// 	// Creating workers
// 	for i := range numWorker {
// 		go worker(i, tasks, results)
// 	}
// 	// now sending information to the workers for them to receive the information
// 	// send tasks to the tasks channel

// 	for i := range numJobs {
// 		tasks <- i
// 	}
// 	close(tasks)

// 	// collect the results
// 	for range numJobs {
// 		result := <-results
// 		fmt.Println(result)
// 	}

// }
