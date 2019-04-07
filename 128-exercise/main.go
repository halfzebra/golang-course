package main

import "fmt"

type user struct {
  name string
}

// *user uses the value at the pointer, allowing us to mutate the instance.
func (u *user) updateName(newName string) {
  u.name = newName
}

func main() {
  u1 := user{
    name: "John",
  }
  u1.updateName("Jake")
  fmt.Println(u1)
}
