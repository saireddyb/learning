// Recurssion
// Sum of first N numbers

package main

import "fmt"

func main() {
	ans := sumn(3)
	fmt.Println(ans)
  sum2(3, 0)
  
}

func sumn(n int) int {
	if n <= 1 {
		return 1
	}
	return n + sumn(n-1)
}

func sum2(n int, sum int) {
  if n < 1 {
    fmt.Println(sum)
    return
  }
  sum2(n-1, sum+n)
}