package main

import (
	"fmt"
	"time"
)

func main() {
	go start()
	time.Sleep(3 * time.Second)
}

func start(){
	go process()
	
	fmt.Println("In start")
	fmt.Println("In start")
	fmt.Println("In start")
	time.Sleep(1 * time.Second)

	fmt.Println("In start")
	fmt.Println("In start")
	fmt.Println("In start")

}

func process(){
	fmt.Println("In process")
}