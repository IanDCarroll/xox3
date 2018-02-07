package menu

import (
  "bitbucket.org/IanDCarroll/xox3/ui"
  "bitbucket.org/IanDCarroll/xox3/ui/display/rec"
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
