package main

import (
	"fmt"
	"time"
)

// go routites that maintains its own state acrros multiople interactions
// keeps track of the state informiation
// good for preserving data between data calls

type StatefullWorker struct {
	count int
	ch    chan int
}

func (w *StatefullWorker) Start() {
	go func() {
		for {
			select {
			case value := <-w.ch:
				w.count += value
				// count increases its previous value
				fmt.Println(w.count)
			}
		}
	}()
}

func (w *StatefullWorker) Send(value int) {
	w.ch <- value // send the value to the channel
}

func main() {

	stWorker := &StatefullWorker{
		ch: make(chan int),
	}
	stWorker.Start()

	for i := range 5 {
		stWorker.Send(i)
		time.Sleep(500 * time.Millisecond)
	}

}
