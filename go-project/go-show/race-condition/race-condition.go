package main

import (
	"fmt"
	"sync"
)

func main() {
	var count int32
	var wg sync.WaitGroup
	wg.Add(3)
	go addI(&count, &wg)
	go addI(&count, &wg)
	go addI(&count, &wg)
	wg.Wait()
	fmt.Println(count)
}

func addI(count *int32, wg *sync.WaitGroup) {
	for i := 0; i < 2000; i++ {
		*count++
	}
	wg.Done()
}
