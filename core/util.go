package core

import (
  "errors"
  "fmt"
  "os"
)

const version = "0.0.1"

func PrintVersion() {
  fmt.Println("charon version: ", version)
}

func IsFileExists(path string) (bool, error) {
  stat, err := os.Stat(path)
  if err == nil {
    if stat.Mode() & os.ModeType == 0 {
      return true, nil
    }
    return false, errors.New(path + " no available...")
  }

  if os.IsNotExist(err) {
    return false, nil
  }

  return false, nil
}
