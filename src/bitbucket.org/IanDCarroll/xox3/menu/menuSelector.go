package menu

import "bitbucket.org/IanDCarroll/xox3/ui/shell"

type menuSelector struct {
  shell shell.ShellIn
}

func NewMenuSelector(shell shell.ShellIn) menuSelector {
  return menuSelector {shell}
}

func (m menuSelector) ReadFromShell() string {
  return m.shell.Read().(string)
}
