package main

import "fmt"

func main() {
  g := 1 == 2
  h := "hello" <= "hola"
  i := 100 >= 100
  j := g != i
  k := 3 < 2
  l := 0 > 1
  fmt.Println(g, h, i, j, k, l)
}
