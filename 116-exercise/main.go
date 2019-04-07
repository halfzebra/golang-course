package main

import "fmt"

type person struct {
  first string
  last string
  age int
}

func (p person) speak() {
  fmt.Println("Hello, I'm", p.first, " and I'm ", p.age)
}

func main() {
  p := person{
    first: "Eduard",
    last: "Kyvenko",
    age: 28,
  }
  p.speak()
}
