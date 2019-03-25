package main

import "fmt"

func main() {
  favs := map[string][]string{
    "bond_james":      {`Shaken, not stirred`, `Martinis`, `Women`},
    `moneypenny_miss`: {`James Bond`, `Literature`, `Computer Science`},
    `no_dr`: {`Being evil`, `Ice cream`, `Sunsets`}}

  for name, list := range favs {
    fmt.Println(name, "likes:")
    for i, v := range list {
      fmt.Println(i, v)
    }
  }
}
