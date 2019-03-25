package main

import "fmt"

func main() {
  var amount int = 1024;
  fmt.Printf("%d %b %X", amount, amount, amount)
  fmt.Println()
  nextAmount := amount << 1
  fmt.Printf("%d %b %X", nextAmount, nextAmount, nextAmount)
}
