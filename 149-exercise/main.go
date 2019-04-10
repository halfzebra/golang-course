package main

import "fmt"

type person struct {
  name string
  age int
}

// https://golang.org/doc/faq#different_method_sets
func (p *person) speak() {
  fmt.Println(p)
}

type human interface {
  speak()
}

func saySomething(h human) {
    h.speak()
}

func main() {
  p := person{
    name: "John",
    age: 30,
  }
  p.speak()

  saySomething(&p)
  // this will crash: saySomething(p)
}
