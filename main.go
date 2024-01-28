package main

import "fmt"

func main() {
	arr := []int{2,1}
	mergeSort(arr, 0, len(arr)-1)
	fmt.Println(arr)
}

func mergeSort(arr []int, low int, high int) {

	if low >= high {
		return
	}
	mid := (low+high)/2

	mergeSort(arr, low, mid)
	mergeSort(arr, mid+1, high)
	merge(arr, low, mid, high)
}

func merge(arr []int, low int, mid int, high int){
	left := low 
	right := mid+1
	temp := make([]int, high-low+1)
	idx := 0

	for left <= mid && right <= high {
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
	for right <= high {
		temp[idx] = arr[right]
		right++
		idx++
	}
	for i:=low; i<=high; i++ {
		arr[i] = temp[i-low]
	}
}