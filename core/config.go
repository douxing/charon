package core

import (
  _ "encoding/json"
  _ "os"
  _ "reflect"
)

type Config struct {
  Server     interface{} `json:"server"`
  ServerPort int         `json:"server_port"`
  Password   string      `json:"password"`
  Encryption string      `json:"method"`

  // used by server only
  Timeout    int         `json:"timeout"`


  // used by client only
  LocalPort  int         `json:"local_port"`


}
