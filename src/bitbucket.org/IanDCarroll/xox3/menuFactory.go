package main

import (
  "bitbucket.org/IanDCarroll/xox3/ui/shell"
  "bitbucket.org/IanDCarroll/xox3/ui/display"
  "bitbucket.org/IanDCarroll/xox3/ui/selector"
  "bitbucket.org/IanDCarroll/xox3/ui"
  "bitbucket.org/IanDCarroll/xox3/menu"
  "bitbucket.org/IanDCarroll/xox3/ui/display/rec"
)

var terminalOut shell.ShellOut = shell.NewTerminal()
var terminalIn shell.ShellIn = shell.NewTerminal()

func BuildMenu() menu.Menu {
  ui := BuildMenuUI()
  rec := BuildEnglishMenuRec()
  return menu.New(ui, rec)
}

func BuildEnglishMenuRec() rec.MenuRec {
  return rec.NewEnglishMenu()
}

func BuildMenuUI() ui.Ui {
  display := BuildMenuDisplay()
  selector := BuildMenuSelector()
  return ui.New(display, selector)
}

func BuildMenuDisplay() display.Display {
  return display.NewMenu(terminalOut)
}

func BuildMenuSelector() selector.Selector {
  return selector.NewMenu(terminalIn)
}
