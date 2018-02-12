package ai

import (
  "bitbucket.org/IanDCarroll/xox3/ui/display"
  "bitbucket.org/IanDCarroll/xox3/ui/selector"
  "bitbucket.org/IanDCarroll/xox3/ui"
  "bitbucket.org/IanDCarroll/xox3/rec"
  "bitbucket.org/IanDCarroll/xox3/rules"
  "bitbucket.org/IanDCarroll/xox3/board"
)

func BuildAiUi(board board.Board, humanSpy display.Display) ui.Ui {
  ai := buildMNK3Ai(board)
  display := buildAiSpyDisplay(humanSpy, ai)
  selector := buildAiSelector(ai)
  return ui.New(display, selector)
}

func buildAiSpyDisplay(humanSpy display.Display, ai Ai) display.Display {
  rec := rec.NewEnglish()
  return NewAiSpyDisplay(rec, humanSpy, ai)
}

func buildAiSelector(ai Ai) selector.Selector {
  return NewAiSelector(ai)
}

func buildMNK3Ai(board board.Board) Ai {
  rules := rules.NewMNKX(3)
  return NewMNKAi(rules, board)
}
