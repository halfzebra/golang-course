package main

import (
  "fmt"
  "math"
)

type square struct {
  L float32
  W float32
}

type circle struct {
  r float32
}

func (s square) area() float32 {
  return s.L * s.W
}

func (c circle) area() float32 {
  return math.Pi * c.r * 2
}

type shape interface {
  area() float32
}

func  info(s shape) {
  fmt.Printf("%T %v\n", s, s.area())
}

func main() {
  s := square{
    L: 23.1,
    W: 17.6,
  }
  fmt.Println(s.area())

  c :=  circle{
    r: 135.1,
  }
  fmt.Println(c.area())

  info(s)
  info(c)
}
