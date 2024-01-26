// Find the number of time a given string is repeated in a given array.

package main

import "fmt"

func main() {
  var arr = []string{"a", "a", "d", "d", "e", "f", "g", "h", "i"}
  // n := 5
  count := make(map[string]int)
  for _, num := range arr {
    count[num]++
  }
  fmt.Println(count)
}
