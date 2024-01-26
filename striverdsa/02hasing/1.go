// Find the number of time a given number is repeated in a given array.

package main

import "fmt"

func main() {
	var arr = []int{1, 2, 3, 4, 5, 6, 7, 8, 9,9 ,9 }
	// n := 5
	count := make(map[int]int)
	for _, num := range arr {
		count[num]++
	}
	fmt.Println(count[10])
}
