package mymath

import (
  "math/rand"
  "testing"
)

func TestCenteredAvg(t *testing.T) {
  tests := []struct {
    data   []int
    answer float64
  }{
    {data: []int{1, 4, 6, 8, 100}, answer: 6,},
    {data: []int{0, 8, 10, 1000}, answer: 9,},
    {data: []int{9000, 4, 10, 8, 6, 12}, answer: 9,},
    {data: []int{123, 744, 140, 200}, answer: 170,},
  }

  for _, v := range tests {
    a := CenteredAvg(v.data)
    if (a != v.answer) {
      t.Errorf("Wrong answer %v for %v", a, v.data)
    }
  }
}

func BenchmarkCenteredAvg(b *testing.B) {
  for i := 0; i < 10; i++ {
    randInts := []int{}
    for j := 0; j < 10; j++ {
      randInts = append(randInts, rand.Int())
    }
    CenteredAvg(randInts)
  }
}
