package main

import (
  "fmt"
  "log"
  "errors"
)

type sqrtError struct {
  lat  string
  long string
  err  error
}

func (se sqrtError) Error() string {
  return fmt.Sprintf("math error: %v %v %v", se.lat, se.long, se.err)
}

func main() {
  _, err := sqrt(-10.23)
  if err != nil {
    log.Println(err)
  }
}

func sqrt(f float64) (float64, error) {
  if f < 0 {
    // write your error code here
    return 0, sqrtError{
      lat:   "50.2289 N",
      long: "99.4656 W",
      err: errors.New("parameter is less than zero"),
    }
  }
  return 42, nil
}
