package main

import "fmt"

func main() {
    switch {
    case false:
      fmt.Println("false")

    case true:
      fmt.Println("true")
      fallthrough
    default:
      fmt.Println("hello")
    }
}
