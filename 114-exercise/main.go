package main

import "fmt"

func foo(s ...int) int {
  sum := 0
  for _, v := range s {
    sum += v
  }
  return sum
}

func bar(s []int) int {
  sum := 0
  for _, v := range s {
    sum += v
  }
  return sum
}

func main() {
  x := []int{1,2,3,4,5}
  fmt.Println(foo(x...))
  fmt.Println(bar(x))
}
