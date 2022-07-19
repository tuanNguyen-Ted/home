package main

import (
	"fmt"
	"time"
)

func sub1() {
	time.Sleep(100 * time.Millisecond)
	fmt.Println("Trứng rán cần mỡ")
	time.Sleep(100 * time.Millisecond)
	fmt.Println("Yêu không cần cớ")
}

func sub2() {
	time.Sleep(150 * time.Millisecond)
	fmt.Println("Bắp cần bơ")
	time.Sleep(150 * time.Millisecond)
	fmt.Println("Cần cậu cơ")
}

func main() {
	go sub1()
	go sub2()
	time.Sleep(2 * time.Second)
	fmt.Println("Good bye")
}
