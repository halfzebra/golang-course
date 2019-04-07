package main

import "fmt"

func main() {
  defer fmt.Println("Bye!")
  for i := 0; i < 5; i++ {
    fmt.Println(i)
  }
}
