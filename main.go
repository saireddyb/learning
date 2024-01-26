package main

import "fmt"

func main() {
	f(1, 5)
}

func f(i int, j int) {
  if i > 5 {
    return
  }
	fmt.Println(i)
	f(i+1, j)
}
