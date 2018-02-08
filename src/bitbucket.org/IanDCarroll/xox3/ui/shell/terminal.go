package shell

import (
  "fmt"
)

type terminal struct {}

func NewTerminal() terminal {
  return terminal {}
}

func (t terminal) Write(message interface{}) {
  fmt.Println(message)
}

func (t terminal) Read() interface{} {
  var input string
  fmt.Scanln(&input)
  return input
}
