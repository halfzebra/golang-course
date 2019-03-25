package main

import "fmt"

func main() {
  favSport := "baseball"
  switch favSport {
    case "baseball":
      fmt.Println("bat")
    case "football":
      fmt.Println("ball")
  }
}
