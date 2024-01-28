// Bubble sort

package main

import "fmt"

func main() {
	arr := []int{1, 3, 4, 8, 7, 6}
	bubbleSort(arr)
  fmt.Println(arr)
}

func bubbleSort(arr []int) {
	n := len(arr)
	for i := 0; i < n; i++ {
		for j := 0; j < n-i-1; j++ {
      if arr[j] > arr[j+1] {
        arr[j], arr[j+1] = arr[j+1], arr[j]
      }
		}
	}
}
