package main

import (
	"fmt"
	"math/rand"
	"sync"
	"time"
)

var cache = map[int]Book{}
var rnd = rand.New(rand.NewSource(time.Now().Local().UnixNano()))

func main() {

	wg := &sync.WaitGroup{} // Coordinate tasks and waiting among multiple go routines
	m := &sync.RWMutex{} // Mutexes Protect shared memory between go routines, to avoid race conditions
	for i := 0; i < 10; i++ {
		id := rnd.Intn(10) + 1
		wg.Add(2) // number of concurrent tasks to wait on
		
		go func(id int, wg *sync.WaitGroup, m *sync.RWMutex) {
			if b, ok := queryCache(id, m); ok {
				// fmt.Println(i)
				fmt.Println("from cache")
				fmt.Println(b)
			}
			wg.Done()
		}(id, wg, m)

		go func(id int, wg *sync.WaitGroup, m *sync.RWMutex) {
			if b, ok := queryDatabase(id, m); ok {
				// fmt.Println(i)
				fmt.Println("from database")
				fmt.Println(b)
			}
			wg.Done()
		}(id, wg, m)

		// fmt.Printf("Book not found with id: %v", id)
		time.Sleep(150 * time.Millisecond)
	}
	wg.Wait()
	// time.Sleep(2 * time.Second)
}

func queryCache(id int, m *sync.RWMutex) (Book, bool) {
	
	//m.Lock() // No efficient in this instance as we need to read often. R > W
	m.RLock() // Allows for multiple readers but only one writer
	b, ok := cache[id]
	m.RUnlock()
	return b, ok
}

func queryDatabase(id int, m *sync.RWMutex) (Book, bool) {
	time.Sleep(100 * time.Millisecond)

	for _, b := range books {
		if b.ID == id {
			m.Lock()
			cache[id] = b
			m.Unlock()
			return b, true
		}
	}
	return Book{}, false
}

