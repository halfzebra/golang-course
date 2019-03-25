package main

import "fmt"

func main() {
  msg := "hello"
  if msg == "hello" {
    fmt.Println(msg, "there")
  } else if msg == "bye" {
    fmt.Println(msg, " bye!")
  } else {
    fmt.Println("hi")
  }
}
