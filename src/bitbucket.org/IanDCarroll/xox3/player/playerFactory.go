package player

import (
  "bitbucket.org/IanDCarroll/xox3/ui/display"
  "bitbucket.org/IanDCarroll/xox3/ui/selector"
  "bitbucket.org/IanDCarroll/xox3/ui"
  "bitbucket.org/IanDCarroll/xox3/rules"
  "bitbucket.org/IanDCarroll/xox3/board"
  "bitbucket.org/IanDCarroll/xox3/player/human"
  "bitbucket.org/IanDCarroll/xox3/player/human/rec"
)


func BuildHumanMNK3Player(marker int, rec rec.Rec, board board.Board) player {
  ui := human.BuildHumanUi(rec, board)
  rules := buildMNK3Rules()
  return New(ui, rules, board, marker)
}

func buildPlayerUI(display display.Display, selector selector.Selector) ui.Ui {
  return ui.New(display, selector)
}

func buildMNK3Rules() rules.Rules {
  return rules.NewMNKX(3)
}
