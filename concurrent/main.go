package main

import (
	"fmt"
	"runtime"
	"sync"
	"time"
)

// concureency vs parallelism

// concureency is the ability of a syustem to handle multiple taks simultaneuosly, it involves managing
// multipe tasks that are in progresss at the same time but not nessarily at the same instand

// parallelism is the ability of a system to handle multiple tasks at the same time, it involves
// executing multiple tasks simultaneously on multiple cores or processors

// in concurency, tasks share the same physical CPU core, while in parallelism, each task is executed on a separate physical core

// parralelism is a subset of concurency

// func printNumbers() {
// 	for i := range 5 {
// 		fmt.Println(time.Now())
// 		fmt.Println(i)
// 		time.Sleep(500 * time.Millisecond)

// 	}
// }

// func printLetters() {
// 	for _, letter := range "ABCDE" {
// 		fmt.Println(string(letter))
// 		time.Sleep(500 * time.Millisecond)
// 	}
// }

func heavyTask(id int, wg *sync.WaitGroup) {
	defer wg.Done()
	fmt.Printf("Task %d is starting!\n", id)

	for range 100_000_000 {
	}
	blud := time.Now()
	fmt.Println(blud)
	fmt.Printf("Task %d is finsihed.\n", id)
}

func main() {

	// go printNumbers()
	// go printLetters()
	// // Go routines are work working in concurency

	// time.Sleep(3 * time.Second)

	numThreads := 12

	runtime.GOMAXPROCS(numThreads) // set the maximum number of CPU cores to use
	var wg sync.WaitGroup

	for i := range numThreads {
		wg.Add(1)
		go heavyTask(i, &wg)
	}

	wg.Wait()
}
