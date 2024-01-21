package main

import "fmt"

var i = 5

func main() {
	fmt.Println("hello")
	recur()
}

func recur() {
	if i == 0 {
		return
	}
	fmt.Println(i)
	i--
  recur()

}
