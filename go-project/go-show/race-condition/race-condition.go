package main

import (
	"fmt"
	"sync"
)

func main() {
	var count int32
	var wg sync.WaitGroup
	var mutex sync.Mutex
	wg.Add(3)
	go addI(&count, &wg, &mutex)
	go addI(&count, &wg, &mutex)
	go addI(&count, &wg, &mutex)
	wg.Wait()
	fmt.Println(count)
}

func addI(count *int32, wg *sync.WaitGroup, mutex *sync.Mutex) {
	for i := 0; i < 2000; i++ {
		mutex.Lock()
		*count++
		mutex.Unlock()
		// atomic.AddInt32(count, 1)
	}
	wg.Done()
}
