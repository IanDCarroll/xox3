package main

import "fmt"

func main() {
    menu := BuildMenu()
    params := menu.GetGameParams()
    fmt.Println(params)
}
