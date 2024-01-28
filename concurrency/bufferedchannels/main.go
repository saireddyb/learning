package main

import (
	"fmt"
	"sync"
)

func main() {
	var wg sync.WaitGroup
	wg.Add(2)
	ch := make(chan int, 3)
	go sell(ch, &wg)
	wg.Wait()
}

func sell(ch chan int, wg *sync.WaitGroup) {
	ch <- 10
	ch <- 11
	ch <- 12
	ch <- 13
	go buy(ch, wg)
	fmt.Println("Send all data to the channel")
	wg.Done()
}

func buy(ch chan int, wg *sync.WaitGroup) {
	fmt.Println("Waiting for data")
	fmt.Println("Receiving data from channel: ", <-ch)
	fmt.Println("Receiving data from channel: ", <-ch)
	wg.Done()
}

