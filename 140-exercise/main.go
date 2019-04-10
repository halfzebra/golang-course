package main

import (
  "fmt"
  "sort"
)

// Alternative: https://play.golang.org/p/BVRZTdlUZ_

type user struct {
  First   string
  Last    string
  Age     int
  Sayings []string
}

func main() {

  u0 := user{
    First: "Abe",
    Last:  "Whatever",
    Age:   27,
    Sayings: []string{
      "Shaken, not stirred",
      "Youth is no guarantee of innovation",
      "In his majesty's royal service",
    },
  }

  u1 := user{
    First: "James",
    Last:  "Bond",
    Age:   32,
    Sayings: []string{
      "Shaken, not stirred",
      "Youth is no guarantee of innovation",
      "In his majesty's royal service",
    },
  }

  u2 := user{
    First: "Miss",
    Last:  "Moneypenny",
    Age:   27,
    Sayings: []string{
      "James, it is soo good to see you",
      "Would you like me to take care of that for you, James?",
      "I would really prefer to be a secret agent myself.",
    },
  }

  u3 := user{
    First: "M",
    Last:  "Hmmmm",
    Age:   27,
    Sayings: []string{
      "Oh, James. You didn't.",
      "Dear God, what has James done now?",
      "Can someone please tell me where James Bond is?",
    },
  }

  users := []user{u0, u1, u2, u3}

  fmt.Println(users)

  sort.Slice(users, func (i, j int) bool {
    if (users[i].Age < users[j].Age) {
      return true
    } else if (users[i].First < users[j].First) {
      return true
    }
    return false
  })

  for _, u := range users {
    sort.Strings(u.Sayings)
  }


  for _, u := range users {
    fmt.Println(u.Age, u.First)
    fmt.Println(u.Sayings)
  }
}
