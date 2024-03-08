package main

import (
	"context"
	"fmt"
	"time"
)

func main() {
	// Create a new context being cancelled in 5 seconds.
	ctx, _ := context.WithTimeout(context.Background(), 5 * time.Second)
	// Start a new goroutine whose lifetime's attached to ctx.
	go Task(ctx)
	fmt.Println("main")
	time.Sleep(time.Second * 10)
}

func Task(ctx context.Context)  {
	time.Sleep(time.Second * 6)
	fmt.Println("hello")
	time.Sleep(time.Second * 4)
	fmt.Println("hello")
}