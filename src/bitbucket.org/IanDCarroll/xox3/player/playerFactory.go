package player

import (
  "bitbucket.org/IanDCarroll/xox3/ui/display"
  "bitbucket.org/IanDCarroll/xox3/ui/selector"
  "bitbucket.org/IanDCarroll/xox3/ui"
  "bitbucket.org/IanDCarroll/xox3/rules"
  "bitbucket.org/IanDCarroll/xox3/board"
  "bitbucket.org/IanDCarroll/xox3/player/human"
)


func BuildHumanMNK3Player(playerNumber int, markers []string, board board.Board) player {
  ui := human.BuildHumanUi(markers, board)
  rules := buildMNK3Rules()
  return New(ui, rules, board, playerNumber)
}

func buildPlayerUI(display display.Display, selector selector.Selector) ui.Ui {
  return ui.New(display, selector)
}

func buildMNK3Rules() rules.Rules {
  return rules.NewMNKX(3)
}
