package dog

import "errors"

func To(age int) (int, error) {
  if age < 0 {
    return 0, errors.New("Please provide age larger than zero.")
  }

  return age / 7, nil
}

func From(age int) (int, error) {
  if age < 0 {
    return 0, errors.New("Please provide age larger than zero.")
  }
  return age * 7, nil
}
