package main

import (
	"fmt"
	"sync"
	"time"
)

func main() {
	wg := &sync.WaitGroup{}
	ch := make(chan int, 1) // allowing one message to sit in the channel without a receiver

	wg.Add(2)
	// example of bidirectional channel
	go func(ch <-chan int, wg *sync.WaitGroup) { // receive-only
		fmt.Println(<-ch)
		// ch <- 27 // can't send here
		wg.Done()
	}(ch, wg)

	go func(ch chan<- int, wg *sync.WaitGroup) {
		ch <- 42
		time.Sleep(5 * time.Millisecond)
		// fmt.Println(<-ch) // send-only. Can't receive message
		wg.Done()
	}(ch, wg)

	wg.Wait()
}