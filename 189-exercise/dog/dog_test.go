package dog

import (
  "fmt"
  "testing"
)

func TestFrom(t *testing.T) {
  age, _ := From(4)

  if age != 28 {
    t.Errorf("Conversion incorrect. Got: %v", age)
  }
}

func TestTo(t *testing.T) {
  age, _ := To(40)

  if age != 5 {
    t.Errorf("Conversion incorrect. Got: %v", age)
  }
}

func BenchmarkFrom(b *testing.B) {
  for i := 1; i < 100; i++ {
    From(i)
  }
}

func BenchmarkTo(b *testing.B) {
  for i := 1; i < 100; i++ {
    To(i)
  }
}


func ExampleFrom() {
  age, err := From(-10)
  fmt.Println(age, err)
}
