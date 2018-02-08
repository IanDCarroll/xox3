package main

import "fmt"

func main() {
  menu := BuildMenu()
  order, marker := menu.GetGameParams()
  fmt.Println("main got : ", order, " ", marker)
}
