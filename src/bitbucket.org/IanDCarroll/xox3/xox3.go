package main

import (
  "fmt"
  "bitbucket.org/IanDCarroll/xox3/menu"
)

func main() {
  menu := menu.BuildMenu()
  order, marker := menu.GetGameParams()
  fmt.Println("main got : ", order, " ", marker)
}
