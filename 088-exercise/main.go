package main

import "fmt"

func main() {
  quotes := [][]string{
    {"James", "Bond", "Shaken, not stirred"},
    {"Miss", "Moneypenny", "Helloooooo, James."}}
  for _,r := range quotes {
    for _, v := range r {
      fmt.Println(v)
    }
  }
}
