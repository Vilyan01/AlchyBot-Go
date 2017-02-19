package boterrors

import (
  "errors"
)

var (
  NILDISCORDTOKEN error
)

func init() {
  NILDISCORDTOKEN = errors.New("The Discord token was nil")
}
