package menu

import (
  "strconv"
  "bitbucket.org/IanDCarroll/xox3/ui"
  "bitbucket.org/IanDCarroll/xox3/ui/display/rec"
)

type Menu interface {
  GetGameParams() (int, int)
}

type menu struct {
  ui ui.Ui
  rec rec.MenuRec
}

func New(ui ui.Ui, rec rec.MenuRec) menu {
  return menu {ui, rec}
}

func (m menu) GetGameParams() (int, int) {
  m.ui.GiveToDisplay(m.rec.Welcome())
  order := m.askForPlayOrder()
  marker := m.askForMarker()
  m.ui.GiveToDisplay(m.rec.StartGame())
  return order, marker
}

func (m menu) askForPlayOrder() int {
  m.ui.GiveToDisplay(m.rec.WhichPlayer())
  playerSeats := 2
  return m.ask(playerSeats)
}

func (m menu) askForMarker() int {
  m.ui.GiveToDisplay(m.rec.WhichMarker())
  markers := m.rec.Markers()
  return m.ask(len(markers))
}

func (m menu) ask(choices int) int {
  rawInput := m.ui.GetFromSelector()
  playOrder, err := strconv.Atoi(rawInput)
  if m.invalid(playOrder, err, choices) { return m.retry(m.askForPlayOrder) }
  return playOrder
}

func (m menu) invalid(input int, err error, choices int) bool {
  return err != nil || input > choices || input < 1
}

func (m menu) retry(function func() int) int {
  m.ui.GiveToDisplay(m.rec.BadSelection())
  return function()
}
