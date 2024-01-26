// Functional recurssion
// 1: reverse a array using recurssion
// 2: Check palindrome using recurssion

package main

import "fmt"

func main() {
	fmt.Println(reverse([]int{1, 2, 3, 4}, 0, 3))
  fmt.Println(isPalindrome([]int{4, 3, 3, 4}, 0, 3))
}

func reverse(arr []int, start int, end int) []int {
	if start > end {
		return arr
	}
	arr[start], arr[end] = arr[end], arr[start]
	return reverse(arr, start+1, end-1)

}
func isPalindrome(arr []int, start int, end int) bool {
	if start > end {
		return true
	}
	if arr[start] != arr[end] {
		return false
	}
	return isPalindrome(arr, start+1, end-1)
} 
