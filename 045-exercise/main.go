package main

import "fmt"

// iota only increments within a single group of constants
const (
  a  = 2005 + iota
  b  = 2005 + iota
  c  = 2005 + iota
  d  = 2005 + iota
)

func main() {
  fmt.Print(a, b, c, d)
}
