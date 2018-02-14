package ai

import (
  "bitbucket.org/IanDCarroll/xox3/rules"
  "bitbucket.org/IanDCarroll/xox3/board"
)

type ai struct {
  rules rules.Rules
  publicBoard board.Board
  identity int
  opponent int
}

func NewMNKAi(rules rules.Rules, publicBoard board.Board) *ai {
  defaultId := 1
  defaultOp := 2
  return &ai {rules, publicBoard, defaultId, defaultOp}
}

func (ai *ai) SetID(id int) {
  if id != ai.identity {
    ai.opponent = ai.identity
    ai.identity = id
  }
}

func (ai *ai) GetMove() int {
  boardCopy := board.Board(ai.publicBoard)
  initialDepth := 0
  initialChoices := map[int]int{}
  return ai.optimumChoice(boardCopy, initialDepth, initialChoices)
}

func (ai *ai) optimumChoice(board board.Board, depth int, choices map[int]int) int {
  over, report := ai.rules.Gameover(board.View())
  if over && report == 0 {
    return depth - 10
  } else if over && report == ai.identity {
    return 10000 - depth
  } else if over && report == ai.opponent {
    return depth - 10000
  }

  spaces := board.AvailableSpaces()
  for i := 0; i < len(spaces); i ++ {
    space := spaces[i]
    player := ai.appropriatePlayer(board)
    board.Mark(space, player)
    choices[space] = ai.optimumChoice(board, depth+1, map[int]int{})
    board.Mark(space, 0)
  }

  if depth != 0 {
    return ai.scoreTheSpace(choices, depth)
  }
  return ai.maxSpace(choices)
}

func (ai *ai) scoreTheSpace(choices map[int]int, depth int) int {
  if  depth % 2 == 1 { return ai.minScore(choices) }
  return ai.maxScore(choices)
}

func (ai *ai) minScore(choices map[int]int) int {
  min := 100
  for _, score := range choices {
    if score < min { min = score }
  }
  return min
}

func (ai *ai) maxScore(choices map[int]int) int {
  max := -100
  for _, score := range choices {
    if score > max { max = score }
  }
  return max
}

func (ai *ai) maxSpace(choices map[int]int) int {
  bestScore := ai.maxScore(choices)
  for space, score := range choices {
    if score == bestScore { return space }
  }
  return bestScore
}

func (ai *ai) appropriatePlayer(board board.Board) int {
  spacesTaken := board.Size() - len(board.AvailableSpaces())
  if spacesTaken % 2 == 0 { return 1 }
  return 2
}
