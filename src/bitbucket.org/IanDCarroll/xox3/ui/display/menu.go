package display

import "bitbucket.org/IanDCarroll/xox3/ui/shell"

type menu struct {
  shell shell.ShellOut
}

func NewMenu(shell shell.ShellOut) menu {
  return menu {shell}
}

func (m menu) WriteToShell(message string) {
  m.shell.Write(message)
}
