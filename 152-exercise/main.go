package main

import (
  "fmt"
  "runtime"
  "sync"
  "sync/atomic"
)

func main() {
  runtime.GOMAXPROCS(2)
  fmt.Println(runtime.NumCPU())
  var counter int32 = 0
  var wg sync.WaitGroup
  wg.Add(2)

  go func() {
    for i := 0; i < 100; i++ {
      atomic.AddInt32(&counter, 1)
      fmt.Println(atomic.LoadInt32(&counter))
    }
    wg.Done()
  }()

  go func() {
    for i := 0; i < 100; i++ {
      atomic.AddInt32(&counter, 1)
      fmt.Println(atomic.LoadInt32(&counter))
    }
    wg.Done()
  }()

  wg.Wait()
  fmt.Println(counter)
}
