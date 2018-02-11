package ai

import (
  "bitbucket.org/IanDCarroll/xox3/rules"
  "bitbucket.org/IanDCarroll/xox3/board"
)

type ai struct {
  rules rules.Rules
  board board.Board
}

func NewMNKAi(rules rules.Rules, board board.Board) ai {
  return ai {rules, board}
}

func (ai ai) GetMove() int {
  return ai.dumbButLegalMove()
}

func (ai ai) dumbButLegalMove() int {
  return ai.board.AvailableSpaces()[0]
}
