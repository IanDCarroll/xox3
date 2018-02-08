package menu

import "bitbucket.org/IanDCarroll/xox3/ui/shell"

type menuDisplay struct {
  shell shell.ShellOut
}

func NewMenuDisplay(shell shell.ShellOut) menuDisplay {
  return menuDisplay {shell}
}

func (m menuDisplay) WriteToShell(message string) {
  m.shell.Write(message)
}
