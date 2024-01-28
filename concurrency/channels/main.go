package main

import (
	"fmt"
	"sync"
)

func main(){
	ch := make(chan string)
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
	fmt.Println("Send data to channel")
	wg.Done()
}

// Buy data to channel
func buy(ch chan string, wg *sync.WaitGroup){
	fmt.Println("Waiting for data")
	val := <- ch 
	fmt.Println("Received data from channel: ", val)
	wg.Done()
}