package main

import (
  "fmt"
  "runtime"
  "sync"
)

func main() {
  runtime.GOMAXPROCS(2)
  fmt.Println(runtime.NumCPU())
  counter := 0
  var wg sync.WaitGroup
  wg.Add(2)

  go func() {
    for i := 0; i < 100; i++ {
      tmp := counter + 1
      //runtime.Gosched()
      counter = tmp
    }
    wg.Done()
  }()

  go func() {
    for i := 0; i < 100; i++ {
      tmp := counter + 1
      //runtime.Gosched()
      counter = tmp
    }
    wg.Done()
  }()


  wg.Wait()
  fmt.Println(counter)
}
