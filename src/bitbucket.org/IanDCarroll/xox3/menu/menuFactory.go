package menu

import (
  "bitbucket.org/IanDCarroll/xox3/ui/shell"
  "bitbucket.org/IanDCarroll/xox3/ui/display"
  "bitbucket.org/IanDCarroll/xox3/ui/selector"
  "bitbucket.org/IanDCarroll/xox3/ui"
  "bitbucket.org/IanDCarroll/xox3/menu/rec"
)

var terminalOut shell.ShellOut = shell.NewTerminal()
var terminalIn shell.ShellIn = shell.NewTerminal()

func BuildMenu() Menu {
  ui := buildMenuUI(terminalOut, terminalIn)
  rec := buildEnglishRec()
  return New(ui, rec)
}

func buildEnglishRec() rec.Rec {
  return rec.NewEnglish()
}

func buildMenuUI(shellOut shell.ShellOut, shellIn shell.ShellIn) ui.Ui {
  display := buildMenuDisplay(shellOut)
  selector := buildMenuSelector(shellIn)
  return ui.New(display, selector)
}

func buildMenuDisplay(shell shell.ShellOut) display.Display {
  return NewMenuDisplay(shell)
}

func buildMenuSelector(shell shell.ShellIn) selector.Selector {
  return NewMenuSelector(shell)
}
