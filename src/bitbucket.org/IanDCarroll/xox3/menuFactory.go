package main

import (
  "bitbucket.org/IanDCarroll/xox3/ui/shell"
  "bitbucket.org/IanDCarroll/xox3/ui/display"
  "bitbucket.org/IanDCarroll/xox3/ui/selector"
  "bitbucket.org/IanDCarroll/xox3/ui"
  "bitbucket.org/IanDCarroll/xox3/menu"
  "bitbucket.org/IanDCarroll/xox3/ui/display/rec"
)

func BuildMenu(shell shell.Shell) menu.Menu {
  ui := BuildMenuUI(shell)
  rec := BuildEnglishMenuRec()
  return menu.New(ui, rec)
}

func BuildEnglishMenuRec() rec.MenuRec {
  return rec.NewEnglishMenu()
}

func BuildMenuUI(shell shell.Shell) ui.Ui {
  display := BuildMenuDisplay(shell)
  selector := BuildMenuSelector(shell)
  return ui.New(display, selector)
}

func BuildMenuSelector(shell shell.Shell) selector.Selector {
  return selector.NewMenu(shell)
}

func BuildMenuDisplay(shell shell.Shell) display.Display {
  return display.NewMenu(shell)
}
