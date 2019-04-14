// Package word provides functions for extracting word usage statistics from text strings.
package word

import "strings"

// UseCount returns the usage per word in a map.
func UseCount(s string) map[string]int {
  xs := strings.Fields(s)
  m := make(map[string]int)
  for _, v := range xs {
    m[v]++
  }
  return m
}

// Count returns an amout of unique words in the string.
func Count(s string) int {
  wordUseCount := UseCount(s)
  var totalUniqueWords int

  for _, v := range wordUseCount {
    totalUniqueWords += v
  }

  return totalUniqueWords
}
