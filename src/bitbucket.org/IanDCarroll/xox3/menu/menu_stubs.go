package menu

import (
  "bitbucket.org/IanDCarroll/xox3/ui/shell"
  "bitbucket.org/IanDCarroll/xox3/ui/display"
  "bitbucket.org/IanDCarroll/xox3/ui/selector"
  "bitbucket.org/IanDCarroll/xox3/ui"
  "bitbucket.org/IanDCarroll/xox3/ui/display/rec"
)

type shellStub struct { shellOut string }
func (s shellStub) Read() string { return "2" }
func (s shellStub) Write(message string) { s.shellOut = message }

func buildEnglishMenuRec() rec.MenuRec {
  return rec.NewEnglishMenu()
}

func buildMenuUI(shell shell.Shell) ui.Ui {
  display := buildMenuDisplay(shell)
  selector := buildMenuSelector(shell)
  return ui.New(display, selector)
}

func buildMenuSelector(shell shell.Shell) selector.Selector {
  return selector.NewMenu(shell)
}

func buildMenuDisplay(shell shell.Shell) display.Display {
  return display.NewMenu(shell)
}
