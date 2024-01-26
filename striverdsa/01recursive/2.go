package main

import "fmt"

func main() {
	fmt.Println("f1ton")
	f1ton(5)
	fmt.Println("fnto1")
	fnto1(5)
}

func f1ton(i int) {
	if i == 0 {
		return
	}
	f1ton(i - 1)
	fmt.Println(i)
}

func fnto1(i int) {
	if i == 0 {
		return
	}
	fmt.Println(i)
	fnto1(i - 1)
}
