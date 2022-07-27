package main

import (
	"fmt"
)

func main() {
	unbuffer := make(chan string, 1)
	unbuffer <- "Có làm thì mới có ăn"

	fmt.Println(<-unbuffer)
}

func read(data <-chan string) {
	fmt.Println(<-data)
}

func write(data chan<- string) {
	data <- "Có làm thì mới có ăn"
}
