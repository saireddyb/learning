// Left Rotate an array by p place

package main

import "fmt"

func main() {
	arr := []int{1,2,3,4,5,6,7}
	p := 1
	rotaten(arr, p)
	fmt.Println(arr)
}

func rotaten(arr []int, p int) {
	var temp []int = make([]int, p)
	n := len(arr)
	for i:=0; i<p; i++{
		temp[i] = arr[i]
	}
	for i:=p; i<n; i++ {
		arr[i-p]= arr[i]
	}
	for i:= n-p; i < n; i++{
		arr[i] = temp[i-(n-p)]
	}
}