package main

import (
  "fmt"
  "runtime"
  "sync"
)

func main() {
  mux := sync.Mutex{}
  runtime.GOMAXPROCS(2)
  fmt.Println(runtime.NumCPU())
  counter := 0
  var wg sync.WaitGroup
  wg.Add(2)

  go func() {
    mux.Lock()
    for i := 0; i < 100; i++ {
      tmp := counter + 1
      //runtime.Gosched()
      counter = tmp
    }
    mux.Unlock()
    wg.Done()
  }()

  go func() {
    mux.Lock()
    for i := 0; i < 100; i++ {
      tmp := counter + 1
      //runtime.Gosched()
      counter = tmp
    }
    mux.Unlock()
    wg.Done()
  }()


  wg.Wait()
  fmt.Println(counter)
}
