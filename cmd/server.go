package main

import (
  "flag"
  "os"
  _ "fmt"
  cc "../core"
)

func main() {
  var printVer bool

  flag.BoolVar(&printVer, "version", false, "print charon version")

  flag.Parse()

  if printVer {
    cc.PrintVersion()
    os.Exit(0)
  }


}