package main

import (
	"fmt"
	"sync"
)

func main() {
	var wg sync.WaitGroup
	go subF(&wg)
	fmt.Println("Process end!")
}

func subF(wg *sync.WaitGroup) {
	fmt.Println("Tuấn Đẹp Trai")
	fmt.Println("Tuấn Đẹp Trai")
	fmt.Println("Tuấn Đẹp Trai")
}
