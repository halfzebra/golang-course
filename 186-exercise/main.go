package main

import (
  "fmt"
  "golang-course/186-exercise/dog"
)

func main() {
  age, _ := dog.Years(28)
  fmt.Printf("I'm %v years old in dog years", age)
}
