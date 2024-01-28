// Merge sort
package main

import "fmt"

func main() {
	arr := []int{2, 1, 3, 4, 5, 6, 7, 8}
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
  fmt.Println("Reached the merge function")
	temp := make([]int, r-l+1)
	left := l
	right := mid+1
	idx := 0


	for left <= mid && right <= r {
		if arr[left] < arr[right] {
			temp[idx] = arr[left]
			left++
		} else if arr[left] >= arr[right] {
			temp[idx] = arr[right]
			right++
		}
		idx++
	}
	for left <= mid {
		temp[idx] = arr[left]
		left++
		idx++
	}
	for right <= r {
		temp[idx] = arr[right]
		right++
		idx++
	}
	fmt.Println("Merged array", temp)
	for i := l; i <= r; i++ {
		arr[i] = temp[i-l]
	}
}
