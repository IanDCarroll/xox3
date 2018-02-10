package player

import (
  "bitbucket.org/IanDCarroll/xox3/ui/display"
  "bitbucket.org/IanDCarroll/xox3/ui/selector"
  "bitbucket.org/IanDCarroll/xox3/ui"
  "bitbucket.org/IanDCarroll/xox3/rules"
  "bitbucket.org/IanDCarroll/xox3/board"
)

func buildPlayerUI(display display.Display, selector selector.Selector) ui.Ui {
  return ui.New(display, selector)
}

func buildMNK3Board() board.Board {
  return board.NewMNKX(3)
}

func buildMNK3Rules() rules.Rules {
  return rules.NewMNKX(3)
}
