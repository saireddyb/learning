// Find the highest and lowest frequence element

package main 

import "fmt"

func main() {
  var arr = []int{1,2,1,2,3,3,4,4,5,5,6,6,6,6,6,6,6,7,8,8,9,9}
  var max = arr[0]
  var maxCount = 1
  var min = arr[0]
  count := make(map[int]int)
  for _, num := range arr {
    count[num]++
    if count[num] > maxCount {
      max = num
      maxCount = count[num]
    }
  }
  var minCount = count[arr[0]]
  
  for key, value := range count {
    if value < minCount {
      minCount = value
      min = key
    }
  }
  fmt.Println(max)
  fmt.Println(min)
}