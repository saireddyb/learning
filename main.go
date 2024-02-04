// quick sort

package main

import "fmt"

func main(){
	arr := []int{4,6,2,5,7,9,1,3}
	n := len(arr)
	quicksort(arr, 0, n-1)
	fmt.Println(arr)
}

func quicksort(arr []int, low int, high int){
	if low >= high {
		return
	}
	i := low 
	j := high 
	pivot := arr[i]
	pivotIndex := i 

	for i < j {
		for arr[i] <= pivot && i < j {
			i++
		}
		for arr[j] >= pivot && i < j {
			j--
		}
		if i < j {
			arr[i], arr[j] = arr[j], arr[i]
		}
	}
	j--
	arr[j], arr[pivotIndex] = arr[pivotIndex], arr[j]
	quicksort(arr, low, pivotIndex-1)
	quicksort(arr, pivotIndex+1, high)
}
