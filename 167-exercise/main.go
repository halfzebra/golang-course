package main

import (
  "fmt"
  "time"
)

func main() {
  stop := make(chan bool)
  numbers := gen(stop)

  receive(numbers, stop)

  fmt.Println("about to exit")
}

func receive(numbers  <-chan int, stop <-chan bool) {
  for {
    select {
    case v := <-numbers:
      fmt.Println(v)

    case v := <-stop:
      fmt.Println(v)
      return
    }
  }
}

func gen(stop chan<- bool) <-chan int {
  c := make(chan int)

  go func() {
    for i := 0; i < 100; i++ {
      c <- i
    }
    time.Sleep(time.Second)
    stop<-true
    close(c)
  }()

  return c
}
