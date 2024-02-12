// Remove duplicates from Sorted array

package main

import "fmt"

func main(){
	arr := []int{2,3,4,4,4,6,6,7,7,7,9}
	removeDuplicate(arr)
	fmt.Println(arr)
}

func removeDuplicate(arr []int) {
	n := len(arr)
	idx := 0 
	for i:=1;i<n; i++ {
		if arr[idx] != arr[i] {
			idx++
			arr[idx] = arr[i]
		}
	}
}