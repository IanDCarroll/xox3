package display

import "bitbucket.org/IanDCarroll/xox3/ui/shell"

type menu struct {
  shell shell.Shell
}

func NewMenu(shell shell.Shell) menu {
  return menu {shell}
}

func (m menu) WriteToShell(message string) {
  m.shell.Write(message)
}

func (m menu) Welcome() string {
  return `
  Welcome to xox3: an unbeatable game of noughts and crosses!
  Be amazed at your inability to ever win against this mighty juggernaut!

  ( Actual game coming soon... )
  `
}
