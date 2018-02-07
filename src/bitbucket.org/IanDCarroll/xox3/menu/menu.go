package menu

import (
  "bitbucket.org/IanDCarroll/xox3/ui"
  "bitbucket.org/IanDCarroll/xox3/rec"
)

type Menu interface {
  GetGameParams() int
}

type menu struct {
  ui ui.Ui
  rec rec.Rec
}

func New(ui ui.Ui, rec rec.Rec) menu {
  return menu {ui, rec}
}

func (m menu) GetGameParams() int {
  m.ui.GiveToDisplay(m.rec.Welcome())
  return 1
}

func (m menu) Welcome() string {
  return `
  Welcome to xox3: an unbeatable game of noughts and crosses!
  Be amazed at your inability to ever win against this mighty juggernaut!

  ( Actual game coming soon... )
  `
}
