// Multiple Recursion calls
// Fibonacci number using recussion
package main

import "fmt"

func main() {
	fmt.Println(f(6))
}

func f(n int) int{
  if n <= 1 {
    return n 
  }
  return f(n-1) + f(n-2)
}