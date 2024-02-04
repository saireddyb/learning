// Largest Element in an Array

package main

import "fmt"

func main(){
	arr := []int{5,43,1,4,5,6,4}
	ans := largestElement(arr)
	fmt.Println(ans)
}

func largestElement(arr []int) int {
	largest := arr[0]
	n := len(arr)
	for i:=0 ; i<n ; i++{
		if arr[i] > largest{
			largest = arr[i]
		}
	}
	return largest
}