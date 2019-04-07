package main

import "fmt"

func counter() func () int {
  i := 0

  return func() int {
    i++
    return i
  }
}

func main() {
  c1 := counter()
  c2 := counter()

  fmt.Println("First counter")
  fmt.Println(c1())
  fmt.Println(c1())
  fmt.Println(c1())
  fmt.Println(c1())
  fmt.Println(c1())

  fmt.Println("Second counter")
  fmt.Println(c2())
  fmt.Println(c2())
}
