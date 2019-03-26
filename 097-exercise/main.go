package main

import "fmt"

type person struct {
  first string
  last string
  favorite []string
}

func main() {
  me := person{
    first: "Eduard",
    last: "Kyvenko",
    favorite: []string{ "fp", "games" },
  }

  max := person{
    first: "Maxim",
    last: "Maumenko",
    favorite: []string{ "magento", "driving" },
  }

 for _, v := range me.favorite {
   fmt.Println(v)
 }
 for _, v := range max.favorite {
   fmt.Println(v)
 }
}
