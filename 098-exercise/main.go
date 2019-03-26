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

  friends := map[string]person{
    me.last: me,
    max.last: max,
  }

  for _, d := range friends {
    fmt.Println(d.first, d.last, "likes:");
    for _, v := range d.favorite {
      fmt.Println(v)
    }
  }
}
