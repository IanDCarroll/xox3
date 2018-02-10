package player

import (
  "bitbucket.org/IanDCarroll/xox3/ui"
  "bitbucket.org/IanDCarroll/xox3/rules"
  "bitbucket.org/IanDCarroll/xox3/board"
)

type Player interface {
  TakeATurn() (bool, int)
}

type player struct {
  ui ui.Ui
  rules rules.Rules
  board board.Board
  marker int
}

func New(ui ui.Ui, rules rules.Rules, board board.Board, marker int) player {
  return player {ui, rules, board, marker}
}

func (p player) TakeATurn() (bool, int) {
  p.showBoard()
  p.markASpace()
  return p.whetherTheGameIsWonOrDrawnOrNot()
}

func (p player) showBoard() {
  p.ui.GiveToDisplay(p.marker)
}

func (p player) markASpace() {
  moveChoice := p.ui.GetFromSelector().(int)
  p.board.Mark(moveChoice, p.marker)
}

func (p player) whetherTheGameIsWonOrDrawnOrNot() (bool, int) {
  return p.rules.Gameover(p.board.View())
}
