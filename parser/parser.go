package parser

import (
  //"fmt"
)

var (
  BASECOMMAND = "!drink"
  MIDCOMMANDS = [3]string{"random", "name", "ingredients"}
)

type Parser struct {
  // Placeholder for a logger later
  id int
}

func NewParser() Parser {
  return Parser{id: 1}
}

func (p Parser) ParseCommand(input string) string {
  var response string
  if input == "Dang" {
    response = "Dong"
  }
  return response
}
