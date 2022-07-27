package main

import (
	"fmt"
	"sync"
)

func main() {
	var wg sync.WaitGroup
	wg.Add(1)
	go subF(&wg)
	wg.Wait()
	fmt.Println("Process end!")
}

func subF(wg *sync.WaitGroup) {
	fmt.Println("Tuấn Đẹp Trai")
	fmt.Println("Tuấn Đẹp Trai")
	fmt.Println("Tuấn Đẹp Trai")
	defer wg.Done()
}
