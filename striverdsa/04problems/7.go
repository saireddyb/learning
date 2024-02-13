// Move all Zeros to the end of the array

package main

import "fmt"

func main() {
	arr := []int{1,0,0,0,2,3,0,0,4,5,0,6,7}
	// zeros1(arr) // With additional space solution
	zeros(arr)
	fmt.Println(arr)
}

func zeros(arr []int) {
	n := len(arr)
	idx := 0 

	for i:=0; i<n; i++{
		if arr[i] != 0 {
			arr[idx], arr[i] = arr[i], arr[idx]
			idx++
		}
	} 
}

func zeros1(arr []int){
	n := len(arr)
	var temp []int = make([]int, n)
	idx := 0
	for i:=0; i<n; i++{
		if arr[i] != 0 {
			temp[idx] = arr[i]
			idx++
		}
	}
	for i:=0; i<n; i++{
		if i < idx {
			arr[i] = temp[i]
		} else {
			arr[i] = 0
		}
	}
}