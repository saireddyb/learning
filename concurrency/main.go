package main

import (
	"fmt"
	"time"
)

func calculateSquare(i int){
    time.Sleep(1 * time.Second)
    fmt.Println(i*i)
}

func main() {
    fmt.Print("hello")
    start := time.Now()
    for i := 1; i <= 10000; i++ {
        go calculateSquare(i)
    }
    elapsed := time.Since(start)
    // time.Sleep(3 * time.Second)
    fmt.Println("Function took ", elapsed)
}