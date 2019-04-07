package main

import "fmt"

func create(msg string) func (){
  return func () {
    fmt.Println(msg);
  }
}

func main() {
  fn := create("Hello!")

  fmt.Printf("%T", fn, );

  fn()
}
