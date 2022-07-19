package main

import (
	"sync"
	"sync/atomic"
)

var mutex = &sync.Mutex{}

func main() {
	for {
		var i atomic.Value
		var wg sync.WaitGroup
		i.Store(int32(0))
		wg.Add(3)
		go Process(&i, &wg)
		go Process(&i, &wg)
		go Process(&i, &wg)
		wg.Wait()
		if i.Load().(int32) != 6000 {
			panic("Race conditon")
		}
	}

}

func Process(variable *atomic.Value, wg *sync.WaitGroup) {
	defer wg.Done()
	for i := 0; i < 2000; i++ {
		// mutex.Lock()
		// *variable++
		variable.Store(variable.Load().(int32) + 1)
		// mutex.Unlock()
		// atomic.AddInt32(variable, 1)
	}
}
