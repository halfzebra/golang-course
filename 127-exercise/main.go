package main

import "fmt"

func main() {
  x := 1
  y := &x
  fmt.Printf("%T %d \n", &x, &x)
  fmt.Printf("%T %d \n" , *y, *y)

  // *y dereferences the value
  *y = 2
  fmt.Println(x)
}
