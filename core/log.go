package core

import (
  _ "log"
  _ "os"
)

type LogLevel int

const (
  Debug LogLevel = iota
  Info
  Warn
  Error
  Fatal
)