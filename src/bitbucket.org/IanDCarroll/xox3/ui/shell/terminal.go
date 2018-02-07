package shell

import "fmt"

type terminal struct {}

func NewTerminal() terminal {
  return terminal {}
}

func (t terminal) Write(message string) {
  fmt.Println(message)
}

func (t terminal) Read() int {
  return 42
}
