package gameRunner

import (
  "bitbucket.org/IanDCarroll/xox3/player"
  "bitbucket.org/IanDCarroll/xox3/ui/display"
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
  gameover := false
  winner := 0
  for move := 1; !gameover; move++ {
      gameover, winner = g.togglePlayers(move)
  }
  g.display.WriteToShell(winner)
}

func (g gameRunner) togglePlayers(move int) (bool, int) {
  modulator := move % 2
  if modulator == 1 { return g.player1.TakeATurn() }
  return g.player2.TakeATurn()
}
