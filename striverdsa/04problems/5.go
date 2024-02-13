// Left Rotate an array by one place

package main

import "fmt"

func main(){
	arr := []int{3,4,5,6,7,8}
	rotate(arr)
	fmt.Println(arr)
}

func rotate(arr []int) {
	first := arr[0]
	n := len(arr)
	for i:=1; i<n; i++{
		arr[i-1] = arr[i]
	}
	arr[n-1] = first
}