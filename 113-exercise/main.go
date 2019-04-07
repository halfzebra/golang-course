package main

import "fmt"

func foo() int {
  return 1
}

func bar() (int, string) {
  return 2, "hello"
}

func main() {
  x := foo()
  fmt.Println(x)
  y, z := bar()
  fmt.Println(y, z)
}
