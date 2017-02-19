package parser

import (
  "strings"
)

var (
  BASECOMMAND = "!drink"
  MIDCOMMANDS = []string{"name", "ingredients"}
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
  commands := strings.Split(input, " ")
  if len(commands) >= 2 {
    if commands[0] == BASECOMMAND {
      if commands[1] != "random" {
        if len(commands) >= 3 {
          if contains(MIDCOMMANDS, commands[1]) {
            response = "Finding drink based on "+commands[1]
          }
        }
      } else {
        // Get a random drink
        response = "Finding a random drink, please stand by..."
      }
    }
  }
  return response
}

// quick contains method

func contains(s []string, i string) bool {
  for _,a := range(s) {
    if a == i {
      return true
    }
  }
  return false
}
