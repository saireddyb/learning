// Merge sort
package main

import "fmt"

func main() {
	arr := []int{1, 2, 3, 4, 5, 6, 7, 8}
	n := len(arr)
	mergeSort(arr, 0, n-1)
}

func mergeSort(arr []int, l int, r int) {
	if l >= r {
		return
	}
	mid := (l + r) / 2
	fmt.Println("l:mid", l, ":", mid, "=", arr[l:mid+1])
  fmt.Println("mid:r", mid+1, ":", r, "=", arr[mid+1:r+1])

	mergeSort(arr, l, mid)
	mergeSort(arr, mid+1, r)
	merge(arr, l, r, mid)
}

func merge(arr []int, l int, r int, mid int) {
  if arr[l] > arr[mid+1] {
    arr[l], arr[mid+1] = arr[mid+1], arr[l]
    
  }
}
