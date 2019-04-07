package main

import (
  "fmt"
  "strings"
)

func mapper(fn func(s string) string, l []string) []string {
  res := []string{}
  for _, v := range l {
    res = append(res, fn(v))
  }
  return res
}

func main() {
  fmt.Println(mapper( strings.ToUpper, []string{"hello", "horld"}))
}
