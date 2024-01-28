package main

import "fmt"

func main() {
	anony := func() {
		fmt.Println("in anonymous function")
	}
	anony()
}