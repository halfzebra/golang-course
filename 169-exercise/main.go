package main

import "fmt"

func main() {
  numbers := make(chan int)

  go func (ch chan<- int) {
    for i := 0; i < 100; i++ {
      ch<-i
    }
    close(ch)
  }(numbers)

  for v := range numbers {
    fmt.Println(v)
  }

  fmt.Println("Done")
}
