package main

import (
  "bitbucket.org/IanDCarroll/xox3/menu"
  "bitbucket.org/IanDCarroll/xox3/gameRunner"
)

func main() {
  menu := menu.BuildMenu()
  gameRunner := gameRunner.BuildAiGameRunner(menu.GetGameParams())
  gameRunner.PlayGame()
}
