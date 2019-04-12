package main

import (
  "fmt"
)

// Fixing the deadlock:
// https://play.golang.org/p/j-EA6003P0

func main() {
  channel := make(chan int)
  buffered := make(chan int, 1)

  buffered <- 42

  go func() {
    channel<-77
  }()

  fmt.Println(<-buffered)
  fmt.Println(<-channel)
}
