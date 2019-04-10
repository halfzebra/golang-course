package main

import (
  "fmt"
  "runtime"
)

func main() {
  fmt.Println(runtime.GOOS, runtime.GOROOT(), runtime.GOARCH)
}
