package word

import "testing"

func TestCount(t *testing.T) {
  total := Count("hello world")

  if total != 2 {
    t.Error("Should return correct amount")
  }
}

func BenchmarkCount(b *testing.B) {
  Count("hello hello and bye bye")
}

func BenchmarkUseCount(b *testing.B) {
  UseCount("hello hello and bye bye")
}
