package main

import (
	"context"
	"fmt"
	"time"
)

func main() {
	ctx := context.Background()
	ctx, cancel := context.WithTimeout(ctx, time.Second*5)
	defer cancel()
	go func() {
		time.Sleep(time.Second * 2)
		fmt.Println("Hello")
	}()

	select {
	case <-ctx.Done():
		fmt.Println(ctx.Err())
		return
	}
	time.Sleep(time.Second * 10)
}
