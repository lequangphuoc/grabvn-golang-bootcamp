package main

import (
	"context"
	"fmt"
	"time"
)

func do(ctx context.Context, c <-chan int) {
	defer fmt.Println("exit")
	for {
		select {
		case value := <-c:
			fmt.Println(value)
			time.Sleep(1 * time.Second)
		case <-ctx.Done():
			return
		}

	}
}

func main() {
	ctx, cancel := context.WithCancel(
		context.Background(),
	)
	c := make(chan int)
	go do(ctx, c)
	for i := 1; i < 10; i++ {
		c <- i
		if i == 5 {
			cancel()
		}
	}
	time.Sleep(2 * time.Second)
}
