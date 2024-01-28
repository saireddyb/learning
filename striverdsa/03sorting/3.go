// Insertion sort.

package main

import "fmt"

func main() {
	arr := []int{1, 3, 5, 8, 7, 2}
	insertionSort(arr)
	fmt.Println(arr)
}

func insertionSort(arr []int) {
	n := len(arr)
	for i := 0; i < n; i++ {
		fmt.Println(i)
		j := i
		for j > 0 && arr[j] < arr[j-1] {
			arr[j], arr[j-1] = arr[j-1], arr[j]
			j--
		}
	}
}
