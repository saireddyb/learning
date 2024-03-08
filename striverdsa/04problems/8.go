// Linear Search

package main

import "fmt"

func main() {
	arr := []int{4,5,6,7,8,9}
	ans := search(arr,11)
	fmt.Println(ans)
}

func search(arr []int, num int) int {
	n := len(arr)
	for i:=0; i<n; i++{
		if arr[i] == num {
			return i
		}
	}
	return -1
}