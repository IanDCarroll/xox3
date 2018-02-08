package selector

import "bitbucket.org/IanDCarroll/xox3/ui/shell"

type menu struct {
  shell shell.ShellIn
}

func NewMenu(shell shell.ShellIn) menu {
  return menu {shell}
}

func (m menu) ReadFromShell() string {
  return m.shell.Read().(string)
}
