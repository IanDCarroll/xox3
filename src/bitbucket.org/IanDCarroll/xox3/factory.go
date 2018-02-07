package main

import (
  "bitbucket.org/IanDCarroll/xox3/ui/shell"
  "bitbucket.org/IanDCarroll/xox3/ui/display"
  "bitbucket.org/IanDCarroll/xox3/ui/selector"
  "bitbucket.org/IanDCarroll/xox3/ui"
  "bitbucket.org/IanDCarroll/xox3/menu"
  "bitbucket.org/IanDCarroll/xox3/rec"
)

var TerminalShell shell.Shell = shell.NewTerminal()

func BuildMenu() menu.Menu {
  ui := BuildMenuUI()
  rec := BuildEnglishRec()
  return menu.New(ui, rec)
}

func BuildEnglishRec() rec.Rec {
  return rec.NewEnglish()
}

func BuildMenuUI() ui.Ui {
  display := BuildMenuDisplay()
  selector := BuildMenuSelector()
  return ui.New(display, selector)
}

func BuildMenuSelector() selector.Selector {
  return selector.NewMenu(TerminalShell)
}

func BuildMenuDisplay() display.Display {
  return display.NewMenu(TerminalShell)
}
