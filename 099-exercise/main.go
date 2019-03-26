package main

import "fmt"

type vehile struct {
  color string
  doors int
}

type sedan struct {
  vehile
  luxury bool
}

type truck struct {
  vehile
  fourWheen bool
}

func main() {
  momsCar := sedan{
    vehile: vehile{
      color: "red",
      doors: 2,
    },
    luxury: false,
  }

  dadsCar := truck{
    vehile: vehile{
      color: "white",
      doors: 5,
    },
    fourWheen: true,
  }

  fmt.Println(momsCar)
  fmt.Println(momsCar.color, momsCar.doors, momsCar.luxury)
  fmt.Println(dadsCar)
}
