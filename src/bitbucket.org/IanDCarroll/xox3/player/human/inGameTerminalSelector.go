package human

import "bitbucket.org/IanDCarroll/xox3/ui/shell"

type terminalSelector struct {
  shell shell.ShellIn
}

func NewInGameTerminalSelector(shell shell.ShellIn) terminalSelector {
  return terminalSelector {shell}
}

func (t terminalSelector) ReadFromShell() interface{} {
  return t.shell.Read()
}
