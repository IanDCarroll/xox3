package ui


import (
  "bitbucket.org/IanDCarroll/xox3/ui/display"
  "bitbucket.org/IanDCarroll/xox3/ui/selector"
)

type Ui interface {
  GiveToDisplay(interface{})
  GetFromSelector() interface{}
}

type ui struct {
  display display.Display
  selector selector.Selector
}

func New(display display.Display, selector selector.Selector) ui {
  return ui {display, selector}
}

func (ui ui) GiveToDisplay(message interface{}) {
  ui.display.WriteToShell(message)
}

func (ui ui) GetFromSelector() interface{} {
  return ui.selector.ReadFromShell()
}
