package human

import (
  "bitbucket.org/IanDCarroll/xox3/ui/display"
)

type terminalDisplay struct {
  messenger display.Display
  carpenter display.Display
}

var craftYourOwnResponse interface{}

func NewInGameTerminalDisplay(messenger, carpenter display.Display) terminalDisplay {
  return terminalDisplay {messenger, carpenter}
}

func (t terminalDisplay) WriteToShell(message interface{}) {
  t.messenger.WriteToShell(message)
  t.carpenter.WriteToShell(craftYourOwnResponse)
}
