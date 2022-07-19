package main

import (
	"fmt"
	"math/rand"
	"time"
)

func main() {
	baseA := make(chan string)
	baseB := make(chan string)

	go callSugarA(baseA)
	go callSugarB(baseB)

	var selected string

	select {
	case sugarA := <-baseA:
		selected = sugarA
	case sugarB := <-baseB:
		selected = sugarB
	}

	fmt.Printf("Picked: %v\n", selected)
}

func callSugarA(ch chan string) {
	randTime := rand.Intn(5)
	fmt.Printf("Cost %v seconds to reach base A\n", randTime)
	time.Sleep(time.Second * time.Duration(randTime))
	ch <- "Sugar from base A"
}

func callSugarB(ch chan string) {
	randTime := rand.Intn(5)
	fmt.Printf("Cost %v second to reach base B\n", randTime)
	time.Sleep(time.Second * time.Duration(randTime))
	ch <- "Sugar from base B"
}
