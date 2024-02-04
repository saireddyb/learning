// Find second largest element in the array

package main

import (
	"fmt"
	"math"
)

func main(){
	arr := []int{5,4,88,1,2,9}
	ans := secondlargestElement(arr)
	fmt.Println(ans)
	ans = secondsmallestElement(arr)
	fmt.Println(ans)
}
func secondsmallestElement(arr []int) int {
	smallest := arr[0]
	secondsmallest := int(math.Inf(1))
	n := len(arr)
	for i:= 0; i<n; i++{
		if arr[i] < smallest {
			secondsmallest = smallest
			smallest = arr[i]
		} else if arr[i]
	}
}

func secondlargestElement(arr []int) int{
	largest := arr[0]
	secondlargest := -1
	n := len(arr)
	for i:=0;i<n;i++{
		if arr[i] > largest {
			secondlargest = largest
			largest = arr[i]
		}
		if arr[i] != largest && arr[i] > secondlargest {
			secondlargest = arr[i]
		} 
	}
	return secondlargest
}