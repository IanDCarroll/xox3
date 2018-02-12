package gameRunner

import (
  "bitbucket.org/IanDCarroll/xox3/player"
  "bitbucket.org/IanDCarroll/xox3/ui/display"
  "fmt"
)

type GameRunner interface {
  PlayGame()
}

type gameRunner struct {
  player1 player.Player
  player2 player.Player
  display display.Display
}

func NewGameRunner(player1 player.Player, player2 player.Player, display display.Display) gameRunner {
  return gameRunner {player1, player2, display}
}

func (g gameRunner) PlayGame() {
  fmt.Println("The game is playing")
}
