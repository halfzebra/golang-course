package main

import (
  "fmt"
  "strings")

func main() {
   config  := struct {
    times int
    capitalize bool
  }{
  times: 4,
    capitalize: true,
  }

   fmt.Println(config)
   for i := 0; i < config.times; i++ {
     fmt.Println(strings.ToUpper("Hello"))
   }
}
