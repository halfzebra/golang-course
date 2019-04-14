// Package dog provides functions for understaning dogs.
package dog

import "errors"

// years converts human years to dog years
func Years(age int) (int, error) {
  if (age <= 0) {
    return 0, errors.New("Age mut be above 0")
  }

  return age / 7, nil
}

