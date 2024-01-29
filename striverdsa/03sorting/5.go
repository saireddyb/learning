// quick sort

package main

import "fmt"

func main(){
	arr := []int{4,3,6,7,8,1,2,9}
	quickSort(arr, 0, len(arr)-1)
	fmt.Println(arr)

	arri := []int{9, 7,8, 6, 5, 4, 3, 2, 1}
	quickSort(arri, 0, len(arri)-1)
	fmt.Println(arri)
	
}

func quickSort(arr []int, low int, high int){
	for low >= high {
		return
	}
	pIndex := sort(arr, low, high)
	quickSort(arr, low, pIndex-1)
	quickSort(arr, pIndex+1, high)
}

func sort(arr []int, low int, high int) int{
	pivot := arr[low]
	pivotIndex := low
	for low < high {
		for arr[low] <= pivot {
			low++
			if low > high {
				break
			}
		}
		for arr[high] > pivot{
			high--
			if low > high {
				break
			}
		}
		if low < high {
			arr[low], arr[high] = arr[high], arr[low]
		} else {
			arr[pivotIndex], arr[high] = arr[high], arr[pivotIndex]
		}
	}
	return high
}