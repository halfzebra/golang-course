package main

import (
  "fmt"
  "sync"
)

func main() {
  var wg sync.WaitGroup
  wg.Add(2)
  go func() {
    fmt.Println("Goroutine 1")
    wg.Done()
  }()

  go func() {
    defer wg.Done()
    fmt.Println("Goroutine 2")
  }()
  fmt.Println("Main Goroutine")
  wg.Wait()
}
