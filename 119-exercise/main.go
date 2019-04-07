package main

import "fmt"

func main() {
  fn := func() string {
    return "Hello!"
  }

  fmt.Printf("%T %v", fn, fn());
}
