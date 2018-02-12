package ai

import (
  "bitbucket.org/IanDCarroll/xox3/rules"
  "bitbucket.org/IanDCarroll/xox3/board"
)

type ai struct {
  rules rules.Rules
  board board.Board
  identity int
  opponent int
}

func NewMNKAi(rules rules.Rules, board board.Board) *ai {
  defaultId := 1
  defaultOp := 2
  return &ai {rules, board, defaultId, defaultOp}
}

func (ai *ai) SetID(id int) {
  if id != ai.identity {
    ai.opponent = ai.identity
    ai.identity = id
  }
}

func (ai *ai) GetMove() int {
  return ai.dumbButLegalMove()
}

func (ai *ai) dumbButLegalMove() int {
  return ai.board.AvailableSpaces()[0]
}
