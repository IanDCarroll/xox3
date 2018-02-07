package main

import (
  "fmt"
  "bitbucket.org/IanDCarroll/xox3/ui/shell"
)

func main() {
  shell := shell.NewTerminal()
  menu := BuildMenu(shell)
  order, marker := menu.GetGameParams()
  fmt.Println("main got : ", order, " ", marker)
}
