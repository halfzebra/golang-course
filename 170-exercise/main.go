package main

import "fmt"

func main() {
  ch := make(chan int)

  for i := 0; i < 10; i++ {
    go func() {
      for j := 0; j < 10; j++ {
        ch <- j
      }
    }()
  }

  for i := 0; i < 100; i ++ {
    fmt.Println(<-ch)
  }
}
