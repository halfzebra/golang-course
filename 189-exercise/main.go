package main

import (
  "fmt"
  "golang-course/189-exercise/dog"
)

func main() {
  dogsAgeInHumanYears, _ := dog.From(3)
  fmt.Printf("Neighbors dog is %v years old in human years.\n", dogsAgeInHumanYears)

  pensionAgeInDogYears, _ := dog.To(65)
  fmt.Printf("Dogs would go to pension when they are %v years old\n", pensionAgeInDogYears)
}
