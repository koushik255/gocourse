package main

import (
	"fmt"
	"time"
)

// go routine handle parralelel task such as input output opreations and more
// we dont have to manualy manage threads
// go routine scheduling usming m n scheduling model//
// M goutines are mappned onto N sytem threads
// multplexing is like switching , the scheudler of the goroutine switches goroutines onto availble os threads, this means it can run many goroutines on a limited number of thread by dyunamically sycheduler gorountines as nedder and this leads to an efficieunt use of resources within the syustem

func main() {
	// delcaring error variabes
	var err error
	// sayHello() // waits for 1 secound then prints out hello from go routine
	fmt.Println("Program start")

	go sayHello() // go extracts the functioj fomr the main thread and it takes it into the background, main function runs in the mainthread main function runs in the main thread, the go routine takes it out of this main room and brings it back to the main thread after its finished, like takes plate cleans it gives it back
	// we didnt get anyput because the funtion did not have anything else so when the go routine leaves the thread the main function is left with nothing
	fmt.Println("After sya hello function")
	// err = doWork() /// you cannot use goroutine to save values into function

	go func() {
		err = doWork() // so we make a anon func to do it for us since now the functions
		// basically
	}()

	go printNumber()
	go printLetter()

	// the time thing is why these dont happen all at teh same time not
	// you should not expect any specfic order for goroutines as its basically now in your controll at all

	time.Sleep(2 * time.Second)

	if err != nil {
		fmt.Println("error:", err)
		return
	} else {
		fmt.Println("Work completed succesfully!")
	}
	// this brings an error now because the dowork error only happens 1 sec into the fcuntion
}

func sayHello() {
	time.Sleep(1 * time.Second)
	fmt.Println("Hello from Goroutine!")
}

// since we did not let this goruntine complete this functions is still in the memeory

func printNumber() {
	for i := 0; i < 5; i++ {
		fmt.Println("Number:", i, time.Now())
		time.Sleep(100 * time.Millisecond)
	}
}

func printLetter() {
	for _, letter := range "abcde" {
		fmt.Println(string(letter), time.Now())
		time.Sleep(200 * time.Millisecond)
	}
}

func doWork() error {
	// something work
	time.Sleep(1 * time.Second)
	return fmt.Errorf("An error occured in doWork")
}

// concurrency versus parralelism
// concurrency means multipke task schelduled across differnt thread
// you can make use of concurreny with go runtimes
//parralalelism is multiple task getting done same time
