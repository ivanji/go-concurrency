package main

import (
	"fmt"
	"sync"
)

// understanding channels

// Unbeffered Channels are blocking constructs.
// It'll block this goRoutine until a message is available
// Receivers must be available in order for it to work
func main() {
	wg := &sync.WaitGroup{}
	ch := make(chan int)

	wg.Add(2)
	go func(ch chan int, wg *sync.WaitGroup) {
		fmt.Println(<-ch)
		close(ch)
		fmt.Println(<-ch) // Attempting to get a message from a closed channel returns 0
		wg.Done()
	}(ch, wg)

	go func(ch chan int, wg *sync.WaitGroup) {
		ch <- 42
		// ch <- 27 // causes deadlock
		wg.Done()
	}(ch, wg)

	wg.Wait()
}