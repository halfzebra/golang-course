package main

import "fmt"

type customErr struct {
  message string
}

func (err customErr) Error() string {
  return err.message
}

func foo(err error) {
  fmt.Println(err)
}

func main () {
  err := customErr{"Hello, this is a custom error!"}
  foo(err)
}
