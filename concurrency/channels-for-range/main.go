package main

import (
	"fmt"
	"sync"
)

func main(){
	ch := make(chan string, 4)
	var wg sync.WaitGroup
	wg.Add(2)
	go sell(ch, &wg)
	go buy(ch, &wg)
	// time.Sleep(3*time.Second)
	wg.Wait()
}

// Send data to channel
func sell(ch chan string, wg *sync.WaitGroup){
	ch <- "Furniture"
	ch <- "soft"
	ch <- "ice"
	close(ch)
	fmt.Println("Send data to channel")
	wg.Done()
}

// Buy data to channel
func buy(ch chan string, wg *sync.WaitGroup){
	fmt.Println("Waiting for data")
	// val := <- ch 
	// fmt.Println("Received data from channel: ", val)
	// val = <- ch 
	// fmt.Println("Received data from channel: ", val)
	for val := range ch {
		fmt.Println("Received data from channel: ", val)
	}

	wg.Done()
}