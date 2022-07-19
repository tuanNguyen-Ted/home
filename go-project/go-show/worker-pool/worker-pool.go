package main

import (
	"fmt"
	"math/rand"
	"sync"
	"time"
)

func main() {
	workers := []string{"worker1", "worker2", "worker3", "worker4", "worker5"}
	toDoJob := 100

	wg := &sync.WaitGroup{}

	roomChannel := make(chan string)

	go defineJob(toDoJob, roomChannel)

	for _, worker := range workers {
		wg.Add(1)
		go doTheJob(wg, worker, roomChannel)
	}

	wg.Wait()
}

func defineJob(toDoJob int, roomChannel chan string) {
	for i := 1; i <= toDoJob; i++ {
		fmt.Println("Define the Job: ", i)
		roomChannel <- fmt.Sprintf("Job%d", i)
	}
	close(roomChannel)
}

func doTheJob(wg *sync.WaitGroup, worker string, roomChannel chan string) {
	defer wg.Done()
	for {
		job, ok := <-roomChannel
		if !ok {
			fmt.Println("No more job need to be done: ", worker)
			return
		}

		fmt.Printf("Worker %s is working on %s\n", worker, job)
		wait(5)
	}
}

func wait(maxNum int) {
	time.Sleep(time.Second * time.Duration(rand.Intn(maxNum)))
}
