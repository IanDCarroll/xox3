package selector

import "bitbucket.org/IanDCarroll/xox3/ui/shell"

type menu struct {
  shell shell.Shell
}

func NewMenu(shell shell.Shell) menu {
  return menu {shell}
}

func (m menu) ReadFromShell() string {
  return m.shell.Read()
}
