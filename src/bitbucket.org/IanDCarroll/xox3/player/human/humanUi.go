package human

import (
  "bitbucket.org/IanDCarroll/xox3/ui/display"
  "bitbucket.org/IanDCarroll/xox3/ui/selector"
  "bitbucket.org/IanDCarroll/xox3/board"
  "bitbucket.org/IanDCarroll/xox3/player/human/messageRec"
  "strconv"
)

type humanUi struct {
  display display.Display
  selector selector.Selector
  board board.Board
  rec messageRec.Rec

}

func NewUI(d display.Display, s selector.Selector, b board.Board, r messageRec.Rec) humanUi {
  return humanUi {d, s, b, r}
}

func (ui humanUi) GiveToDisplay(message interface{}) {
  ui.display.WriteToShell(message)
}

func (ui humanUi) GetFromSelector() interface{} {
  rawInput := ui.selector.ReadFromShell()
  moveSelection, err := strconv.Atoi(rawInput.(string))
  if ui.invalid(moveSelection, err) { return ui.retry(ui.GetFromSelector) }
  return ui.zeroIndex(moveSelection)
}

func (ui humanUi) invalid(input int, err error) bool {
  return err != nil || ui.spaceIsUnavailable(input)
}

func (ui humanUi) spaceIsUnavailable(input int) bool {
  selection := ui.zeroIndex(input)
  availableSpaces := ui.board.AvailableSpaces()
  for i := 0; i < len(availableSpaces); i++ {
    if selection == availableSpaces[i] { return false }
  }
  return true
}

func (ui humanUi) zeroIndex(selection int) int { return selection - 1 }

func (ui humanUi) retry(function func() interface{}) interface{} {
  ui.GiveToDisplay(ui.rec.BadSelection())
  return function()
}
