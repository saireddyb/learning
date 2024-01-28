package main

import (
	"fmt"
	"sync"
	"time"
)

func calculateSquare(i int, wg *sync.WaitGroup){
	time.Sleep(1*time.Second)
	fmt.Println(i*i)
	wg.Done()
}
func main(){
	start := time.Now()
	var wg sync.WaitGroup
	wg.Add(10)
	for i := 0; i < 10; i++ {
		go calculateSquare(i, &wg)
	}
	elapsed := time.Since(start)
	wg.Wait()
	fmt.Println("Function took", elapsed)
}