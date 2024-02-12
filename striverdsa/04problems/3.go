// Check if the array is already sorted

package main

import "fmt"

func main(){
	arr := []int{3,4,5,5,6,9,9,8}
	fmt.Println(checkSorted(arr))
}

func checkSorted(arr []int) bool {
	var n int = len(arr)
	if n < 2 {
		return true 
	}
	for i:=1; i < n ; i++ {
		if arr[i] < arr[i-1] {
			return false
		}
	}
	return true
}

